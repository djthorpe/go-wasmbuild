---
description: OverflowMenuItem wraps Carbon's selectable overflow-menu item with value, label, size, danger, divider, and enabled-state helpers.
---

# OverflowMenuItem

## Constructors

|Constructor|Description|
|----|----|
|`carbon.OverflowMenuItem(args ...any)`|Returns a `cds-overflow-menu-item` view.|

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Value() string` / `SetValue(string)`|Gets or sets the item value.|
|`SetLabel(string)`|Sets the item label HTML.|
|`SetDanger(bool)`|Marks the item as dangerous.|
|`SetDivider(bool)`|Adds or removes a divider.|
|`Size() Attr` / `SetSize(Attr)`|Gets or sets the normalized item size.|

## References

* [OverflowMenu](OverflowMenu.md)
