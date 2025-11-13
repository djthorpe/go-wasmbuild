package dom_test

import (
	"bytes"
	"testing"

	// Modules
	. "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Doctype_001(t *testing.T) {
	doc := GetWindow().Document()
	if doc.Doctype() == nil {
		t.Error("Doctype() returned nil")
	} else {
		t.Log(doc.Doctype())
	}
}

func Test_Doctype_002(t *testing.T) {
	win := GetWindow()
	doc := win.Document()
	w := new(bytes.Buffer)
	if _, err := win.Write(w, doc); err != nil {
		t.Error(err)
	} else {
		t.Log(w.String())
	}
}
