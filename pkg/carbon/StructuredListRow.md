---
description: StructuredListRow renders a body row for a Carbon structured list and exposes active and value state for selectable lists.
---

# StructuredListRow

## Constructors

|Constructor|Description|
|----|----|
|`carbon.StructuredListRow(args ...any)`|Returns a structured list body row. Raw children are wrapped in `StructuredListCell` views.|

## State

|Method|Description|
|----|----|
|`Active() bool`|Returns true when the row is selected.|
|`SetActive(bool)`|Selects or deselects the row.|
|`Value() string`|Returns the row selection value.|
|`SetValue(string)`|Sets the row selection value.|

## References

* [StructuredList](StructuredList.md)
* [StructuredListCell](StructuredListCell.md)
