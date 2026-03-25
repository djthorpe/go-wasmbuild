# Carbon Package

The `pkg/carbon` package provides Carbon Design System view wrappers on top of the local `pkg/mvc` layer.
Most files in this directory follow the same construction pattern so new components stay predictable and existing ones are easy to extend.

## View Structure

Carbon views are built in four layers:

1. A registered view name in [view.go](view.go)
2. A concrete Go type which embeds the shared `base`
3. An `init()` function which registers the view with `mvc.RegisterView`
4. One or more exported constructors which call `mvc.NewView`

The shared `base` type stores the inner `mvc.View` implementation:

```go
type base struct {
    mvc.View
}
```

That embedded field is wired by `setView`, which means each Carbon component only needs to embed `base` and can then rely on the common MVC behavior.

## Common Pattern

Most components look like this:

```go
type button struct{ base }

func init() {
    mvc.RegisterView(ViewButton, func(element dom.Element) mvc.View {
        return mvc.NewViewWithElement(new(button), element, setView)
    })
}

func Button(args ...any) *button {
    return mvc.NewView(new(button), ViewButton, "cds-button", setView, args).(*button)
}
```

This pattern does two different jobs:

- `mvc.RegisterView(... NewViewWithElement ...)` reconstructs a Carbon view from an existing DOM element
- `mvc.NewView(...)` creates a fresh DOM element or template-backed subtree and returns the concrete Carbon type

Constructors accept a mixed argument list because `mvc.NewView` splits arguments into:

- `mvc.Opt` values such as `mvc.WithClass`, `mvc.WithAttr`, `mvc.WithStyle`, or `mvc.WithID`
- Carbon attribute helpers returned by `carbon.With(...)`
- child views
- DOM elements
- strings and other content inserted into the default content slot

For example:

```go
carbon.Button(
    carbon.With(carbon.KindPrimary, carbon.SizeLarge),
    mvc.WithID("save"),
    "Save",
)
```

`carbon.With(...)` returns `[]mvc.Opt`, so it can be passed directly into any constructor.

## Simple Views

Simple views render from a single tag or custom element name. Examples include:

- [button.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/button.go)
- [button_group.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/button_group.go)
- [checkbox.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/checkbox.go)
- [checkbox_group.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/checkbox_group.go)
- [code.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/code.go)
- [code_snippet.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/code_snippet.go)
- [code_block.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/code_block.go)
- [head.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/head.go)
- [para.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/para.go)
- [lead.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/lead.go)
- [grid.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/grid.go)
- [dropdown.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/dropdown.go)
- [dropdown_item.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/dropdown_item.go)
- [tag.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/tag.go)
- [tag_group.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/tag_group.go)

Typical templates for simple views are:

- a native tag such as `"DIV"`, `"P"`, or `"SECTION"`
- a Carbon web component tag such as `"cds-button"`

Examples:

```go
func Section(args ...any) *container {
    return mvc.NewView(new(container), ViewSection, "SECTION", setView,
        mvc.WithClass("cds--content"), args).(*container)
}

func Head(level int, args ...any) *text {
    tag := fmt.Sprintf("H%d", level)
    cls := fmt.Sprintf("cds--heading-%02d", 7-level)
    return mvc.NewView(new(text), ViewText, tag, setView,
        mvc.WithClass(cls), args).(*text)
}
```

These constructors mostly differ only in:

- the registered view name
- the root tag or Carbon element
- default classes or attributes
- any extra validation or argument normalization

Supporting helper methods should stay close to that construction logic. Use them to normalize constructor arguments or to adapt Carbon-specific DOM details behind the public view API. For example, [button.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/button.go) normalizes icon arguments so icon children are moved into the `slot="icon"` position expected by Carbon.

## Composite Views And Slots

Some Carbon views use an HTML template instead of a single tag. These templates declare named slots with `data-slot`, and the underlying MVC layer keeps track of those slots.

Examples:

- [table.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table.go)
- [table_header.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table_header.go)
- [table_row.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table_row.go)
- [table_toolbar.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table_toolbar.go)
- [table_toolbar_search.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table_toolbar_search.go)
- [blockquote.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/blockquote.go)
- [structured_list.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/structured_list.go)
- [structured_list_row.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/structured_list_row.go)

The table component is representative:

```go
const templateTable = `
    <cds-table size="sm">
        <cds-table-head data-slot="header"></cds-table-head>
        <cds-table-body data-slot="body"></cds-table-body>
    </cds-table>
`

func Table(args ...any) *table {
    return mvc.NewView(new(table), ViewTable, templateTable, setView, args).(*table)
}
```

This lets methods target specific regions of the component:

- `Content(...)` writes to the default `body` slot
- `ReplaceSlot(...)` replaces a named slot
- `ReplaceSlotChildren(...)` replaces the children inside a named slot

Use a template-backed view when a component needs a stable internal structure rather than a single root element.

## Styling And Attributes

Carbon-specific appearance values are defined in [attr.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/attr.go) as typed `Attr` constants.

Examples include:

- kinds such as `KindPrimary` and `KindGhost`
- sizes such as `SizeSmall` and `SizeLarge`
- themes such as `ThemeWhite` and `ThemeG100`
- component-specific flags such as `LinkInline` or `CodeWrapText`

`carbon.With(...)` converts those typed values into `mvc.Opt` values. Most attrs become HTML attributes such as `kind`, `size`, or `type`. Theme attrs are handled as CSS classes like `cds--g10`.

## Component State

View state should be exposed through the interface contracts in [state.go](/Users/djt/projects/go-wasmbuild/pkg/mvc/state.go), not through ad hoc component-specific mutation methods.

Use the existing MVC state interfaces when a Carbon component has readable or writable state, for example:

- `mvc.LabelState`
- `mvc.ValueState`
- `mvc.ActiveState`
- `mvc.EnabledState`
- `mvc.VisibleState`
- `mvc.PaginationState`
- `mvc.ActiveGroup`, `mvc.EnabledGroup`, and `mvc.VisibleGroup` for container-managed state

That means state should be accessed and written through the established getter and setter pairs such as:

- `Label()` and `SetLabel(...)`
- `Value()` and `SetValue(...)`
- `Active()` and `SetActive(...)`
- `Enabled()` and `SetEnabled(...)`
- `Visible()` and `SetVisible(...)`

Do not add custom methods whose only purpose is to alter state when an existing `mvc/state.go` contract already covers it. For example, prefer implementing `mvc.EnabledState` over introducing a separate method such as `Disable()` or `SetDisabled(...)`, and prefer `mvc.VisibleState` over custom show or hide APIs.

If a component needs state, implement the matching MVC interface, add the corresponding interface assertion in the view file, and keep the DOM mapping inside the standard `Set...` method. This keeps Carbon components interchangeable, predictable, and easy for controllers to work with generically.

## Events

Carbon event names are defined in [event.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/event.go).

Those constants are the public event surface for Carbon views and should be used when:

- registering supported events in `mvc.RegisterView(...)`
- attaching listeners with `AddEventListener(...)`
- bridging or normalizing Carbon custom events into package-level event names

A Carbon view component may emit events, either from the component itself or by exposing events from child components that are part of the view composition. Event targets should remain at the view boundary:

- the event target may be the component's own root element
- the event target may be a child component when that child is part of the containing view
- the event target should not be an implementation detail inside a shadow root

In practice, that means Carbon wrappers should normalize or bridge events so controllers observe component-level targets rather than private internal nodes created by Carbon web components. If an event originates from internal Carbon markup, expose it from the wrapping view or from a child view that is already part of the public view structure. Event registration in `mvc.RegisterView(...)` should describe that public event surface, not shadow-DOM implementation details.

This keeps event handling stable even when Carbon internals change, and it ensures application code works with views and components instead of shadow-DOM details.

## Recommended Pattern For New Components

When adding a new Carbon view, follow this order:

1. Add a new `View...` constant in [view.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/view.go)
2. Create a concrete type embedding `base`
3. Register it in `init()` with `mvc.RegisterView(... mvc.NewViewWithElement(...))`
4. Add an exported constructor that calls `mvc.NewView`
5. Use a single tag for simple components, or a template with `data-slot` markers for structured components
6. Add typed behavior methods only when the component needs more than plain content and attributes
7. Add or update the per-component markdown doc next to the Go file

## File Layout

The package is organized roughly like this:

- [view.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/view.go): registered view names and shared `base` wiring
- [attr.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/attr.go): typed Carbon attrs and `With(...)`
- component `*.go` files: concrete constructors and behavior
- component `*.md` files: usage-focused docs for each exported view

If you need an example to copy, start with a simple file such as [grid.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/grid.go) or a structured one such as [table.go](/Users/djt/projects/go-wasmbuild/pkg/carbon/table.go), depending on whether the new component needs slots.
