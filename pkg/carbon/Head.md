---
description: Head renders Carbon-styled headings from h1 through h6, mapping semantic HTML heading levels onto Carbon's type scale.
---

# Head

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Head(level int, args ...any)`|Returns an `H1` through `H6` text view styled with the matching Carbon heading class.|

## Basic Usage

```go
title := carbon.Head(1, "Page title")
section := carbon.Head(2, "Deployment status")
detail := carbon.Head(4, "Build metadata")
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
