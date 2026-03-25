---
description: Checkboxes expose boolean and tri-state selection, and checkbox groups coordinate shared label, helper text, and group-level state.
---

# Checkbox

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Checkbox(args ...any)`|Returns a `cds-checkbox` view. A leading string becomes the `label-text` attribute.|
|`carbon.CheckboxGroup(helperText string, args ...any)`|Returns a `cds-checkbox-group` view. The `helperText` argument becomes the `helper-text` attribute when non-empty.|

## Basic Usage

```go
marketing := carbon.Checkbox("Marketing emails").
 SetValue("marketing").
 SetActive(true)

group := carbon.CheckboxGroup("Choose all that apply.",
 marketing,
 carbon.Checkbox("Product updates").SetValue("product"),
).SetLabel("Notifications")
```

## Appearance

Checkbox groups can be laid out vertically or horizontally:

|Property|Values|
|----|----|
|Orientation|`CheckboxOrientationVertical`, `CheckboxOrientationHorizontal`|

## State

|Method|Description|
|----|----|
|`Active() bool`|Returns true when the checkbox is checked.|
|`SetActive(bool)`|Checks or unchecks the checkbox.|
|`State() CheckboxState`|Returns `CheckboxStateTrue`, `CheckboxStateFalse`, or `CheckboxStateUndefined`.|
|`SetState(CheckboxState)`|Sets checked, unchecked, or indeterminate state.|
|`Enabled() bool`|Returns true when the checkbox is not disabled.|
|`SetEnabled(bool)`|Enables or disables the checkbox.|
|`Label() string`|Returns the `label-text` attribute.|
|`SetLabel(string)`|Sets the `label-text` attribute.|
|`Value() string`|Returns the current `value`.|
|`SetValue(string)`|Sets the `value` property and attribute.|
|`Content(args ...any)`|Replaces a checkbox group's children with the provided checkboxes.|
|`Active() []mvc.View`|Returns the currently checked child checkboxes in a group.|
|`SetActive(views ...mvc.View)`|Checks the supplied group children and unchecks the rest.|
|`Enabled() []mvc.View`|Returns the currently enabled child checkboxes in a group.|
|`SetEnabled(views ...mvc.View)`|Enables the supplied group children and disables the rest.|
|`SetOrientation(CheckboxOrientation)`|Switches a checkbox group between vertical and horizontal layout.|

Tri-state checkboxes use `CheckboxStateUndefined` for the indeterminate state:

```go
consent := carbon.Checkbox("Partial consent")
consent.SetState(carbon.CheckboxStateUndefined)
```

`CheckboxGroup` manages checkbox children only. `Content(args ...any)` panics if any child is not a `*checkbox`.

## Events

Both `Checkbox` and `CheckboxGroup` normalize Carbon's checkbox change event to `EventChange`.

|Event|Description|
|----|----|
|`EventChange`|Fires when the checked state changes.|

## Notes

* `Checkbox.Active()` reports only the checked state. Use `State()` when indeterminate matters.
* Group-level `EventChange` is useful because child checkbox events bubble.
* `CheckboxGroup.SetEnabled()` with no arguments disables every child checkbox.

## References

* [Carbon Design System](https://carbondesignsystem.com/components/checkbox/usage/)
