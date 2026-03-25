---
description: Page renders a plain wrapper without Carbon's cds--content spacing, suitable for composing full custom page layouts inside a Section.
---

# Page

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Page(args ...any)`|Returns a plain `<div>` view with no default Carbon content class. Accepts content, `carbon.With(...)`, and standard `mvc` options.|

## Basic Usage

```go
page := carbon.Page(
 carbon.Head(1, "Dashboard"),
 carbon.Grid(
  carbon.Col4(carbon.Compact("Sidebar")),
  carbon.Col12(carbon.Para("Primary content")),
 ),
)
```

## Appearance

Page always renders as a plain `<div>` with no default Carbon content class.

|Property|`With` and `Apply` values|
|----|----|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

Example:

```go
carbon.Page(
 carbon.With(carbon.ThemeWhite),
 carbon.Head(1, "Workspace"),
)
```

`Page` does not expose additional appearance attrs beyond theme application on the container.

## Notes

* Use `Page` when you need a neutral wrapper without the default `cds--content` styling.
* In the example app, `Page` is used as the per-route wrapper inside the main content `Section`.

## References

* [Section](Section.md)
* [Grid](Grid.md)
* [Carbon Design System](https://carbondesignsystem.com/elements/2x-grid/overview/)
