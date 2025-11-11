package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Code() mvc.View {
	return bs.Container(bs.WithSize(bs.SizeFluid)).Append(
		bs.Heading(2).Append(
			"Code",
		),
		bs.Para(
			"This is some ",
			bs.Code("inline code"),
			" inside a paragraph.",
		),
		bs.CodeBlock().Append(
			`<html>
	<head>
		<title>My Page</title>
	</head>
	<body>
		<h1>Hello, World!</h1>
	</body>
</html>
`,
		),
	)
}
