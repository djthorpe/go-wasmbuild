---
description: HeaderPanel wraps Carbon's UI-shell right-side overlay panel and exposes a simple visible-state API through the shared MVC pattern.
---

# HeaderPanel

## Constructors

|Constructor|Description|
|----|----|
|`carbon.HeaderPanel(args ...any)`|Returns a `cds-header-panel` view.|

## Basic Usage

```go
panel := carbon.HeaderPanel(
 mvc.HTML("DIV",
  carbon.Head(4, "Notifications"),
  carbon.Para("Supplementary content lives in the panel."),
 ),
)

panel.SetVisible(true)
```

## Appearance

Carbon's header panel is intended as a right-side overlay attached to the UI shell. Width and positioning are usually controlled with normal `mvc.WithStyle(...)` options on the host element.

```go
panel.Root().SetAttribute("style", "position:absolute;top:0;right:0;height:100%;inline-size:24rem")
```

## State

|Method|Description|
|----|----|
|`Visible() bool`|Returns true when the panel is expanded.|
|`SetVisible(bool)`|Sets Carbon's `expanded` property and attribute.|

Because `HeaderPanel` implements `mvc.VisibleState`, helper components such as [CloseButton](CloseButton.md) can dismiss it automatically.

```go
panel := carbon.HeaderPanel(
 mvc.HTML("DIV",
  mvc.WithStyle("position:relative;padding:1rem 1.5rem"),
  carbon.CloseButton(mvc.WithStyle("position:absolute;top:0;right:0")),
  carbon.Head(4, "Panel"),
 ),
)
```

## Events

`HeaderPanel` does not register any custom component events.

## Notes

* `HeaderPanel` is a shell overlay, not a generic modal or drawer abstraction.
* Visibility is controlled through the `expanded` property, which the wrapper exposes as `Visible()` and `SetVisible(...)`.
* Layout and content are fully caller-defined through normal child views.

## References

* [CloseButton](CloseButton.md)
* [Carbon Design System](https://carbondesignsystem.com/components/UI-shell-right-panel/usage/)
