package carbon

// Code returns a <cds-code-snippet type="inline"> suitable for embedding short
// code fragments within a sentence.
//
//	carbon.Code("go build ./...")
func Code(args ...any) *codeSnippet {
	return newCodeSnippet("inline", args...)
}
