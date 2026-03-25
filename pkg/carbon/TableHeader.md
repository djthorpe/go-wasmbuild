---
description: TableHeader renders a Carbon table header row and wraps raw children in header cells.
---

# TableHeader

## Constructors

|Constructor|Description|
|----|----|
|`carbon.TableHeader(args ...any)`|Returns a Carbon table header row. Raw children are wrapped in `cds-table-header-cell` elements.|

## Basic Usage

```go
header := carbon.TableHeader("Name", "Role", "Status")
```

## Notes

* Strings, DOM elements, and views are wrapped automatically in header cells.

## References

* [Table](Table.md)
* [TableRow](TableRow.md)
