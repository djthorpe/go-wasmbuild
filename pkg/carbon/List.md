---
description: List provides semantic unordered and ordered lists, plus list items, for structured content in Carbon pages and markdown output.
---

# List

## Constructors

|Constructor|Description|
|----|----|
|`carbon.List(args ...any)`|Returns an unordered list (`<ul>`) view. Non-`ListItem` children are wrapped automatically in list items.|
|`carbon.OrderedList(args ...any)`|Returns an ordered list (`<ol>`) view. Non-`ListItem` children are wrapped automatically in list items.|
|`carbon.ListItem(args ...any)`|Returns a list item (`<li>`) view.|

## Basic Usage

```go
items := carbon.List(
 carbon.ListItem("Primary action"),
 carbon.ListItem("Secondary action"),
 carbon.ListItem("Tertiary action"),
)
```

Ordered lists use the same item helper:

```go
steps := carbon.OrderedList(
 carbon.ListItem("Create the workspace"),
 carbon.ListItem("Add the route"),
 carbon.ListItem("Validate the build"),
)
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Unordered list styles|`ListDisc`, `ListCircle`, `ListSquare`|
|Ordered list styles|`ListDecimal`, `ListLowerAlpha`, `ListUpperAlpha`, `ListLowerRoman`, `ListUpperRoman`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

Example:

```go
carbon.List(
 carbon.With(carbon.ListSquare),
 carbon.ListItem("Primary action"),
 carbon.ListItem("Secondary action"),
)
```

## Notes

* `List(...)` and `OrderedList(...)` automatically wrap non-`ListItem` children in `<li>` elements.
* List style attrs are applied with `carbon.With(...)` or `Apply(...)` on the list container, not on individual list items.

## References

* [Markdown](Markdown.md)
* [Carbon Design System](https://carbondesignsystem.com/patterns/content-patterns/global-elements/#lists)
