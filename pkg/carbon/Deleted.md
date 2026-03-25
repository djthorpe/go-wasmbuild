---
description: Deleted renders inline deleted text using the semantic HTML del element.
---

# Deleted

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Deleted(args ...any)`|Returns a `DEL` text view for deleted or removed inline content.|

## Basic Usage

```go
copy := carbon.Para(
 "Previous wording: ",
 carbon.Deleted("deprecated"),
)
```

## References

* [Carbon Design System](https://carbondesignsystem.com/elements/typography/type-sets/)
