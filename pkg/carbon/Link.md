---
description: Link wraps Carbon's link component for inline, standalone, and icon-only navigation with accessible labeling and optional icon slots.
---

# Link

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Link(href string, args ...any)`|Returns a `cds-link` view. The `href` is always applied first; pass `carbon.With(...)`, text, and optional `carbon.Icon(...)` content in `args`.|

## Basic Usage

```go
link := carbon.Link(
 "#docs",
 carbon.With(carbon.LinkInline, carbon.SizeMedium),
 "Read the documentation",
)
```

Standalone and icon-only links use the same constructor:

```go
standalone := carbon.Link(
 "#docs",
 "Open docs",
 carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize20)),
)

iconOnly := carbon.Link(
 "#docs",
 carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize20)),
).SetLabel("Open docs")
```

## Appearance

|Property|Values|
|----|----|
|Size|`SizeSmall`, `SizeMedium`, `SizeLarge`|
|Variant|`LinkInline` for inline presentation|
|Icon slot|`carbon.Icon(...)` is normalized into the `icon` slot and inherits `currentColor`|

## State

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the link is not disabled.|
|`SetEnabled(bool)`|Enables or disables the link.|
|`Value() string`|Returns the current `href` value.|
|`SetValue(string)`|Sets or clears the `href` value.|
|`Label() string`|Returns the current accessible label (`aria-label`).|
|`SetLabel(string)`|Sets or clears the accessible label. Useful for icon-only links.|
|`Rel() string`|Returns the current `rel` attribute.|
|`SetRel(string)`|Sets or clears the `rel` attribute.|
|`Target() string`|Returns the current `target` attribute.|
|`SetTarget(string)`|Sets or clears the `target` attribute.|

## Events

|Event|Description|
|----|----|
|`EventClick`|Fires when the link is activated.|
|`EventHover`|Fires when the pointer enters the link.|
|`EventNoHover`|Fires when the pointer leaves the link.|

## Notes

* Icon-only links should always set an accessible label with `SetLabel(...)`.
* Decorative link icons default to `aria-hidden="true"` unless the icon already has its own accessible label.
* Use `SetTarget("_blank")` and an appropriate `SetRel(...)` value for external navigation.

## References

* [Carbon Design System](https://carbondesignsystem.com/components/link/usage/)
