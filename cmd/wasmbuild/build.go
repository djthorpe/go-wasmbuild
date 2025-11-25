package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	// Packages
	"github.com/djthorpe/go-wasmbuild/etc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type BuildPath struct {
	Path string `arg:"" default:"." help:"Source path to WASM application"`
}

type BuildCmd struct {
	BuildPath
	Output string `short:"o" help:"Output path (uses temp dir if not specified)"`
}

type BuildContext struct {
	Config

	// Source path for build
	Path string `json:"input,omitempty"`

	// Output path for build
	Output string `json:"output,omitempty"`

	// Go tool command and flags
	GoCmd  string   `json:"go_cmd,omitempty"`
	GoRoot string   `json:"go_root,omitempty"`
	GoArgs []string `json:"go_args,omitempty"`
	GoEnv  []string `json:"go_env,omitempty"`

	// WasmExec Javascript path
	WasmExecJS   *File `json:"wasm_exec_js,omitempty"`
	WasmExecHTML *File `json:"wasm_exec_html,omitempty"`
	FavIcon      *File `json:"favicon,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// BuildContext creates a BuildContext from the Config, returning all the
// information needed to build a WASM application.
func (c Config) BuildContext(ctx *Context, path, output string, watch bool) (*BuildContext, error) {
	context := &BuildContext{
		Config: c,
		GoCmd:  ctx.Go,
	}

	// Make input path absolute
	if !filepath.IsAbs(path) {
		var err error
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, fmt.Errorf("failed to determine absolute path: %w", err)
		}
	}
	context.Path = path

	// Determine the output path
	if output == "" {
		// Use temporary directory
		var err error
		output, err = os.MkdirTemp("", fmt.Sprintf("wasmbuild-%s-*", filepath.Base(path)))
		if err != nil {
			return nil, fmt.Errorf("failed to create temp directory: %w", err)
		}
	} else {
		// Make output path absolute
		if filepath.IsAbs(output) == false {
			var err error
			output, err = filepath.Abs(output)
			if err != nil {
				return nil, fmt.Errorf("failed to determine absolute path: %w", err)
			}
		}

		// Ensure build directory exists
		if err := os.MkdirAll(output, 0755); err != nil {
			return nil, fmt.Errorf("failed to create build directory: %w", err)
		}
	}
	context.Output = output

	// goroot
	gocmd, goroot, err := GoEnvFromCmd(ctx.Go)
	if err != nil {
		return nil, fmt.Errorf("failed to determine GOROOT or TINYGOROOT: %w", err)
	}
	context.GoCmd = gocmd
	context.GoRoot = goroot

	// goargs
	context.GoArgs = append([]string{"build"}, strings.Fields(ctx.GoFlags)...)

	// goenv
	context.GoEnv = []string{"GOOS=js", "GOARCH=wasm"}

	// wasm_exec.js
	wasmPathExecJS := RegularFileFromPathList(ctx.WasmExec, goroot)
	if wasmPathExecJS == "" {
		return nil, fmt.Errorf("wasm_exec.js not found in GOROOT")
	}
	wasmFileExecJS, err := NewFileFromSource(wasmPathExecJS, "wasm_exec.js")
	if err != nil {
		return nil, fmt.Errorf("failed to read wasm_exec.js: %w", err)
	}
	context.WasmExecJS = wasmFileExecJS

	//  wasm_exec.html
	funcs := template.FuncMap{
		"Title": func() string {
			if title, ok := c.Vars["Title"]; ok {
				return title
			}
			return filepath.Base(path)
		},
		"Header": func() string {
			if head, ok := c.Vars["Header"]; ok {
				return head
			}
			return ""
		},
		"Footer": func() string {
			if foot, ok := c.Vars["Footer"]; ok {
				return foot
			}
			return ""
		},
		"Notify": func() string {
			if watch {
				if notify, ok := c.Vars["Notify"]; ok {
					return notify
				} else {
					return string(etc.NotifyHTML)
				}
			}
			return ""
		},
		"WasmFile": func() string {
			if wasmFile, ok := c.Vars["WasmFile"]; ok {
				return wasmFile
			}
			return filepath.Base(path) + ".wasm"
		},
	}
	wasmExecHTML, err := NewFileFromTemplate(etc.WasmExecHTML, "wasm_exec.html", c.Vars, funcs)
	if err != nil {
		return nil, fmt.Errorf("failed to create wasm_exec.html: %w", err)
	}
	context.WasmExecHTML = wasmExecHTML

	// favicon
	context.FavIcon = NewFile(etc.FaviconPNG, "favicon.png")

	// Return build context
	return context, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *BuildContext) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *BuildCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveFile(ctx.Config, c.Path)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create a compiler context from the configuration
	buildContext, err := config.BuildContext(ctx, c.Path, c.Output, false)
	if err != nil {
		return err
	}

	// Compile
	file, err := buildContext.CompileExec(ctx)
	if err != nil {
		return err
	}

	// Copy files to output directory
	for _, files := range []*File{
		file,
		buildContext.WasmExecHTML,
		buildContext.WasmExecJS,
		buildContext.FavIcon,
	} {
		// Write file
		if sz, err := files.WriteTo(buildContext.Output); err != nil {
			return fmt.Errorf("failed to copy %s: %w", files.Path, err)
		} else {
			ctx.log.Info("cp ", files.Path, " ", buildContext.Output, " (", sz, " bytes)")
		}
	}

	// Copy assets to output directory
	for _, asset := range config.Assets {
		if filepath.IsAbs(asset) == false {
			asset = filepath.Join(c.Path, asset)
		}
		dest := fmt.Sprintf("%s/%s", buildContext.Output, filepath.Base(asset))

		// Walk the asset path and copy files
		err := filepath.Walk(asset, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasPrefix(info.Name(), ".") {
				return nil
			}

			// Relative path from asset root
			relPath, err := filepath.Rel(asset, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(dest, relPath)
			if info.Mode().IsDir() {
				ctx.log.Info("mkdir ", destPath)
				if err := os.MkdirAll(destPath, 0755); err != nil {
					return err
				}
				// Walk into directory
				return nil
			}

			// Copy file across
			if sz, err := CopyFile(path, destPath); err != nil {
				return err
			} else {
				ctx.log.Info("cp ", path, " ", filepath.Dir(destPath), " (", sz, " bytes)")
			}

			// Return success
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to copy asset %s: %w", asset, err)
		}
	}

	// Print out the destination to stdout
	fmt.Println(buildContext.Output)

	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Return a exec.Cmd for building the WASM application
func (bc *BuildContext) GoBuildCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(bc.GoCmd, append(bc.GoArgs, args...)...)
	cmd.Dir = bc.Path
	cmd.Env = append(os.Environ(), bc.GoEnv...)
	return cmd
}

func (c *BuildContext) CompileExec(ctx *Context) (*File, error) {
	// Create temporary directory for build
	tmpDir, err := os.MkdirTemp("", "wasmbuild-compile-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Build to temporary location
	wasmPath := filepath.Join(tmpDir, filepath.Base(c.Path)+".wasm")
	cmd := c.GoBuildCmd("-o", wasmPath)
	// Update the output path to temp location
	for i, arg := range cmd.Args {
		if arg == "-o" && i+1 < len(cmd.Args) {
			cmd.Args[i+1] = wasmPath
			break
		}
	}

	// Log the compile
	ctx.log.Info(cmd.String())

	// Capture stderr for error messages
	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	// Run the command
	if err := cmd.Run(); err != nil {
		stderr := stderrBuf.String()
		return nil, fmt.Errorf("compilation failed: %w\n%s", err, stderr)
	}

	// Read the compiled wasm file into memory
	wasmData, err := os.ReadFile(wasmPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read compiled wasm file: %w", err)
	}

	// Return as File object
	return NewFile(wasmData, filepath.Base(c.Path)+".wasm"), nil
}

func CopyFile(src, dest string) (int64, error) {
	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	// Create destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	// Copy data
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		return 0, err
	}

	// Return success
	return n, nil
}

// Return the path to file wasm_exec.js
func RegularFileFromPathList(path, base string) string {
	for _, path := range strings.Split(path, string(filepath.ListSeparator)) {
		if filepath.IsAbs(path) == false {
			path = filepath.Join(base, path)
		}
		if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
			return path
		}
	}
	return ""
}

func ResolveFile(path, base string) (string, error) {
	if filepath.IsAbs(path) == false {
		if base == "" {
			base = "."
		}
		path = filepath.Join(base, path)
	}
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if stat.IsDir() {
		return "", fmt.Errorf("expected file but found directory: %s", path)
	}
	return path, nil
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func GoEnvFromCmd(cmd string) (string, string, error) {
	if !filepath.IsAbs(cmd) {
		var err error
		cmd, err = exec.LookPath(cmd)
		if err != nil {
			return "", "", fmt.Errorf("failed to locate go executable: %w", err)
		}
	}

	// Run 'tinygo env TINYGOROOT' to get TINYGOROOT
	output, _ := exec.Command(cmd, "env", "TINYGOROOT").Output()
	if output = bytes.TrimSpace(output); len(output) > 0 {
		return cmd, string(output), nil
	}

	// Run 'go env GOROOT' to get GOROOT
	output, err = exec.Command(cmd, "env", "GOROOT").Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to determine GOROOT: %w", err)
	}
	if output = bytes.TrimSpace(output); len(output) > 0 {
		return cmd, string(output), nil
	}
	return "", "", fmt.Errorf("GOROOT not found")
}

/*

	// wasm_exec.js
	wasmPathExecJS := RegularFileFromPathList(ctx.WasmExec, goroot)
	if wasmPathExecJS == "" {
		return nil, fmt.Errorf("wasm_exec.js not found in GOROOT")
	}


	// Return success
}
*/
