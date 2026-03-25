---
description: Dropdowns wrap Carbon's single-select menu field with label, helper text, selected-item control, and theme-aware presentation tweaks.
---

# Dropdown

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Dropdown(helperText string, args ...any)`|Returns a `cds-dropdown` view. Non-empty `helperText` sets the `helper-text` attribute.|

## Basic Usage

```go
draft := carbon.DropdownItem("Draft").SetValue("draft")
review := carbon.DropdownItem("Review").SetValue("review")

status := carbon.Dropdown("Used to control publish flow.", draft, review).
 SetLabel("Status").
 SetActive(review)
```

## Appearance

`Dropdown.Apply(carbon.With(...))` supports the normal Carbon size and theme attrs.

|Property|Values|
|----|----|
|Size|`SizeSmall`, `SizeMedium`, `SizeLarge`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

The wrapper also injects a small white-theme style shim when the dropdown is themed with `ThemeWhite` and no explicit field/layer surface tokens are already set.

## State

|Method|Description|
|----|----|
|`Enabled() bool`|Returns true when the dropdown is not disabled.|
|`SetEnabled(bool)`|Enables or disables the field.|
|`Label() string`|Returns the current title text.|
|`SetLabel(string)`|Writes the label into Carbon's `title-text` slot.|
|`Value() string`|Returns the current selected value.|
|`SetValue(string)`|Sets the selected value on the host element.|
|`Active() []mvc.View`|Returns the selected item views.|
|`SetActive(views ...mvc.View)`|Marks the supplied items selected and clears the rest.|
|`Content(args ...any)`|Replaces the dropdown items. Panics if any child is not a `*dropdownItem`.|
|`Value() string`|Returns the item's value.|
|`SetValue(string)`|Sets the item's value.|
|`Active() bool`|Returns true when the item is marked selected.|
|`SetActive(bool)`|Adds or removes the `selected` attribute.|

## Events

|Event|Description|
|----|----|
|`EventSelected`|Fires when the user selects an item.|

Typical event handling reads back the selected value from the dropdown view:

```go
status.AddEventListener(carbon.EventSelected, func(e dom.Event) {
 if view := mvc.ViewFromEventTarget(e, carbon.ViewDropdown); view != nil {
  _ = view.Value()
 }
})
```

## Notes

* `SetActive(...)` updates both the selected item and the host dropdown value.
* `SetLabel(...)` manages the `title-text` slot content directly.
* Use `mvc.WithStyle("width:100%")` when the field should stretch to its container width.

## References

* [DropdownItem](DropdownItem.md)
* [Carbon Design System](https://carbondesignsystem.com/components/dropdown/usage/)
