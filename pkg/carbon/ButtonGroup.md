---
description: ButtonGroup arranges related Carbon buttons horizontally with the correct group spacing.
---

### Constructor

|Constructor|Description|
|----|----|
|`carbon.ButtonGroup(args ...any)`|Returns a `cds-button-group` view for grouping buttons with Carbon spacing.|

The group constructor accepts standard `mvc` options on the group root, such as `mvc.WithID(...)`, `mvc.WithClass(...)`, `mvc.WithStyle(...)`, and attributes like `mvc.WithAttr("aria-label", ...)`.

Button kind and size should be applied to the child buttons rather than the group root.

### Appearance

|Property|`With` and `Apply` values|
|----|----|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

Themes apply to the group root:

```go
group := carbon.ButtonGroup(
 carbon.With(carbon.ThemeG90),
 mvc.WithAttr("aria-label", "Primary actions"),
)
```

Button appearance still belongs on the children:

```go
group := carbon.ButtonGroup(
 carbon.With(carbon.ThemeG10),
).Content(
 carbon.Button(carbon.With(carbon.KindSecondary, carbon.SizeSmall), "Back"),
 carbon.Button(carbon.With(carbon.KindPrimary, carbon.SizeLarge), "Continue"),
)
```

### Basic usage

```go
back := carbon.Button(carbon.With(carbon.KindTertiary), "Back")
next := carbon.Button(carbon.With(carbon.KindPrimary), "Next")

group := carbon.ButtonGroup().Content(back, next)
```

If the group needs a label or layout styling, apply that on the group root and keep button appearance on the children:

```go
group := carbon.ButtonGroup(
 mvc.WithAttr("aria-label", "Wizard actions"),
 mvc.WithStyle("margin-top:1rem"),
).Content(
 carbon.Button(carbon.With(carbon.KindSecondary, carbon.SizeLarge), "Back"),
 carbon.Button(carbon.With(carbon.KindPrimary, carbon.SizeLarge), "Next"),
)
```

### Content

`ButtonGroup.Content(args ...any)` appends buttons to the group.

All children must be `*button` values. Passing another type will panic.

```go
cancel := carbon.Button(carbon.With(carbon.KindGhost), "Cancel")
save := carbon.Button(carbon.With(carbon.KindPrimary), "Save")

group := carbon.ButtonGroup().Content(cancel, save)
```

### Group state

`ButtonGroup` can manage enabled state across all of its children:

|Method|Description|
|----|----|
|`Enabled() []mvc.View`|Returns the currently enabled buttons.|
|`SetEnabled(views ...mvc.View)`|Enables the supplied buttons and disables the rest.|

```go
group.SetEnabled(back, next)
```

With no arguments, `SetEnabled()` disables all buttons.

### Events

`ButtonGroup` accepts the same button-style listeners as `Button`.

|Event|Description|
|----|----|
|EventClick|Triggered when the group receives a click event.|
|EventHover|Triggered when the group receives a hover event.|
|EventNoHover|Triggered when the group receives a pointer-leave event.|
|EventFocus|Triggered when the group receives a focus event.|
|EventNoFocus|Triggered when the group receives a focus-out event.|

When the event target is a child view, use `mvc.ViewFromEvent(event, carbon.ViewButton)` to resolve the originating button.

### Notes

* `ButtonGroup` is a layout and enabled-state wrapper around `Button` instances.
* Themes and general `mvc` options can be applied to the group root.
* Kind and size belong on the child buttons.
* Use individual button methods when only one child needs changing.
* Use group methods when enabled state should be coordinated.

### References

* [Button](Button.md)
* [Carbon Design System](https://carbondesignsystem.com/components/button/usage/)
