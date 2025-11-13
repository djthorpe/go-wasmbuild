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

A view is any component that implements the `mvc.View` interface. Views can be composed together to create complex user interfaces. A view can consist of a root DOM element, which can be accessed via the `Root()` method. A few pieces are needed to create a new view component:

```go
// The specific implementation of a view can vary, but it must
// implement the mvc.View interface and the mvc.ViewWithSelf interfaces
type container struct {
    mvc.View
}

// Register the view type with the MVC package
func init() {
    mvc.RegisterView("container-view", CreateContainerViewFromElement)
}

// Satisfy the mvc.ViewWithSelf interface
func (c *container) SetView(v mvc.View) {
    c.View = v
}

// Create a new container view programmatically
func CreateContainerView(args ...any) mvc.View {
    return mvc.NewView(new(container), "container-view", "DIV", args)
}

// Create a container view from an existing DOM element,
// which is used by mvc.ViewFromEvent
func CreateContainerViewFromElement(element dom.Element) mvc.View {
    return new(container)
}
```

In this example, a new view is created with a parent view, a name, a root element of type `DIV`, and optional arguments which set classes, attributes and content for the element. You can then use this view in your application like any other view:

```go
// Resulting HTML:
// <div data-mvc="container-view">hello, world!</div>
doc.Body().AppendChild(CreateContainerView("hello, world!").Root())
```

Views can also be structured with specific sub-elements:

- Optional `Body` content, where child views or elements can be added using methods
  like `Content()` or `Append()`. If the body element is not defined, child views are added
  directly to the root element.
- `Header` and `Footer` elements, which can set their content using the `Header()` and `Footer()` methods.
  If either the header or footer is defined, the body also needs to be defined.
- A `Label` element, which can set their content using the `Label()` method. Labels can
  optionally be inserted at the top or bottom of the root element children.

In order to create a view with these sub-elements, the view implementation must define them and provide methods to access them. For example:

```go
// Create a new container view with a header, footer, body and label
// programmatically
func CreateContainerView(args ...any) mvc.View {
    // Define the sub-elements
    headerElement := mvc.HTML("div", mvc.WithClass("header"))
    bodyElement := mvc.HTML("div", mvc.WithClass("body"))
    footerElement := mvc.HTML("div", mvc.WithClass("footer"))
    labelElement := mvc.HTML("div", mvc.WithClass("label"))

    // Return the view with the sub-elements, with the label at the bottom
    return mvc.NewViewEx(
        new(container), "container-view", "DIV", 
        headerElement, bodyElement, footerElement, labelElement,
        args
    )
}
```

Then the view can set content for the sub-elements:

```go
// Resulting HTML:
// <div data-mvc="container-view">
//  <div class="header">This is the header</div>
//  <div class="body">This is the body</div>
//  <div class="footer">This is the footer</div>
//  <div class="label">This is the label</div>
// </div>
doc.Body().AppendChild(CreateContainerView("This is the body").Header(
    "This is the header"
).Footer(
    "This is the footer"
).Label(
    "This is the label"
).Root())
```

These methods can be chained together to create complex views that include multiple sub-elements,
views, text nodes, and view options.

## View Options

View options are helper functions that can be passed as arguments when creating a new view. They allow you to set classes and attributes in a declarative way. Some common view options include:

- `mvc.WithClass("class-name")`: Adds a class to the view's root element.
- `mvc.WithAttr("attr-name", "attr-value")`: Sets an attribute on the view's root element.
- `mvc.WithoutClass("class-name")`: Removes a class from the view's root element.
- `mvc.WithoutAttr("attr-name")`: Removes an attribute from the view's root element.

You can apply these options by calling `mvc.NewView`, `mvc.NewViewEx`,`mvc.Content`, `mvc.Header`, `mvc.Footer` and `mvc.Label` with the desired options as arguments. You can also define your own custom view options by creating functions that match the `mvc.Opt` type signature.

## Event Listeners

You can attach event listeners to views using `AddEventListener` method. This method takes an event type and a callback function as arguments. The callback function receives an `Event` object, which contains information about the event, and
you can determine the target view:

```go
doc.Body().AppendChild(CreateContainerView("This is the body").AddEventListener("click", func(e mvc.Event) {
    view := mvc.ViewFromEvent(e)
    fmt.Println("Clicked view:", view)
}).Root())
```

## Models

The `mvc.Model` interface represents the data layer of the MVC architecture. You can create a model with a prototype object in order to define the data structure. A model can contain one or more records, each represented by an instance of the prototype object. You can define the fields of the prototype object using struct tags to specify the column names and types.

```go
// Define a person struct with the Name field as the unique key
type Person struct {
    Name string `mvc:"name,key"`
    Age  int  `mvc:"age"`
}

model := mvc.NewModel(new(Person))

// Returns []string{ "name", "age" }
columns := model.Columns()

// Insert new records
model.Insert(&Person{Name: "Alice", Age: 30})
model.Insert(&Person{Name: "Bob", Age: 25})

// Count the number of records
count := model.Count() // returns 2

// Delete records by index, starting at zero.
model.DeleteByIndex(1)

// Delete records by key
model.DeleteByKey("Bob")

// Returns []dom.Node{ "Alice", "30" }
nodes := model.ByKey("Alice")

// Delete all records
model.DeleteAll()
```

The `key` property in the struct tag indicates which field should be used as the unique identifier for each record. You can insert, delete, and query records in the model using the provided methods.

The model defines the following events, and you can use `AddEventListener` to the model to listen for these events:

- insert
- delete
- update

## Controller

The `mvc.ModelController` interface represents the controller layer of the MVC architecture. Controllers are responsible for handling user input, updating the model, and refreshing the view. You can create a controller by implementing the `mvc.Controller` interface and defining methods to handle specific actions.

```go
// Create a new model controller and link to a table view
// This keeps the view "in sync" with the model data (but not vice versa)
contoller := mvc.NewModelController(model, mvc.TableView())
```

## Provider

A provider is an external data source, which can be used to update the model. You generally attach the provider to a controller, which updates the model when the data from the provider changes.

TODO

## Router

`mvc.Router()` registers a hash-based router view that swaps child pages based on `window.location.hash`. Register each page with `Router.Page("#route", view)` and place the router within your application tree. The router listens for `hashchange` events and updates the active page automatically.
