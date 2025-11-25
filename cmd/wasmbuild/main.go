package main

import (
	// Packages
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Context struct {
	Go       string `default:"go" help:"Path to go tool"`
	WasmExec string `default:"lib/wasm/wasm_exec.js:misc/wasm/wasm_exec.js:targets/wasm_exec.js" help:"Path to wasm_exec.js relative to GOROOT"`
	GoFlags  string `help:"Additional flags to pass to go build"`
	Config   string `default:"wasmbuild.yaml" help:"Path to configuration YAML file (relative to source path)"`
	Verbose  bool   `short:"v" help:"Enable verbose output"`

	// Private
	log    *Logger
	ctx    context.Context
	cancel context.CancelFunc
}

type CLI struct {
	Context
	Build BuildCmd `cmd:"" help:"Build a WASM application"`
	Serve ServeCmd `cmd:"" help:"Serve a WASM application"`
	Dep   DepCmd   `cmd:"" help:"Show dependencies of a WASM application"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func main() {
	cli := new(CLI)
	kong := kong.Parse(cli)

	// Additional context setup
	cli.Context.log = NewLogger(cli.Verbose)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	cli.Context.ctx = ctx
	cli.Context.cancel = cancel
	defer cancel()

	// Run the selected command
	if err := kong.Run(&cli.Context); err != nil {
		cli.Context.log.Error(err)
		os.Exit(-1)
	}
}
