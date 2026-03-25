---
description: Compact renders dense Carbon body copy for metadata, short descriptions, and layouts where vertical rhythm needs to stay tight.
---

# Compact

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Compact(args ...any)`|Returns a `P` text view styled with Carbon's `cds--body-compact-01` type token.|

## Basic Usage

```go
meta := carbon.Compact("Compact text is useful for dense interface regions and short supporting details.")
```

`Compact` accepts the same child content and standard `mvc` options as the other text helpers:

```go
meta := carbon.Compact(
 mvc.WithClass("metadata"),
 "Keep supporting details readable while conserving space.",
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
