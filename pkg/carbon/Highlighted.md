---
description: Highlighted renders inline highlighted text using the semantic HTML mark element.
---

# Highlighted

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Highlighted(args ...any)`|Returns a `MARK` text view for highlighted inline content.|

## Basic Usage

```go
copy := carbon.Para(
 "Status: ",
 carbon.Highlighted("needs review"),
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
