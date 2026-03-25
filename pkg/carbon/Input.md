---
description: Input, SecureInput, and NumberInput provide Carbon text-entry fields with label, helper text, validation, and normalized event behavior.
---

# Input

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Input(args ...any)`|Returns a `cds-text-input` view.|
|`carbon.SecureInput(args ...any)`|Returns a `cds-password-input` view with the same public API as `Input`.|
|`carbon.NumberInput(args ...any)`|Returns a `cds-number-input` view for numeric entry.|

## Basic Usage

```go
name := carbon.Input().
 SetLabel("Project name").
 SetHelperText("Used in deployment dashboards.").
 SetInvalidText("Enter a project name.").
 SetRequired(true)

name.Root().SetAttribute("placeholder", "My service")
```

Shared `Input` and `SecureInput` methods:

|Method|Description|
|----|----|
|`Label() string` / `SetLabel(string)`|Gets or sets the field label.|
|`HelperText() string` / `SetHelperText(string)`|Gets or sets helper text.|
|`InvalidText() string` / `SetInvalidText(string)`|Gets or sets invalid-state text.|
|`Required() bool` / `SetRequired(bool)`|Gets or sets the required flag.|
|`Value() string` / `SetValue(string)`|Gets or sets the current field value.|
|`CheckValidity() bool`|Evaluates constraints and updates invalid UI.|
|`SetCustomValidity(string)`|Sets or clears a custom invalid message.|

`SecureInput` uses the same methods and event model as `Input`:

```go
password := carbon.SecureInput().
 SetLabel("Password").
 SetRequired(true)
```

Numeric input usage:

```go
retries := carbon.NumberInput().
 SetLabel("Retry count").
 SetValue("3").
 SetMin("0").
 SetMax("10").
 SetStep("1")
```

`NumberInput` supports the same label/helper/invalid/required/value helpers plus numeric-specific methods:

## Appearance

`Input` and `SecureInput` support the same Carbon field presentation attrs used in Storybook.

|Property|Values|
|----|----|
|Size|`SizeSmall`, `SizeMedium`, `SizeLarge`|
|Theme|`ThemeWhite`, `ThemeG10`, `ThemeG90`, `ThemeG100`|

`NumberInput` uses the same form-field presentation model, with min/max/step and stepper visibility controlled through state helpers.

## State

|Method|Description|
|----|----|
|`Label() string` / `SetLabel(string)`|Gets or sets the field label.|
|`HelperText() string` / `SetHelperText(string)`|Gets or sets helper text.|
|`InvalidText() string` / `SetInvalidText(string)`|Gets or sets invalid-state text.|
|`Required() bool` / `SetRequired(bool)`|Gets or sets the required flag.|
|`Value() string` / `SetValue(string)`|Gets or sets the current field value.|
|`CheckValidity() bool`|Evaluates constraints and updates invalid UI.|
|`SetCustomValidity(string)`|Sets or clears a custom invalid message.|
|`Min() string` / `SetMin(string)`|Gets or sets the minimum allowed value.|
|`Max() string` / `SetMax(string)`|Gets or sets the maximum allowed value.|
|`Step() string` / `SetStep(string)`|Gets or sets the step interval.|
|`AllowEmpty() bool` / `SetAllowEmpty(bool)`|Controls whether an empty value is valid.|
|`HideSteppers() bool` / `SetHideSteppers(bool)`|Shows or hides the stepper controls.|
|`CheckValidity() bool`|Evaluates numeric validity and updates invalid UI.|
|`SetCustomValidity(string)`|Sets or clears a custom invalid message.|

## Events

Text and secure inputs register the following normalized events:

|Event|Description|
|----|----|
|`EventInput`|Fires as the value changes.|
|`EventChange`|Bridged to fire on effective value changes when focus leaves the field.|
|`EventInvalid`|Fires when validity checking fails.|
|`EventFocus`|Fires when the field gains focus.|
|`EventNoFocus`|Fires when the field loses focus.|

`NumberInput` uses the same event names, with `EventInput` and `EventChange` normalized from Carbon's number-input change event.

## Notes

* `Input` and `SecureInput` keep an internal change baseline so `EventChange` only fires on actual value changes.
* `SetCustomValidity("")` clears the invalid UI.
* Placeholder text and non-wrapped attributes can be set directly on the root element with `Root().SetAttribute(...)`.

## References

* [Carbon Design System: Text input](https://carbondesignsystem.com/components/text-input/usage/)
* [Carbon Design System: Password input](https://carbondesignsystem.com/components/password-input/usage/)
* [Carbon Design System: Number input](https://carbondesignsystem.com/components/number-input/usage/)
