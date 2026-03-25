---
description: Checkbox exposes boolean and tri-state selection through Carbon's checkbox component.
---

# Checkbox

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Checkbox(args ...any)`|Returns a `cds-checkbox` view. A leading string becomes the `label-text` attribute.|

## Basic Usage

```go
marketing := carbon.Checkbox("Marketing emails").
 SetValue("marketing").
 SetActive(true)
```

## Appearance

Checkbox appearance is controlled by its own label, value, active state, and enabled state.

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

Tri-state checkboxes use `CheckboxStateUndefined` for the indeterminate state:

```go
consent := carbon.Checkbox("Partial consent")
consent.SetState(carbon.CheckboxStateUndefined)
```

## Events

`Checkbox` normalizes Carbon's checkbox change event to `EventChange`.

|Event|Description|
|----|----|
|`EventChange`|Fires when the checked state changes.|

## Notes

* `Checkbox.Active()` reports only the checked state. Use `State()` when indeterminate matters.
* Use `CheckboxGroup` for coordinated group label, helper text, orientation, and child state.

## References

* [CheckboxGroup](CheckboxGroup.md)
* [Carbon Design System](https://carbondesignsystem.com/components/checkbox/usage/)
