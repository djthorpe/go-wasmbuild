# Badge

A **Badge** is a small count and labeling component. Badges are often used to add a visual cue to an element, such as a button or a link. It implements <https://getbootstrap.com/docs/5.3/components/badge/>.

## Creating a Badge

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func (opts ...mvc.Opt) Badge() mvc.View {
    // Programmatically create a Badge
    return bs.Badge(opts...)
}


func (opts ...mvc.Opt) PillBadge() mvc.View {
    // Programmatically create a PillBadge, which has rounded edges
    return bs.PillBadge(opts...)
}
```

## Supported Options

In addition to general options, **Badge** supports the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the badge color (default is "primary") |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set a badge border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |

## Examples

```go
return bs.Badge(
    bs.WithColor("danger"),
    bs.WithMargin(bs.PositionAll, bs.Spacing1),
).Append("Error")
```
