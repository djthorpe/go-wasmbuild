---
description: Buttons are used to initialize an action. Button labels express what action will occur when the user interacts with it.
---

#### Constructors

|Constructor|Description|
|----|----|
|`carbon.Button(args ...any)`|Returns a `cds-button` view. Accepts content, `carbon.With(...)` options, and standard `mvc` options.|

Related components:

* [CloseButton](CloseButton.md)
* [ButtonGroup](ButtonGroup.md)

#### Basic usage

```go
save := carbon.Button(
 carbon.With(carbon.KindPrimary, carbon.SizeLarge),
 "Save",
)

save.AddEventListener(carbon.EventClick, func(e dom.Event) {
 // handle save
})
```

Buttons can contain text, an icon, or both.

```go
next := carbon.Button(
 carbon.With(carbon.KindSecondary),
 "Next",
 carbon.Icon(carbon.IconArrowRight, carbon.With(carbon.IconSize20)),
)
```

#### Icon-only buttons

If a button contains only an icon and no text, choose the button kind explicitly to match the emphasis you want.

Icon-only buttons should have an accessible label:

```go
settings := carbon.Button(
 carbon.With(carbon.KindGhost),
 carbon.Icon(carbon.IconSettings, carbon.With(carbon.IconSize20)),
).SetLabel("Settings")
```

`SetLabel` updates both `aria-label` and the Carbon tooltip text.

#### Appearance

|Property|`With` and `Apply` values|
|----|----|
|Variant|`KindPrimary`, `KindSecondary`, `KindTertiary`, `KindGhost`, `KindDanger`, `KindDangerTertiary`, `KindDangerGhost`|
|Size|`SizeExtraSmall`, `SizeSmall`, `SizeMedium`, `SizeLarge`, `SizeExtraLarge`, `Size2XLarge`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

Example:

```go
carbon.Button(
 carbon.With(carbon.KindDangerGhost, carbon.SizeSmall),
 "Delete",
)
```

Themes are applied with the same `carbon.With(...)` helper:

```go
carbon.Button(
 carbon.With(carbon.KindPrimary, carbon.SizeExtraLarge, carbon.ThemeG90),
 "Deploy",
)
```

#### State and values

Buttons support enabled, value, and label helpers:

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the button is not disabled.|
|`SetEnabled(bool)`|Adds or removes the `disabled` attribute.|
|`Value() string`|Returns the `value` attribute.|
|`SetValue(string)`|Sets the `value` attribute for use by event handlers.|
|`Label() string`|Returns the current accessible label (`aria-label`).|
|`SetLabel(string)`|Sets the accessible label and tooltip text.|

```go
run := carbon.Button("Run").SetValue("build")

run.AddEventListener(carbon.EventClick, func(e dom.Event) {
  view := mvc.ViewFromEvent(e)
  if view != nil {
    _ = view.Value()
  }
})
```

#### Events

|Event|Description|
|----|----|
|EventClick|User clicks the button.|
|EventHover|Pointer enters the button.|
|EventNoHover|Pointer leaves the button.|
|EventFocus|Button gains focus.|
|EventNoFocus|Button loses focus.|

#### Notes

* Icons added to a button are automatically placed into the `icon` slot.
* Button icons default to `aria-hidden="true"` unless you assign them an accessible label.
* Icon-only buttons should set their kind explicitly with `carbon.With(...)`.

#### References

* [Carbon Design System](https://carbondesignsystem.com/components/button/usage/)
