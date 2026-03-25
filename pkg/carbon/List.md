---
description: List provides semantic unordered lists for structured content in Carbon pages and markdown output.
---

# List

## Constructors

|Constructor|Description|
|----|----|
|`carbon.List(args ...any)`|Returns an unordered list (`<ul>`) view. Non-`ListItem` children are wrapped automatically in list items.|

## Basic Usage

```go
items := carbon.List(
 carbon.ListItem("Primary action"),
 carbon.ListItem("Secondary action"),
 carbon.ListItem("Tertiary action"),
)
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Unordered list styles|`ListDisc`, `ListCircle`, `ListSquare`|
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

* `List(...)` automatically wraps non-`ListItem` children in `<li>` elements.
* List style attrs are applied with `carbon.With(...)` or `Apply(...)` on the list container, not on individual list items.

## References

* [OrderedList](OrderedList.md)
* [ListItem](ListItem.md)
* [Markdown](Markdown.md)
* [Carbon Design System](https://carbondesignsystem.com/patterns/content-patterns/global-elements/#lists)
