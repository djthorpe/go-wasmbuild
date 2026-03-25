---
description: CheckboxGroup coordinates grouped Carbon checkboxes with shared helper text, legend text, orientation, and group-level active and enabled state.
---

# CheckboxGroup

## Constructors

|Constructor|Description|
|----|----|
|`carbon.CheckboxGroup(helperText string, args ...any)`|Returns a `cds-checkbox-group` view. The `helperText` argument becomes the `helper-text` attribute when non-empty.|

## Basic Usage

```go
marketing := carbon.Checkbox("Marketing emails").
 SetValue("marketing").
 SetActive(true)

group := carbon.CheckboxGroup("Choose all that apply.").
 Content(
  marketing,
  carbon.Checkbox("Product updates").SetValue("product"),
 ).
 SetLabel("Notifications")
```

## Appearance

Checkbox groups can be laid out vertically or horizontally:

|Property|Values|
|----|----|
|Orientation|`CheckboxOrientationVertical`, `CheckboxOrientationHorizontal`|

## State

|Method|Description|
|----|----|
|`Content(args ...any)`|Replaces the group's children with the provided checkboxes.|
|`Active() []mvc.View`|Returns the currently checked child checkboxes in the group.|
|`SetActive(views ...mvc.View)`|Checks the supplied group children and unchecks the rest.|
|`Enabled() []mvc.View`|Returns the currently enabled child checkboxes in the group.|
|`SetEnabled(views ...mvc.View)`|Enables the supplied group children and disables the rest.|
|`Label() string`|Returns the group's `legend-text` value.|
|`SetLabel(string)`|Sets the group's `legend-text` value.|
|`Orientation() CheckboxOrientation`|Returns the current group orientation.|
|`SetOrientation(CheckboxOrientation)`|Switches the group between vertical and horizontal layout.|

`CheckboxGroup` manages checkbox children only. `Content(args ...any)` panics if any child is not a `*checkbox`.

## Events

`CheckboxGroup` normalizes Carbon's checkbox change event to `EventChange`.

|Event|Description|
|----|----|
|`EventChange`|Fires when a child checkbox changes checked state.|

## Notes

* Group-level `EventChange` is useful because child checkbox events bubble.
* `CheckboxGroup.SetEnabled()` with no arguments disables every child checkbox.
* Use `Checkbox` methods when changing individual children, and use group methods when coordinating shared state.

## References

* [Checkbox](Checkbox.md)
* [Carbon Design System](https://carbondesignsystem.com/components/checkbox/usage/)