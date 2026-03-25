package carbon

// CodeBlock returns a <cds-code-snippet type="multi"> — a multi-line code
// block that collapses long content behind a "Show more" button.
//
//	carbon.CodeBlock("line1\nline2\nline3")
func CodeBlock(args ...any) *codeSnippet {
	return newCodeSnippet("multi", args...)
}
