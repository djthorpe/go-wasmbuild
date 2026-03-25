---
description: Lead renders larger introductory Carbon body copy for summaries, openings, and prominent contextual text.
---

# Lead

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Lead(args ...any)`|Returns a `P` text view styled with Carbon's `cds--body-02` type token.|

## Basic Usage

```go
intro := carbon.Lead("Lead text works well for opening summaries and higher-emphasis explanatory content.")
```

`Lead` accepts the same child content and standard `mvc` options as the other text helpers:

```go
intro := carbon.Lead(
 mvc.WithClass("page-intro"),
 "Use lead copy when a paragraph needs more visual presence than standard body text.",
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
