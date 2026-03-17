# MVC Package

The `github.com/djthorpe/go-wasmbuild/pkg/mvc` package provides a thin "model, view and controller" style abstraction on top of the DOM wrappers that ship with go-wasmbuild. It is designed for Go/WebAssembly applications that implement applications a declarative, component-oriented approach to building
user interfaces.

## Getting Started

Import the package alongside the base DOM facade:

```go
import (
    mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)
```

Create an application root and populate it with content:

```go
func main() {
    app := mvc.New()
    app.Content("Hello, world!")

    select {} // keep the WASM module alive
}
```

Views expose fluent helpers such as `Content`, `Append`, and `AddEventListener` so you can assemble larger components quickly. Many sub-packages (for example, `pkg/bootstrap`) build on these primitives to provide higher-level UI widgets.

## View

A `View` is any component that implements the `mvc.View` interface. Views wrap a root DOM element and can be composed to build complex UIs.

```go
type View interface {
    Name() string                                       // registered component name
    ID() string                                         // value of the id attribute, if set
    Root() dom.Element                                  // the root DOM element
    Parent() View                                       // nearest ancestor view, or nil
    Self() View                                         // concrete implementation (for embedding)
    Slot(name string) dom.Element                       // named slot element, or nil
    ReplaceSlot(name string, node any, opts ...Opt) View
    ReplaceSlotChildren(name string, args ...any) View
    Apply(opts ...Opt) View                             // apply class/attribute options
    Content(args ...any) View                           // replace body slot children
    AddEventListener(event string, fn func(dom.Event)) View
    RemoveEventListener(event string) View
}
```

### Creating a view

Views must be registered before use. A minimal custom view embeds `mvc.View` and provides an initialisation function that stores the inner view:

```go
type myWidget struct {
    mvc.View
}

func init() {
    mvc.RegisterView("my-widget", func(el dom.Element) mvc.View {
        return mvc.NewViewWithElement(new(myWidget), el, func(self, child mvc.View) {
            self.(*myWidget).View = child
        })
    })
}

func MyWidget(args ...any) mvc.View {
    return mvc.NewView(new(myWidget), "my-widget", "DIV", func(self, child mvc.View) {
        self.(*myWidget).View = child
    }, args...)
}
```

`NewView` creates the root element from a tag name (e.g. `"DIV"`), applies any `Opt` arguments, and inserts any content arguments into the body slot. The `data-mvc` attribute is set automatically to the registered name.

### Slots

Slots are named sub-elements within a view's HTML template. The default slot is named `"body"` (the `ContentSlot` constant). `Content(...)` is shorthand for `ReplaceSlotChildren("body", ...)`.

Custom templates can define additional slots by giving child elements a `data-slot` attribute:

```html
<div data-mvc="my-widget">
    <header data-slot="header"></header>
    <div data-slot="body"></div>
    <footer data-slot="footer"></footer>
</div>
```

You can then target each slot by name:

```go
w := MyWidget()
w.ReplaceSlotChildren("header", "Page title")
w.Content("Main body text")           // same as ReplaceSlotChildren("body", ...)
w.ReplaceSlotChildren("footer", "Footer text")
```

### Creating elements

`mvc.HTML` creates a standalone DOM element (not a registered view) that can be passed as content to any view:

```go
mvc.HTML("SPAN", "hello", mvc.WithClass("text-muted"))
mvc.HTML("BR")
mvc.HTML("A", "click me", mvc.WithAttr("href", "#home"))
```

### Application root

`mvc.New` creates the root application container, prepends it to `document.body`, and returns it. Call `Run()` to block the main goroutine and keep the WASM module alive:

```go
func main() {
    mvc.New(
        bs.Container(
            mvc.WithClass("p-3"),
            "Hello, world!",
        ),
    ).Run()
}
```

## View Options

View options are functions with signature `func(OptSet) error` (the `mvc.Opt` type) that declaratively set classes and attributes on a view's root element. They can be passed to `NewView`, `Content`, `ReplaceSlotChildren`, `ReplaceSlot`, and `Apply`.

Common built-in options:

- `mvc.WithClass("name")` — add a CSS class
- `mvc.WithoutClass("name")` — remove a CSS class
- `mvc.WithAttr("key", "value")` — set an attribute
- `mvc.WithoutAttr("key")` — remove an attribute
- `mvc.WithID("id")` — set the element id
- `mvc.WithStyle("css")` — set the inline style attribute
- `mvc.WithSlotAttr("slot", "key", "value")` — set an attribute on a named slot element

Custom options can be created by implementing the `mvc.Opt` function signature:

## Event Listeners

Attach event listeners to a view's root element with `AddEventListener`. The method returns the view so it can be chained:

```go
btn := bs.Button("Click me").AddEventListener("click", func(e dom.Event) {
    fmt.Println("clicked")
})
```

`RemoveEventListener` removes a previously registered listener by event type.

## Models

There are two generic model types. Both are zero-value usable structs.

### Model[T]

`Model[T]` stores an ordered slice of items of any type and notifies listeners on every mutation.

```go
var m mvc.Model[string]

m.AddEventListener(func(items []string) {
    fmt.Println("changed:", items) // do not modify the slice
})

m.Set([]string{"a", "b", "c"}) // fires listener
m.Append("d")                   // fires listener
m.Clear()                       // fires listener

fmt.Println(m.Len())    // 0
fmt.Println(m.Items())  // [] (copy)
```

| Method | Description |
|---|---|
| `Set([]T)` | Replace all items; notify listeners |
| `Append(...T)` | Add items to the end; notify listeners |
| `Clear()` | Remove all items; notify listeners |
| `Items() []T` | Return a shallow copy |
| `Len() int` | Number of items |
| `AddEventListener(func([]T))` | Register change listener |

### KeyedModel[K, T]

`KeyedModel[K comparable, T Keyed[K]]` is like `Model[T]` but requires items to implement `PrimaryKey() K`. It maintains an internal index for O(1) key lookups and emits fine-grained added/deleted/changed events in addition to the full-list `OnSet` event.

```go
type Station struct{ Abbr, Name string }
func (s Station) PrimaryKey() string { return s.Abbr }

var m mvc.KeyedModel[string, Station]

m.OnSet(func(items []Station) { /* any mutation */ })
m.OnAdded(func(e mvc.AddedEvent[Station]) {
    fmt.Println("added at index", e.Index, ":", e.Item.Name)
})
m.OnDeleted(func(e mvc.DeletedEvent[Station]) {
    fmt.Println("removed from index", e.Index, ":", e.Item.Name)
})
m.OnChanged(func(e mvc.ChangedEvent[Station]) {
    fmt.Println("updated at index", e.Index, ":", e.Item.Name)
})

m.Set([]Station{{"RICH", "Richmond"}, {"12TH", "12th St."}}) // OnSet only
m.Append(Station{"EMBR", "Embarcadero"})                     // OnAdded + OnSet
m.Update(Station{"RICH", "Richmond (updated)"})              // OnChanged + OnSet
m.Remove("12TH")                                             // OnDeleted + OnSet

station, ok := m.Get("EMBR") // O(1) lookup
```

Each event struct carries `Item T`, `Index int`, and `Items []T` (full post-mutation slice), giving views everything needed for either a targeted DOM update or a full re-render.

To use `KeyedModel`, implement `mvc.Keyed[K]` on your type:

```go
type Keyed[K comparable] interface {
    PrimaryKey() K
}
```

## Controller

`mvc.Controller` is an interface for wiring views to application logic. A controller attaches to one or more views and reacts to events they emit.

```go
type Controller interface {
    Views() []View
    Attach(...View)
    Detach(...View)
    EventListener(eventType string, source View)
}
```

Create a controller by embedding `*mvc.controller` in your own struct and calling `mvc.NewController`:

```go
type myController struct {
    *mvc.controller // not exported; access via Controller interface
}

func NewMyController(views ...mvc.View) mvc.Controller {
    c := &myController{}
    c.controller = mvc.NewController(c, nil, views...)
    return c
}

func (c *myController) EventListener(event string, v mvc.View) {
    fmt.Printf("event %q from %s\n", event, v.Name())
}
```

For domain-specific controllers (e.g. wrapping a data provider and models), define your own struct that owns the provider and models directly and exposes typed methods — see the BART app's `bart.Controller` for a worked example.

## Provider

Providers fetch data from remote sources and fan results out to registered listeners. There are two provider types.

### Provider

`mvc.Provider` performs raw HTTP fetches and delivers `*js.FetchResponse` to listeners:

```go
type Provider interface {
    Fetch(path string, opts ...js.FetchOption)
    FetchWithInterval(path string, interval time.Duration, opts ...js.FetchOption)
    Cancel()
    AddEventListener(fn func(*js.FetchResponse, error))
}

p := mvc.NewProvider(baseURL)
p.AddEventListener(func(resp *js.FetchResponse, err error) {
    if err != nil { return }
    resp.Text().Done(func(v js.Value, err error) {
        fmt.Println(v.String())
    })
})
p.Fetch("data.json")
```

`FetchWithInterval` cancels any existing interval, fires immediately, then repeats. Call `Cancel()` to stop.

### JSONProvider[T]

`mvc.JSONProvider[T]` wraps `Provider` and automatically decodes JSON responses into `T`. It supports all HTTP verbs and marshals request bodies for mutation methods:

```go
type JSONProvider[T any] interface {
    Get(path string, opts ...js.FetchOption)
    GetWithInterval(path string, interval time.Duration, opts ...js.FetchOption)
    Cancel()
    Post(path string, body T, opts ...js.FetchOption)
    Put(path string, body T, opts ...js.FetchOption)
    Patch(path string, body T, opts ...js.FetchOption)
    Delete(path string, opts ...js.FetchOption)
    AddEventListener(fn func(T, error))
}

type MyData struct {
    Name string `json:"name"`
}

p := mvc.NewJSONProvider[MyData](baseURL)
p.AddEventListener(func(data MyData, err error) {
    if err != nil { return }
    fmt.Println(data.Name)
})
p.Get("items/1")
```

A 204 No Content response calls the listener with the zero value of `T` and a nil error. Use `js.WithQuery(url.Values{...})` to append query parameters.

## Router

`mvc.Router()` creates a hash-based router view that swaps the displayed page based on `window.location.hash`. It listens for `hashchange` events and updates the content area automatically. If no hash matches, the first registered page is shown as the default.

### Basic usage

```go
router := mvc.Router().
    Page("#home",     homeView).
    Page("#stations", stationsView)

mvc.New(router).Run()
```

Navigating to `#stations` (via an `<a href="#stations">` link or `window.location.hash = "#stations"`) swaps the content area to `stationsView`.

### Integrating with a nav group

Pass any `mvc.ActiveGroup` to `Active(...)` so the router automatically marks the correct nav items active on every navigation:

```go
navHome     := bs.NavItem("Home",     mvc.WithAttr("href", "#home"))
navStations := bs.NavItem("Stations", mvc.WithAttr("href", "#stations"))
nav         := bs.Nav(navHome, navStations)  // Nav implements mvc.ActiveGroup

router := mvc.Router().
    Active(nav).
    Page("#home",     homeView,     navHome).
    Page("#stations", stationsView, navStations)
```

When the route changes to `#stations`, the router calls `nav.SetActive(navStations)`. The `ActiveGroup` implementation marks `navStations` active and all other members inactive.

### ActiveState and ActiveGroup interfaces

`mvc.ActiveState` is implemented by any view that can be marked active or inactive:

```go
type ActiveState interface {
    SetActive(bool)
}
```

`mvc.ActiveGroup` is implemented by a container (such as a nav bar) that manages which of its members are currently active:

```go
type ActiveGroup interface {
    SetActive(views ...View)  // activates given views; deactivates the rest
}
```

Passing no arguments to `SetActive` deactivates all members.
