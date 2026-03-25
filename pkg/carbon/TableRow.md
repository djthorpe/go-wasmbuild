---
description: TableRow renders a Carbon table body row and wraps raw children in body cells.
---

# TableRow

## Constructors

|Constructor|Description|
|----|----|
|`carbon.TableRow(args ...any)`|Returns a Carbon table body row. Raw children are wrapped in `cds-table-cell` elements.|

## Basic Usage

```go
row := carbon.TableRow("Ada", "Maintainer", "Active")
```

## Notes

* Strings, DOM elements, and views are wrapped automatically in body cells.

## References

* [Table](Table.md)
* [TableHeader](TableHeader.md)
