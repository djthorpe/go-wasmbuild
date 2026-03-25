---
description: Markdown renders parsed markdown content into Carbon text helpers and semantic HTML for documentation-style pages.
---

# Markdown

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Markdown(text string, args ...any)`|Returns a block-level markdown view rendered inside a `DIV`. The `text` is parsed immediately and any additional `args` can include standard `mvc` options plus markdown-specific options such as `carbon.WithMarkdownLinkResolver(...)`.|

## Basic Usage

```go
doc := carbon.Markdown(`## Release Notes

Markdown supports **strong text**, *emphasis*, lists, links, tables, blockquotes, and fenced code blocks.
`)
```

You can rewrite markdown links before they are rendered:

```go
doc := carbon.Markdown(
 "See [Button](Button.md) for action patterns.",
 carbon.WithMarkdownLinkResolver(func(url string) string {
  if strings.HasSuffix(url, ".md") {
   return "#docs/" + strings.TrimSuffix(url, ".md")
  }
  return url
 }),
)
```

## Appearance

|Property|Description|
|----|----|
|Root element|Rendered as a `DIV` with the `markdown` class applied by default.|
|Content mapping|Headings, paragraphs, emphasis, deletion, lists, blockquotes, tables, images, links, rules, and code are expanded into Carbon or semantic HTML views.|
|Theme|Markdown inherits the theme of its parent container, such as `carbon.ThemeWhite` or `carbon.ThemeG10`.|

## Notes

* Relative markdown links can be rewritten with `carbon.WithMarkdownLinkResolver(...)` before they are emitted into the DOM.

## References

* [Blockquote](Blockquote.md)
* [CommonMark](https://commonmark.org/)
