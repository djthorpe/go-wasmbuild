package dom_test

import (
	"bytes"
	"testing"

	// Packages
	. "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Attr_001(t *testing.T) {
	win := GetWindow()
	doc := win.Document()
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"name", "", `name=""`},
		{"name", "test", `name="test"`},
		{"name", "&", `name="&amp;"`},
		{"name", "<test>", `name="&lt;test&gt;"`},
		{"name", `"test"`, `name="&#34;test&#34;"`},
		{"name", `'test'`, `name="&#39;test&#39;"`},
	}
	for _, test := range tests {
		attr := doc.CreateAttribute(test.name)
		if attr.Name() != test.name {
			t.Errorf("Expected %q, got %q", test.name, attr.Name())
		}
		if attr.Value() != "" {
			t.Errorf("Expected empty string, got %q", attr.Value())
		}
		if test.value != "" {
			attr.SetValue(test.value)
			if attr.Value() != test.value {
				t.Errorf("Expected %q, got %q", test.value, attr.Value())
			}
		}
		buf := new(bytes.Buffer)
		if _, err := win.Write(buf, attr); err != nil {
			t.Error(err)
		} else if buf.String() != test.expected {
			t.Errorf("Expected <%v>, got <%v>", test.expected, buf.String())
		}
	}
}
