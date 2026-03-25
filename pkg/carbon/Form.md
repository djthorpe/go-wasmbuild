---
description: Form wraps Carbon's form container and normalizes bubbling input, change, invalid, and focus events across descendant controls.
---

# Form

## Constructors

|Constructor|Description|
|----|----|
|`carbon.Form(args ...any)`|Returns a `cds-form` view.|

## Events

|Event|Description|
|----|----|
|`EventInput`|Observed from descendant form controls.|
|`EventChange`|Observed from descendant form controls, including normalized checkbox and number-input changes.|
|`EventInvalid`|Bridged because descendant invalid events do not bubble naturally.|
|`EventFocus`|Observed via bubbling focus signal.|
|`EventNoFocus`|Observed when descendant focus leaves the form.|

## References

* [FormGroup](FormGroup.md)
