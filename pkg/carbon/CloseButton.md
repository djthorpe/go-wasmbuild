---
description: CloseButton creates a ready-made icon-only Carbon button for dismiss actions.
---

# CloseButton

## Constructors

|Constructor|Description|
|----|----|
|`carbon.CloseButton(args ...any)`|Returns a ghost icon-only close button and forwards any additional args to `carbon.Button(...)`.|

## Basic Usage

```go
modal := carbon.Modal(
 carbon.CloseButton(),
)
```

You can still apply normal button options:

```go
dismiss := carbon.CloseButton(
 mvc.WithID("dismiss"),
 mvc.WithStyle("margin-left:auto"),
)
```

If you need a different kind, pass it explicitly and it will override the default:

```go
dismiss := carbon.CloseButton(
 carbon.With(carbon.KindTertiary),
)
```

## Appearance

`CloseButton` defaults to `carbon.KindGhost` and renders the `IconClose` icon at `IconSize20`.

If you need a different kind, pass it explicitly and it will override the default:

```go
dismiss := carbon.CloseButton(
 carbon.With(carbon.KindTertiary),
)
```

## State

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the button is not disabled.|
|`SetEnabled(bool)`|Enables or disables the button.|
|`Value() string`|Returns the current value attribute.|
|`SetValue(string)`|Sets the `value` attribute.|
|`Label() string`|Returns the current accessible label.|
|`SetLabel(string)`|Overrides the default accessible label and tooltip text.|

## Events

`CloseButton` uses the same event model as `Button`:

|Event|Description|
|----|----|
|`EventClick`|User clicks the button.|
|`EventHover`|Pointer enters the button.|
|`EventNoHover`|Pointer leaves the button.|
|`EventFocus`|Button gains focus.|
|`EventNoFocus`|Button loses focus.|

## Notes

* When clicked, `CloseButton` walks up the parent view chain and looks for the nearest ancestor implementing `mvc.VisibleState`.
* If one is found, it calls `SetVisible(false)` on that ancestor.
* This makes it useful for dismissing views such as modals, toasts, panels, or other hideable containers.
* If no ancestor implements `mvc.VisibleState`, the click handler does nothing beyond the normal button event.

## References

* [Button](Button.md)
* [Carbon Design System](https://carbondesignsystem.com/components/button/usage/)
