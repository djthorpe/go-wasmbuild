---
description: Tag, DismissibleTag, OperationalTag, and TagGroup wrap Carbon's metadata and status labels, including icon slots, selection, dismissal, and group-level observation.
---

# Tag

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Tag(args ...any)`|Returns a `cds-tag` view.|
|`carbon.DismissibleTag(args ...any)`|Returns a `cds-dismissible-tag` view. A leading string becomes the `text` attribute.|
|`carbon.OperationalTag(args ...any)`|Returns a `cds-operational-tag` view. A leading string becomes the `text` attribute.|
|`carbon.TagGroup(args ...any)`|Returns a flex container for grouping tags.|

## Basic Usage

```go
status := carbon.Tag(
 carbon.Icon(carbon.IconWarningFilled),
 "Review",
 carbon.With(carbon.TagYellow, carbon.SizeMedium),
)

dismissible := carbon.DismissibleTag(
 "Beta",
 carbon.Icon(carbon.IconLaunch),
 carbon.With(carbon.TagBlue),
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
|`Visible() bool` / `SetVisible(bool)`|Gets or sets the `open` state used by dismissible and operational tags.|
|`Active() bool` / `SetActive(bool)`|Gets or sets the `selected` state used by operational tags.|
|`Label() string` / `SetLabel(string)`|Gets or sets the visible label. Plain tags use child content; dismissible and operational tags use the `text` attribute.|
|`Content(args ...any)`|Replaces the group's children.|
|`Active() []mvc.View` / `SetActive(views ...mvc.View)`|Reads or updates selected child tags.|
|`Enabled() []mvc.View` / `SetEnabled(views ...mvc.View)`|Reads or updates enabled child tags.|
|`Visible() []mvc.View` / `SetVisible(views ...mvc.View)`|Reads or updates visible child tags.|

`TagGroup` accepts only `*tag` children. `Content(args ...any)` panics for other types.

## Events

`TagGroup` observes bubbled events from its children.

|Event|Description|
|----|----|
|`EventTagDismissibleClosed`|Fires when a dismissible tag is closed.|
|`EventTagOperationalSelected`|Fires when an operational tag is selected.|

## Notes

* `TagGroup` is a plain `DIV` wrapper with Carbon-friendly flex layout defaults.
* Plain tags store visible text in child content; dismissible and operational tags store it in the `text` attribute.
* Decorative tag icons inherit `currentColor` automatically.

## References

* [Carbon Design System](https://carbondesignsystem.com/components/tag/usage/)
