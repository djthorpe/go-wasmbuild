# Nav and NavBar

The **NavBar** provides a simple way to create and manage navigation bars in your application. It implements <https://getbootstrap.com/docs/5.3/components/navbar/>. The **NavItem** and **NavDropdown** components can be used to create items within the NavBar.

## Creating a NavBar

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NewNavBar(items ...any) mvc.ViewWithCaption {
    opts := []mvc.Opt{
        // Set options here, e.g. bs.WithColor("dark"), bs.WithSize(bs.SizeLg), etc.
    }
    // Programmatically create a NavBar, and append NavItem or NavDropdown items
    return bs.NavBar(opts...).Append(items...)
}

func NewNavItem(content ...any) mvc.ViewWithState {
    // Programmatically create a NavItem, and append text, icons, or other views
    return bs.NavItem(opts...).Append(content...)
}

func AddBrand(navbar mvc.ViewWithCaption, content ...any) mvc.ViewWithCaption {
    // Add a brand to the NavBar
    return navbar.Caption(content...)
}


```

## Supported Options

**NavBar** supports the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the navbar color (default is "primary") |
| `bs.WithSize(Size)` | Allow NavBar to collapse at a specific breakpoint, displaying a toggle button instead |

**NavItem** supports the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the item color (default is "primary") |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set an item border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |
| `bs.WithDisabled(bool)` | Set NavItem state to disabled |

## Interactivity

You can add an event listener to a NavBar to respond to clicked NavItems

```go
func AddNavItemListener(navbar mvc.ViewWithCaption, callback func(mvc.View)) {
  navbar.AddEventListener("click", func(node Node) {
    view := mvc.ViewFromNode(node)
    if view != nil && view.Name() == bs.ViewNavItem {
      return callback(view)
    }
  })
}
```
