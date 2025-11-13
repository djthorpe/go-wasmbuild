# Container

The **Container** view is a layout component that centers and constrains the width of its content. Containers are used to create responsive layouts that adapt to different screen sizes. It implements <https://getbootstrap.com/docs/5.3/layout/containers/>.

## Creating a Container

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func (opts ...mvc.Opt) Container() mvc.View {
    // Programmatically create a Container
    return bs.Container(opts...)
}
```

## Supported Options

In addition to general options, **Container**  support the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the text color (default is theme-based) |
| `bs.WithBackground(Color)` | Set the background color (default is theme-based) |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set the container border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |
| `bs.WithSize(Size)` | Set breakpoint, `bs.Fluid` extends the container width to the window width |

These mirror options available in Bootstrap containers (see link above).

## Examples

```go
return bs.Container(
    bs.WithMargin(bs.PositionY, bs.Spacing4),
    bs.WithSize(bs.Fluid),
)
```
