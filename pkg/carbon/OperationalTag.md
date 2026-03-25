---
description: OperationalTag renders Carbon's selectable operational tag variant with active and visible state helpers.
---

# OperationalTag

## Constructors

|Constructor|Description|
|----|----|
|`carbon.OperationalTag(args ...any)`|Returns a `cds-operational-tag` view. A leading string becomes the `text` attribute.|

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Active() bool` / `SetActive(bool)`|Gets or sets the `selected` state.|
|`Visible() bool` / `SetVisible(bool)`|Gets or sets the `open` state.|
|`Label() string` / `SetLabel(string)`|Gets or sets the visible label via the `text` attribute.|

## Events

|Event|Description|
|----|----|
|`EventTagOperationalSelected`|Fires when the tag is selected.|

## References

* [Tag](Tag.md)
* [TagGroup](TagGroup.md)
