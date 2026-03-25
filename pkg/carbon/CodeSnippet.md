---
description: CodeSnippet renders Carbon's single-line copyable code snippet component for commands, expressions, and other one-line examples.
---

# CodeSnippet

## Constructors

|Constructor|Description|
|----|----|
|`carbon.CodeSnippet(args ...any)`|Returns a `cds-code-snippet` with `type="single"` for one-line commands or expressions with copy support.|

## Basic Usage

```go
snippet := carbon.CodeSnippet("GOOS=js GOARCH=wasm go build ./...")
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Variant|Always `type="single"` via `carbon.CodeSnippet(...)`|
|Boolean appearance attrs|`CodeWrapText`, `CodeHideCopyButton`|
|Custom appearance opts|`WithCodeFeedback(string)`, `WithCodeCopyText(string)`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

Example:

```go
carbon.CodeSnippet(
 carbon.With(carbon.ThemeG10),
 carbon.WithCodeFeedback("Copied command"),
 carbon.WithCodeCopyText("go test ./..."),
 "go test ./...",
)
```

## State

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the snippet is not disabled.|
|`SetEnabled(bool)`|Enables or disables snippet interaction, including copy affordances.|

## Notes

* Use `carbon.CodeSnippet(...)` for single commands users are likely to copy directly.
* Inline snippets use the separate `carbon.Code(...)` constructor.
* Multi-line snippets use `carbon.CodeBlock(...)`.
* `CodeWrapText`, `CodeHideCopyButton`, `WithCodeFeedback(...)`, and `WithCodeCopyText(...)` are presentation options, so apply them during construction or with `Apply(...)`.

## References

* [Code](Code.md)
* [CodeBlock](CodeBlock.md)
* [Carbon Design System](https://carbondesignsystem.com/components/code-snippet/usage/)
