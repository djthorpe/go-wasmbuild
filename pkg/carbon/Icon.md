---
description: Icon renders a Carbon icon from the bundled registry, with size and accessible labeling applied through the icon wrapper.
---

# Icon

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Icon(name IconName, args ...any)`|Returns a `cds-icon` view using a bundled Carbon icon name. Accepts `carbon.With(...)` size attrs plus standard `mvc` options.|

## Basic Usage

```go
search := carbon.Icon(
 carbon.IconSearch,
 carbon.With(carbon.IconSize20),
)
```

Icons can also be labeled when they need to be announced by assistive technology:

```go
warning := carbon.Icon(
 carbon.IconWarningFilled,
 carbon.With(carbon.IconSize24),
).SetLabel("Warning")
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Icon name|Pass any bundled `IconName` constant such as `IconSearch`, `IconSettings`, `IconLaunch`, `IconClose`, `IconUserAvatar`|
|Size|`IconSize16`, `IconSize20`, `IconSize24`, `IconSize32`|

Example:

```go
carbon.Icon(
 carbon.IconLaunch,
 carbon.With(carbon.IconSize32),
)
```

## State

|Method|Description|
|----|----|
|`Value() string`|Returns the currently assigned bundled icon name.|
|`SetValue(string)`|Updates the icon glyph using a bundled icon name.|
|`Label() string`|Returns the current accessible label (`aria-label`).|
|`SetLabel(string)`|Sets or clears the accessible label and tooltip text.|

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/icons/library/)
