---
description: OrderedList provides semantic ordered lists for step-by-step or ranked content in Carbon pages and markdown output.
---

# OrderedList

## Constructors

|Constructor|Description|
|----|----|
|`carbon.OrderedList(args ...any)`|Returns an ordered list (`<ol>`) view. Non-`ListItem` children are wrapped automatically in list items.|

## Basic Usage

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
|Ordered list styles|`ListDecimal`, `ListLowerAlpha`, `ListUpperAlpha`, `ListLowerRoman`, `ListUpperRoman`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

## References

* [List](List.md)
* [ListItem](ListItem.md)
* [Carbon Design System](https://carbondesignsystem.com/patterns/content-patterns/global-elements/#lists)
