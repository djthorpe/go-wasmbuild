---
description: Grid documents Carbon's CSS grid container variants and the column helpers used to build 16-column layouts.
---

# Grid

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Grid(args ...any)`|Returns a standard Carbon 16-column CSS grid container.|
|`carbon.GridFullWidth(args ...any)`|Returns a full-width grid container using Carbon's full-width grid class.|
|`carbon.GridNarrow(args ...any)`|Returns a narrow-gutter grid container.|
|`carbon.GridCondensed(args ...any)`|Returns a condensed-gutter grid container.|
|`carbon.Col1(args ...any)`|Returns a column spanning 1 of 16 columns.|
|`carbon.Col2(args ...any)`|Returns a column spanning 2 of 16 columns.|
|`carbon.Col4(args ...any)`|Returns a column spanning 4 of 16 columns.|
|`carbon.Col6(args ...any)`|Returns a column spanning 6 of 16 columns.|
|`carbon.Col8(args ...any)`|Returns a column spanning 8 of 16 columns.|
|`carbon.Col10(args ...any)`|Returns a column spanning 10 of 16 columns.|
|`carbon.Col12(args ...any)`|Returns a column spanning 12 of 16 columns.|
|`carbon.Col16(args ...any)`|Returns a column spanning all 16 columns.|
|`carbon.Col(n int, args ...any)`|Returns a column spanning `n` columns, where `n` must be between 1 and 16.|

## Basic Usage

```go
layout := carbon.Grid(
 carbon.Col4(carbon.Compact("Sidebar")),
 carbon.Col12(carbon.Para("Primary content")),
)
```

Grid variants use the same column helpers:

```go
wide := carbon.GridFullWidth(
 carbon.Col8(carbon.Para("Main")),
 carbon.Col8(carbon.Para("Secondary")),
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/2x-grid/overview/)
