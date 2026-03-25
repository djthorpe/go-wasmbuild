---
description: Header renders Carbon's UI-shell header with name, nav body, and optional global actions slots.
---

# Header

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Header(args ...any)`|Returns a `cds-header` shell header.|

## State

|Method|Description|
|----|----|
|`Label() string`|Returns the header aria-label.|
|`SetLabel(href, prefix string, args ...any)`|Sets the header name slot using a `cds-header-name` link.|
|`Active() []mvc.View` / `SetActive(views ...mvc.View)`|Reads or updates active nav items.|
|`Item(href string)`|Returns the first matching nav item by href.|

## References

* [HeaderNavGlobal](HeaderNavGlobal.md)
* [HeaderNavItem](HeaderNavItem.md)
