---
description: DismissibleTag renders Carbon's dismissible tag variant with closable visibility state.
---

# DismissibleTag

## Constructors

|Constructor|Description|
|----|----|
|`carbon.DismissibleTag(args ...any)`|Returns a `cds-dismissible-tag` view. A leading string becomes the `text` attribute.|

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Visible() bool` / `SetVisible(bool)`|Gets or sets the `open` state.|
|`Label() string` / `SetLabel(string)`|Gets or sets the visible label via the `text` attribute.|

## Events

|Event|Description|
|----|----|
|`EventTagDismissibleClosed`|Fires when the tag is dismissed.|

## References

* [Tag](Tag.md)
* [TagGroup](TagGroup.md)
