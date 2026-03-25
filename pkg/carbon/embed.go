package carbon

import "embed"

// DocsFS embeds the Markdown documentation files for Carbon components.
//
//go:embed *.md
var DocsFS embed.FS
