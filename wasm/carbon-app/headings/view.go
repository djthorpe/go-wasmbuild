package headings

import (
	"github.com/djthorpe/go-wasmbuild/pkg/carbon"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func View() mvc.View {
	return carbon.Head(1, "Headings")
}
