---
description: TagGroup wraps one or more Carbon tags and coordinates group-level visible, enabled, and active state.
---

# TagGroup

## Constructors

|Constructor|Description|
|----|----|
|`carbon.TagGroup(args ...any)`|Returns a flex container for grouping tags.|

## State

|Method|Description|
|----|----|
|`Content(args ...any)`|Replaces the group's children with tags.|
|`Active() []mvc.View` / `SetActive(views ...mvc.View)`|Reads or updates selected child tags.|
|`Enabled() []mvc.View` / `SetEnabled(views ...mvc.View)`|Reads or updates enabled child tags.|
|`Visible() []mvc.View` / `SetVisible(views ...mvc.View)`|Reads or updates visible child tags.|

## Events

|Event|Description|
|----|----|
|`EventTagDismissibleClosed`|Fires when a dismissible child tag is closed.|
|`EventTagOperationalSelected`|Fires when an operational child tag is selected.|

## References

* [Tag](Tag.md)
* [DismissibleTag](DismissibleTag.md)
* [OperationalTag](OperationalTag.md)
