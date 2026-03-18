package headings

import (
	"github.com/djthorpe/go-wasmbuild/pkg/carbon"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func View() []any {
	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Headings")),
	}
}
