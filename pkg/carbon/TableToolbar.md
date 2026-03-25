---
description: TableToolbar arranges Carbon table search and action controls, wrapping non-search content in a toolbar-content container.
---

# TableToolbar

## Constructors

|Constructor|Description|
|----|----|
|`carbon.TableToolbar(args ...any)`|Returns a `cds-table-toolbar` wrapper for table actions and search controls.|

## Basic Usage

```go
toolbar := carbon.TableToolbar(
 carbon.TableToolbarSearch(),
 carbon.Button("Add row"),
)
```

## Notes

* `TableToolbar` inserts search children directly.
* Other action content is wrapped automatically in `TableToolbarContent`.

## References

* [TableToolbarContent](TableToolbarContent.md)
* [TableToolbarSearch](TableToolbarSearch.md)
