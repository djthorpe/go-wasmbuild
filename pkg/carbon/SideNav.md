---
description: SideNav renders Carbon's side navigation shell and coordinates active state across nested nav items.
---

# SideNav

## Constructors

|Constructor|Description|
|----|----|
|`carbon.SideNav(args ...any)`|Returns a `cds-side-nav` shell panel.|

## State

|Method|Description|
|----|----|
|`Active() []mvc.View` / `SetActive(views ...mvc.View)`|Reads or updates active nav items.|
|`Item(href string)`|Returns the first matching nav item by href.|

## Events

|Event|Description|
|----|----|
|`EventSectionToggling`|Fires while a side-nav section is being toggled.|
|`EventSectionToggle`|Fires when a side-nav section finishes toggling.|

## References

* [SideNavLink](SideNavLink.md)
* [SideNavGroup](SideNavGroup.md)
* [SideNavGroupItem](SideNavGroupItem.md)
