# Code

**Code** and **CodeBlock** views provide a way to display inline code snippets and preformatted code blocks, respectively, using Bootstrap styles.

## Creating a Code and CodeBlock view

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func (opts ...mvc.Opt) Code(content string) mvc.ViewWithState {
    // Programmatically create inline Code
    return bs.Code(content)
}

func (opts ...mvc.Opt) CodeBlock(content string) mvc.ViewWithState {
    // Programmatically create a CodeBlock
    return bs.CodeBlock().Content(content)
}
```

## Supported Options

In addition to general options, **Code** and **CodeBlock** support the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the default text color (default based on theme) |
| `bs.WithBackgroundColor(Color)` | Set the button background color (default based on theme) |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set a button border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |
