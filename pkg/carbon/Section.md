---
description: Section renders a content container with Carbon's standard cds--content spacing for page sections and story layouts.
---

# Section

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Section(args ...any)`|Returns a `<section>` view with the `cds--content` class applied. Accepts content, `carbon.With(...)`, and standard `mvc` options.|

## Basic Usage

```go
content := carbon.Section(
 carbon.Head(2, "Overview"),
 carbon.Para("Section applies Carbon's standard content container spacing."),
)
```

## Appearance

Section always renders as a `<section>` with the `cds--content` class.

|Property|`With` and `Apply` values|
|----|----|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

Example:

```go
carbon.Section(
 carbon.With(carbon.ThemeG10),
 carbon.Head(2, "Filters"),
)
```

`Section` does not expose additional appearance attrs beyond theme application on the container.

## References

* [Page](Page.md)
* [Grid](Grid.md)
* [Carbon Design System](https://carbondesignsystem.com/elements/2x-grid/overview/)
