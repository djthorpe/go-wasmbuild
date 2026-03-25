---
description: Code renders a short inline Carbon code snippet for commands, identifiers, and brief source fragments within body copy.
---

# Code

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Code(args ...any)`|Returns a `cds-code-snippet` with `type="inline"` for short code fragments embedded in surrounding text.|

## Basic Usage

```go
cmd := carbon.Para(
 "Run ",
 carbon.Code("go test ./..."),
 " before publishing changes.",
)
```

Inline code is intended for short tokens such as commands, file names, environment variables, and compact expressions.

## Appearance

|Property|`With` and `Apply` values|
|----|----|
|Variant|Always `type="inline"` via `carbon.Code(...)`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100` on the surrounding container|

Example:

```go
carbon.Code(
 carbon.With(carbon.ThemeG10),
 "go test ./...",
)
```

Inline code does not expose additional code-snippet-specific appearance helpers.

## References

* [CodeBlock](CodeBlock.md)
* [Carbon Design System](https://carbondesignsystem.com/components/code-snippet/usage/)
