---
description: Pagination wraps Carbon's pagination component and exposes enabled and pagination-state helpers for page, limit, and total count.
---

# Pagination

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Pagination(args ...any)`|Returns a `cds-pagination` view.|

## State

|Method|Description|
|----|----|
|`Enabled() bool` / `SetEnabled(bool)`|Gets or sets the disabled state.|
|`Offset() uint` / `SetOffset(uint)`|Gets or sets the current item offset.|
|`Limit() uint` / `SetLimit(uint)`|Gets or sets the page size.|
|`Count() uint` / `SetCount(uint)`|Gets or sets the total item count.|
|`SetPage(uint)`|Sets the current page number and updates the derived offset.|
|`PagesUnknown() bool` / `SetPagesUnknown(bool)`|Gets or sets the pages-unknown mode.|

## Events

|Event|Description|
|----|----|
|`EventPaginationChanged`|Fires when the current page changes.|
|`EventPaginationPageSize`|Fires when the page size changes.|

## References

* [Carbon Design System](https://carbondesignsystem.com/components/pagination/usage/)
