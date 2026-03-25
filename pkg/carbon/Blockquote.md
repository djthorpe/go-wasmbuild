---
description: Blockquote renders quoted content using Carbon-styled block quotation markup with optional citation or attribution text.
---

# Blockquote

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Blockquote(args ...any)`|Returns a styled blockquote view backed by a figure template with quote body and optional label slots.|

## Basic Usage

```go
quote := carbon.Blockquote(
 "Good content systems need a clear typographic hierarchy so that long-form material remains readable without losing structure or emphasis.",
).SetLabel("Carbon text storybook")
```

`Blockquote` accepts the quoted body as child content and supports normal `mvc` options at construction time:

```go
quote := carbon.Blockquote(
 mvc.WithID("quote"),
 "Quoted content goes here.",
)
```

## State

|Method|Description|
|----|----|
|`Label() string`|Returns the current citation or attribution text.|
|`SetLabel(string)`|Sets or replaces the citation or attribution text in the blockquote label slot.|

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
