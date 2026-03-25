---
description: Tile and TileDecorator provide Carbon's simple content surface, with optional fill, height, background, and decorator-slot presentation helpers.
---

# Tile

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Tile(args ...any)`|Returns a `cds-tile` view.|
|`carbon.TileDecorator(args ...any)`|Alias for `Tile(...)`, kept for backward compatibility. Decorator treatment comes from slotted child content, not a different host element.|

## Basic Usage

```go
summary := carbon.Tile(
 carbon.Head(4, "Summary"),
 carbon.Para("Use tiles for quiet structure and grouping."),
)
```

## Appearance

Carbon tiles support a top-right `decorator` slot. The wrapper keeps this as plain child content:

```go
tile := carbon.Tile(
 mvc.HTML("SPAN", mvc.WithAttr("slot", "decorator"), "AI"),
 carbon.Head(4, "Generated result"),
 carbon.Compact("Decorator content is optional and caller-defined."),
)
```

`Tile.Apply(...)` re-computes inline presentation styles from a few helper opts.

|Helper|Effect|
|----|----|
|`carbon.WithFill()`|Adds `width:100%` and `display:block`.|
|`carbon.WithHeight(string)`|Sets a fixed height.|
|`carbon.WithBackground(string)`|Sets the Carbon layer token used for the tile background.|

Example:

```go
metric := carbon.Tile(
 carbon.WithFill(),
 carbon.WithHeight("9rem"),
 carbon.WithBackground("var(--cds-layer-02,#e0e0e0)"),
 carbon.Head(3, "42"),
)
```

## State

`Tile` does not expose component-specific state helpers beyond normal content and attribute updates on the underlying view.

## Events

`Tile` does not register any custom component events.

## Notes

* `Tile` is intentionally static; it does not add a selection model or custom events.
* Generated inline presentation styles are merged with any caller-provided `style` attribute.
* `TileDecorator(...)` exists only for API continuity; new code can use `Tile(...)` directly.

## References

* [Carbon Design System](https://carbondesignsystem.com/components/tile/usage/)
