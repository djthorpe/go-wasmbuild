---
description: Table renders Carbon's basic data table wrapper with header and body slots.
---

# Table

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Table(args ...any)`|Returns a minimal Carbon table wrapper with header and body slots.|

## Basic Usage

```go
table := carbon.Table(
 carbon.TableHeader("Name", "Role"),
 carbon.TableRow("Ada", "Maintainer"),
)
```

## Notes

* Use `TableHeader` for header rows and `TableRow` for body rows.
* `Content(...)` writes into the body slot by default.

## References

* [TableHeader](TableHeader.md)
* [TableRow](TableRow.md)
* [Carbon Design System](https://carbondesignsystem.com/components/data-table/usage/)
