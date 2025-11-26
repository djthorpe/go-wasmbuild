# Text

Text can be output in various formats including paragraphs, headings, blockquotes and blockquotes with citations.
Its possible to modify text color and alignment, and use the in-build Markdown renderer to output text which uses the markdown syntax.

* Inline text formatting options include `bs.Deleted()`, `bs.Highlighted()`, `bs.Strong()`, `bs.Smaller()`, `bs.Em()`,  and `bs.Code()`
* Links can be created with `bs.Link()` and `bs.IconLink()`
* Blocks of code can be created with `bs.CodeBlock()`

Options for text formatting include:

* `WithColor(bs.Color)` - set the text color
* `WithPosition(bs.Position)` - set the text alignment (`bs.Start`, `bs.Center`, `bs.End`)

The [Bootstrap documentation on Typography](https://getbootstrap.com/docs/5.3/content/typography/) provides more details and examples of text formatting options using class names.
