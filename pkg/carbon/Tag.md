---
description: Tag wraps Carbon's plain metadata and status label component, including icon slot normalization and basic enabled state.
---

# Tag

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Tag(args ...any)`|Returns a `cds-tag` view.|

## Basic Usage

```go
status := carbon.Tag(
 carbon.Icon(carbon.IconWarningFilled),
 "Review",
 carbon.With(carbon.TagBlue, carbon.SizeMedium),
)
```

## Appearance

Common tag attrs:

|Property|Values|
|----|----|
|Type|`TagRed`, `TagMagenta`, `TagPurple`, `TagBlue`, `TagCyan`, `TagTeal`, `TagGreen`, `TagGray`, `TagCoolGray`, `TagWarmGray`, `TagHighContrast`, `TagOutline`|
|Size|`SizeSmall`, `SizeMedium`, `SizeLarge`|

Icons passed to any tag variant are normalized into Carbon's `icon` slot and default to `aria-hidden="true"` unless a label is set on the icon.

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Label() string` / `SetLabel(string)`|Gets or sets the visible label using child content.|

## Notes

* Plain tags store visible text in child content.
* Decorative tag icons inherit `currentColor` automatically.

## References

* [DismissibleTag](DismissibleTag.md)
* [OperationalTag](OperationalTag.md)
* [TagGroup](TagGroup.md)
* [Carbon Design System](https://carbondesignsystem.com/components/tag/usage/)
