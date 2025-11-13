package dom_test

import (
	"bytes"
	"testing"

	// Packages
	. "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Window_001(t *testing.T) {
	window := GetWindow()
	if window == nil {
		t.Fatal("Window() returned nil")
	} else {
		t.Log("window=", window)
	}

	document := window.Document()
	if document == nil {
		t.Fatal("Document() returned nil")
	}
	w := new(bytes.Buffer)
	if _, err := window.Write(w, document); err != nil {
		t.Fatal(err)
	}
	t.Log("document=", w.String())
}

/*
func Test_Window_002(t *testing.T) {
	window := GetWindow()
	tests := []struct {
		mime, data string
	}{
		{"text/xml", "<warning>Beware of the tiger</warning>"},
		{"image/svg+xml", "<circle cx=\"50\" cy=\"50\" r=\"50\"/>"},
		{"text/html", "<strong>Beware of the leopard</strong>"},
		{"text/html", "<!DOCTYPE html><html><head><title>test</title><meta charset=\"utf-8\"></head><body></body></html>"},
	}
	for _, test := range tests {
		buf := bytes.NewBufferString(test.data)
		document, err := window.Read(buf, test.mime)
		if err != nil {
			t.Error(err)
			continue
		}
		buf.Reset()
		if _, err := window.Write(buf, document); err != nil {
			t.Error(err)
			continue
		}
		t.Log("document=", buf.String())
	}
}
*/
