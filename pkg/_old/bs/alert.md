# Alert

An **Alert** provides a simple way to create and manage alert messages in your application. Alerts can be used to notify users of important information, warnings, or errors. It implements <https://getbootstrap.com/docs/5.3/components/alerts/>.

## Creating an Alert

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func (opts ...mvc.Opt) View() mvc.ViewWithVisibility {
    // Programmatically create an Alert, which is by default visible
    return bs.Alert(opts...)
}
```

## Supported Options

In addition to general options, **Alert** supports the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the alert color (default is "primary") |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set an alert border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |
| `bs.WithoutVisibility()` | Create an alert which is not initially visible |

## Interactivity

Alerts can be shown or hidden using the general visibility options:

```go
func ToggleVisibility(view mvc.ViewWithVisibility) {
    if view.Visible() {
        view.Hide()
    } else {
        view.Show()
    }
}
```

## Examples

```go
return bs.Alert(
    bs.WithColor("danger"),
    bs.WithBorder(bs.PositionTop|bs.PositionBottom, bs.Dark),
    bs.WithPadding(bs.PositionAll, bs.Spacing3),
    bs.WithMargin(bs.PositionBottom, bs.Spacing4),
)
```
