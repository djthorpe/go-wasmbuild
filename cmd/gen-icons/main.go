// gen-icons generates icon_names.go by parsing the icon registry in
// npm/carbon/icons-generated.js. Run via go generate in the target package.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	pkg := flag.String("pkg", "carbon", "package name for the generated file")
	typeName := flag.String("type", "IconName", "type to use for icon name values")
	structName := flag.String("struct", "", "if set, generate []StructName{ID, Name} entries instead of []TypeName")
	importPath := flag.String("importpath", "", "import path for the icon name type (if external)")
	importAlias := flag.String("importalias", "", "import alias for -importpath")
	outPath := flag.String("out", "icon_names.go", "output file path")
	flag.Parse()

	// Walk up from CWD to find npm/carbon/icons-generated.js or index.js.
	indexPath := findIconsJS()
	if indexPath == "" {
		fmt.Fprintln(os.Stderr, "gen-icons: cannot find npm/carbon/icons-generated.js or npm/carbon/index.js")
		os.Exit(1)
	}

	names, err := parseIconNames(indexPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "gen-icons:", err)
		os.Exit(1)
	}
	if len(names) == 0 {
		fmt.Fprintln(os.Stderr, "gen-icons: no icon names found in", indexPath)
		os.Exit(1)
	}

	if err := writeOutput(*outPath, *pkg, *typeName, *structName, *importPath, *importAlias, names); err != nil {
		fmt.Fprintln(os.Stderr, "gen-icons:", err)
		os.Exit(1)
	}
	fmt.Printf("gen-icons: wrote %d icons to %s\n", len(names), *outPath)
}

// findIconsJS walks up from CWD looking for npm/carbon/icons-generated.js,
// falling back to npm/carbon/index.js if the generated file is absent.
func findIconsJS() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		p := filepath.Join(dir, "npm", "carbon", "icons-generated.js")
		if _, err := os.Stat(p); err == nil {
			return p
		}
		p2 := filepath.Join(dir, "npm", "carbon", "index.js")
		if _, err := os.Stat(p2); err == nil {
			return p2
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

// parseIconNames extracts keys from the goWasmBuildCarbonIcons object in the JS file.
func parseIconNames(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Matches lines like:  add:  or  'arrow--right':  or  "4K":  or  AI:
	keyRE := regexp.MustCompile(`^\s+['"]?([A-Za-z0-9][A-Za-z0-9\-]*(?:--[A-Za-z0-9\-]+)*)['"]?\s*:`)

	var names []string
	inBlock := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "goWasmBuildCarbonIcons = {") {
			inBlock = true
			continue
		}
		if inBlock {
			if strings.HasPrefix(strings.TrimSpace(line), "};") {
				break
			}
			if m := keyRE.FindStringSubmatch(line); m != nil {
				names = append(names, m[1])
			}
		}
	}
	sort.Strings(names)
	return names, scanner.Err()
}

// idToName converts an icon ID to a display name.
// "--" separates a base name from a variant (e.g. "arrow--right" → "arrow: right"),
// while a single "-" is a word separator (e.g. "zoom-in" → "zoom in").
// Known acronyms at the start of an ID are uppercased.
func idToName(id string) string {
	s := strings.ReplaceAll(id, "--", ": ")
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.Join(strings.Fields(s), " ")
	// Uppercase leading acronyms.
	for _, acronym := range []string{
		// IBM mainframe / middleware
		"BSAM", "CICS", "CICSPLEX", "COBOL", "DB2", "DFSORT", "IMS", "TSQ", "ZOS",
		// IBM products / brands
		"ASM", "BPMN", "CAD", "CBL", "CDN", "CEO", "CICS", "CLP", "CMS",
		"CO2", "CPU", "CXL",
		// Networking / protocols
		"DNS", "FTP", "HTTP", "HTTPS", "IP", "MQTT", "NFS", "P2P",
		"RSS", "SAN", "SSH", "SSL", "TCP", "TLS", "UDP", "VPN", "VLAN", "VMDK",
		// General tech
		"AI", "API", "EDT", "GPU", "GUI", "HTML", "IBM", "ID", "ICA", "IoT",
		"IT", "JS", "JSON", "ML", "NLP", "OWASP", "PDLC", "PDF", "PHP",
		"QR", "RAM", "REST", "SDK", "SLO", "SQL", "SVG", "UI", "URL", "USB",
		"UV", "VM", "WiFi", "XML", "YAML",
	} {
		lower := strings.ToLower(acronym)
		if s == lower || strings.HasPrefix(s, lower+" ") || strings.HasPrefix(s, lower+":") {
			s = acronym + s[len(lower):]
			break
		}
	}
	// Capitalise the first letter if it is still lowercase.
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	// Replace acronyms that appear as whole words anywhere in the name
	// (i.e. after a space or ": "). Order matters: longer matches first.
	// Replace compound terms that need special joining.
	s = strings.ReplaceAll(s, "TCP ip", "TCP/IP")
	s = strings.ReplaceAll(s, "TCP IP", "TCP/IP")
	// Replace acronyms that appear as whole words anywhere in the name
	// (i.e. after a space or ": "). Order matters: longer matches first.
	for _, repl := range []struct{ from, to string }{
		{"epdf", "ePDF"},
		{"pdf", "PDF"},
		{"db2", "DB2"},
		{"zos", "ZOS"},
		{"christian", "Christian"},
		{"jewish", "Jewish"},
		{"muslim", "Muslim"},
	} {
		s = strings.ReplaceAll(s, " "+repl.from, " "+repl.to)
		s = strings.ReplaceAll(s, ": "+repl.from, ": "+repl.to)
	}
	return s
}

func writeOutput(path, pkg, typeName, structName, importPath, importAlias string, names []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "// Code generated by go generate; DO NOT EDIT.")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "package %s\n", pkg)
	fmt.Fprintln(w, "")

	// Build import block.
	var imports []string
	imports = append(imports, `"strings"`)
	if importPath != "" {
		if importAlias != "" {
			imports = append(imports, fmt.Sprintf("%s %q", importAlias, importPath))
		} else {
			imports = append(imports, fmt.Sprintf("%q", importPath))
		}
	}
	fmt.Fprintln(w, "import (")
	for _, imp := range imports {
		fmt.Fprintf(w, "\t%s\n", imp)
	}
	fmt.Fprintln(w, ")")
	fmt.Fprintln(w, "")

	if structName != "" {
		// Struct-based output: []StructName with ID and Name fields.
		fmt.Fprintf(w, "var iconRegistry = []%s{\n", structName)
		for _, name := range names {
			fmt.Fprintf(w, "\t{ID: %s(%q), Name: %q},\n", typeName, name, idToName(name))
		}
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "// Icons returns all bundled icons, optionally filtered by substring queries.")
		fmt.Fprintln(w, "// Searches both the display Name and the raw ID. All terms must match (AND semantics); matching is case-insensitive.")
		fmt.Fprintf(w, "func Icons(q ...string) []%s {\n", structName)
		fmt.Fprintf(w, "\tvar result []%s\n", structName)
		fmt.Fprintln(w, "outer:")
		fmt.Fprintln(w, "\tfor _, entry := range iconRegistry {")
		fmt.Fprintln(w, "\t\thaystack := strings.ToLower(entry.Name + \" \" + string(entry.ID))")
		fmt.Fprintln(w, "\t\tfor _, term := range q {")
		fmt.Fprintln(w, "\t\t\tif !strings.Contains(haystack, term) {")
		fmt.Fprintln(w, "\t\t\t\tcontinue outer")
		fmt.Fprintln(w, "\t\t\t}")
		fmt.Fprintln(w, "\t\t}")
		fmt.Fprintln(w, "\t\tresult = append(result, entry)")
		fmt.Fprintln(w, "\t}")
		fmt.Fprintln(w, "\treturn result")
		fmt.Fprintln(w, "}")
	} else {
		// Simple slice output: []TypeName.
		fmt.Fprintf(w, "var iconNames = []%s{\n", typeName)
		for _, name := range names {
			fmt.Fprintf(w, "\t%q,\n", name)
		}
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "// Icons returns all bundled icon names, optionally filtered by substring queries.")
		fmt.Fprintln(w, "// All terms must match (AND semantics).")
		fmt.Fprintf(w, "func Icons(q ...string) []%s {\n", typeName)
		fmt.Fprintf(w, "\tvar result []%s\n", typeName)
		fmt.Fprintln(w, "outer:")
		fmt.Fprintln(w, "\tfor _, name := range iconNames {")
		fmt.Fprintln(w, "\t\tfor _, term := range q {")
		fmt.Fprintf(w, "\t\t\tif !strings.Contains(string(name), term) {\n")
		fmt.Fprintln(w, "\t\t\t\tcontinue outer")
		fmt.Fprintln(w, "\t\t\t}")
		fmt.Fprintln(w, "\t\t}")
		fmt.Fprintln(w, "\t\tresult = append(result, name)")
		fmt.Fprintln(w, "\t}")
		fmt.Fprintln(w, "\treturn result")
		fmt.Fprintln(w, "}")
	}
	return w.Flush()
}
