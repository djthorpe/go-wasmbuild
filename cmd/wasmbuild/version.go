package main

import "github.com/djthorpe/go-wasmbuild/pkg/version"

// Packages

///////////////////////////////////////////////////////////////////////////////
// TYPES

type VersionCmd struct {
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *VersionCmd) Run(ctx *Context) error {
	if ctx.Verbose {
		if version.GitSource != "" {
			ctx.log.Info("Git Source: ", version.GitSource)
		}
		if version.GitTag != "" {
			ctx.log.Info("Git Tag: ", version.GitTag)
		}
		if version.GitBranch != "" {
			ctx.log.Info("Git Branch: ", version.GitBranch)
		}
		if version.GitHash != "" {
			ctx.log.Info("Git Hash: ", version.GitHash)
		}
		if version.GoBuildTime != "" {
			ctx.log.Info("Build Time: ", version.GoBuildTime)
		}
		ctx.log.Info("Compiler: ", version.Compiler())
	} else {
		ctx.log.Info(version.Version())
	}
	return nil
}
