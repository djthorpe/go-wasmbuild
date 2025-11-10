package dom_test

import (
	"fmt"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Attr_001(t *testing.T) {
	doc := dom.GetWindow().Document()
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
		if attr == nil {
			t.Fatal("Expected attribute, got nil")
		}
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
		if str := fmt.Sprint(attr); str != test.expected {
			t.Errorf("Expected <%v>, got <%v>", test.expected, str)
		}
	}
}
