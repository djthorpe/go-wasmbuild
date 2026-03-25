---
description: TableToolbarSearch renders Carbon's table toolbar search field with label, placeholder, enabled, value, and expanded state helpers.
---

# TableToolbarSearch

## Constructors

|Constructor|Description|
|----|----|
|`carbon.TableToolbarSearch(args ...any)`|Returns a `cds-table-toolbar-search` view. Defaults the label to `Search` when unset.|

## State

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the search control is enabled.|
|`SetEnabled(bool)`|Enables or disables the control.|
|`Value() string`|Returns the current search value.|
|`SetValue(string)`|Sets the current search value.|
|`Label() string`|Returns the current label text.|
|`SetLabel(string)`|Sets the label text.|
|`Placeholder() string`|Returns the placeholder text.|
|`SetPlaceholder(string)`|Sets the placeholder text.|
|`Expanded() bool`|Returns whether the search is expanded.|
|`SetExpanded(bool)`|Expands or collapses the search control.|

## Events

|Event|Description|
|----|----|
|`EventInput`|Fires from Carbon's search input event bridge.|
|`EventChange`|Fires when the value changes.|
|`EventFocus`|Fires when the control receives focus.|
|`EventNoFocus`|Fires when the control loses focus.|

## References

* [TableToolbar](TableToolbar.md)
