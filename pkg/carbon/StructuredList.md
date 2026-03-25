---
description: StructuredList provides Carbon's table-like structured list for read-only or single-select grouped content.
---

# StructuredList

## Constructors

|Constructor|Description|
|----|----|
|`carbon.StructuredList(args ...any)`|Returns a Carbon structured list wrapper with header and body slots.|
|`carbon.StructuredListHeader(args ...any)`|Returns a structured list header row. Raw children are wrapped in header cells.|
|`carbon.StructuredListRow(args ...any)`|Returns a structured list body row. Raw children are wrapped in body cells.|
|`carbon.StructuredListHeaderCell(args ...any)`|Returns an explicit structured list header cell.|
|`carbon.StructuredListCell(args ...any)`|Returns an explicit structured list body cell.|

## Basic Usage

```go
plans := carbon.StructuredList(
 carbon.StructuredListHeader("Plan", "Support", "Notes"),
 carbon.StructuredListRow("Starter", "Community", "Good for quick evaluation."),
 carbon.StructuredListRow("Team", "Business hours", "Balanced option for internal tools."),
)
```

Selectable structured lists use a selection name on the wrapper and selection values on rows:

```go
team := carbon.StructuredListRow("Team", "Business hours", "Balanced option for internal tools.")
team.SetValue("team")
starter := carbon.StructuredListRow("Starter", "Community", "Good for quick evaluation.")
starter.SetValue("starter")

plans := carbon.StructuredList(
 mvc.WithAttr("selection-name", "pricing-plan"),
 carbon.StructuredListHeader("Plan", "Support", "Notes"),
 starter,
 team,
)

plans.SetActive(team)
```

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Size|`StructuredListCondensed`|
|Alignment|`StructuredListFlush`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

## State

|Method|Description|
|----|----|
|`list.Active() []mvc.View`|Returns the currently selected rows. In the selectable variant this is at most one row.|
|`list.SetActive(rows ...mvc.View)`|Selects the first supplied row and deselects the rest. Calling it with no rows clears the selection.|
|`row.Active() bool`|Returns true when a structured list row is selected.|
|`row.SetActive(bool)`|Selects or deselects a structured list row.|
|`row.Value() string`|Returns the row selection value.|
|`row.SetValue(string)`|Sets or clears the row selection value.|

## Events

|Event|Description|
|----|----|
|`EventChange`|Fires on the structured list when the selected row changes in the selectable variant.|

Use low-level DOM `click`, `focus`, or keyboard listeners on rows only if you need interaction details beyond selection changes.

## Notes

* `StructuredListFlush` is for the flush alignment. The default hang alignment is used when `flush` is not applied.
* `StructuredList` implements `mvc.ActiveGroup`; the selectable variant behaves more like a radio group than a checkbox group.
* Carbon's selectable variant supports a single selected row at a time.
* Flush alignment is not recommended for the selectable variant in Carbon's guidance.

## References

* [List](List.md)
* [Carbon Design System](https://carbondesignsystem.com/components/structured-list/usage/)
