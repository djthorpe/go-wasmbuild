---
description: CodeBlock documents Carbon's multi-line Carbon code snippet component for expandable source samples and longer examples.
---

# CodeBlock

## Constructors

|Constructor|Description|
|----|----|
|`carbon.CodeBlock(args ...any)`|Returns a `cds-code-snippet` with `type="multi"` for longer multi-line snippets that can expand and collapse.|

## Basic Usage

```go
block := carbon.CodeBlock(`package main

func main() {
    println("hello")
}`)
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Variant|Always `type="multi"` via `carbon.CodeBlock(...)`|
|Boolean appearance attrs|`CodeWrapText`, `CodeHideCopyButton`|
|Custom appearance opts|`WithCodeFeedback(string)`, `WithCodeCopyText(string)`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

Example:

```go
carbon.CodeBlock(
 carbon.With(carbon.ThemeG10, carbon.CodeWrapText),
 carbon.WithCodeFeedback("Copied snippet"),
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

* Use `carbon.CodeBlock(...)` for longer examples, source samples, or configuration fragments.
* Single-line snippets use `carbon.CodeSnippet(...)`.
* Inline snippets use the separate `carbon.Code(...)` constructor.
* `CodeWrapText`, `CodeHideCopyButton`, `WithCodeFeedback(...)`, and `WithCodeCopyText(...)` are presentation options, so apply them during construction or with `Apply(...)`.

## References

* [Code](Code.md)
* [CodeSnippet](CodeSnippet.md)
* [Carbon Design System](https://carbondesignsystem.com/components/code-snippet/usage/)
