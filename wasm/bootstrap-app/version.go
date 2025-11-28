package main

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	GitSource   string
	GitTag      string
	GitBranch   string
	GitHash     string
	GoBuildTime string
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func Version() string {
	if GitTag != "" {
		return GitTag
	}
	if GitBranch != "" {
		return GitBranch
	}
	if GitHash != "" {
		return GitHash
	}
	return "dev"
}
