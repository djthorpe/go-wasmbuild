# Button

**Button** and **OutlineButton** are clickable views that trigger an action or event. Buttons are often used to submit forms, open dialogs, or perform other interactive tasks. They implement <https://getbootstrap.com/docs/5.3/components/buttons/>.

A **ButtonGroup** is a view for grouping multiple buttons together. Button groups are useful for creating toolbars or button sets that share a common purpose. It implements <https://getbootstrap.com/docs/5.3/components/button-group/>.

## Creating a Button or ButtonGroup

```go
import (
    "github.com/djthorpe/go-wasmbuild/pkg/bs"
    "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func (opts ...mvc.Opt) Button(text string) mvc.ViewWithState {
    // Programmatically create a Button
    return bs.Button(opts...).Insert(text)
}

func (opts ...mvc.Opt) OutlineButton(text string) mvc.ViewWithState {
    // Programmatically create an OutlineButton, which has a border but
    // no background color
    return bs.OutlineButton(opts...).Insert(text)
}

func (opts ...mvc.Opt) CloseButton() mvc.ViewWithState {
    // Programmatically create a CloseButton, which has a close icon
    // but no text
    return bs.CloseButton(opts...)
}


func (opts ...mvc.Opt) ButtonGroup(buttons ...mvc.ViewWithState) mvc.ViewWithGroupState {
    // Programmatically create a ButtonGroup
    return bs.ButtonGroup(opts...).Insert(buttons...)
}
```

## Supported Options

In addition to general options, **Button** and **ButtonGroup** support the following specific options:

| Option | Description |
|--------|-------------|
| `bs.WithColor(Color)` | Set the button color (default is "primary") |
| `bs.WithBorder(Position, Color)`<br>`bs.WithBorder(Position)` | Set a button border |
| `bs.WithPadding(Position, Spacing)` | Set padding |
| `bs.WithMargin(Position, Spacing)` | Set margin |
| `bs.WithSize(Size)` | Set size, `bs.Default`, `bs.Small` and `bs.Large` are supported |
| `bs.WithDisabled(bool)` | Set disabled state |
| `bs.WithActive(bool)` | Set active state |

These mirror options available for Bootstrap buttons and groups (see link above).

## Interactivity

Append an event listener to a button or a button group to handle click events.

```go
func HandleButtonClick(button mvc.ViewWithState) {
    button.AddEventListener("click", func(node dom.Node) {
        button := mvc.ViewFromNode(node)
        if button != nil {
            // Handle button click event
        }        
    })
}

func HandleButtonGroupClick(buttongroup mvc.ViewWithGroupState) {
    buttongroup.AddEventListener("click", func(node dom.Node) {
        button := mvc.ViewFromNode(node)
        if button != nil {
            // Handle button click event
        }
    })
}
```

## Examples

```go
return bs.ButtonGroup(
    bs.WithMargin(bs.PositionY, bs.Spacing4),
    bs.WithSize(bs.Fluid),
).Insert(
    bs.Button(
        bs.WithColor("success"),
        bs.WithPadding(bs.PositionX, bs.Spacing3),
        bs.WithID("save-button"),
    ).Insert("Save"),
    bs.OutlineButton(
        bs.WithColor("danger"),
        bs.WithPadding(bs.PositionX, bs.Spacing3),
        bs.WithDisabled(true),
    ).Insert("Delete"),
).AddEventListener("click", func(node dom.Node) {
    button := mvc.ViewFromNode(node).(mvc.ViewWithState)
    if button != nil && button.ID() == "save-button" {
        // Disabled save button until document is saved
        button.SetDisabled(true)
    }        
})
```
