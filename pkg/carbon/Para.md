---
description: Para renders standard Carbon body copy for readable paragraphs, descriptions, and general supporting text.
---

# Para

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Para(args ...any)`|Returns a `P` text view styled with Carbon's `cds--body-01` type token.|

## Basic Usage

```go
body := carbon.Para("Paragraph text supports the primary task without competing with headings or controls.")
```

`Para` accepts the same child content and standard `mvc` options as the other text helpers:

```go
body := carbon.Para(
 mvc.WithID("summary"),
 "Body copy can include plain strings and other inline views.",
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
