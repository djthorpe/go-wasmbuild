---
description: OverflowMenu wraps Carbon's overflow-menu trigger and body, including visibility, size, flipped layout, and grouped item content.
---

# OverflowMenu

## Constructors

|Constructor|Description|
|----|----|
|`carbon.OverflowMenu(args ...any)`|Returns a `cds-overflow-menu` view.|

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Visible() bool` / `SetVisible(bool)`|Gets or sets the `open` state.|
|`Size() Attr` / `SetSize(Attr)`|Gets or sets the normalized menu size.|
|`SetFlipped(bool)`|Flips the menu placement.|
|`Label() string` / `SetLabel(string)`|Gets or sets the tooltip-content slot text.|
|`Content(args ...any)`|Routes menu items into the overflow body and leaves other host children in place.|

## References

* [OverflowMenuItem](OverflowMenuItem.md)
