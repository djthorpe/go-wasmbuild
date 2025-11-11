package bs

import (
	"slices"
	"testing"

	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TestViewTextConstant(t *testing.T) {
	expected := "mvc-bs-text"
	if ViewText != expected {
		t.Errorf("ViewText = %v, want %v", ViewText, expected)
	}
}

func TestTextRegistered(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Text view should be registered, but got panic: %v", r)
		}
	}()
	_ = Para()
}

func TestTextType(t *testing.T) {
	var _ *text = &text{}
	txt := &text{}
	if _, ok := any(txt).(interface{ View }); !ok {
		t.Error("text should embed View")
	}
}

func TestTextTagNames(t *testing.T) {
	expectedTags := []string{"P", "SPAN", "DEL", "MARK", "SMALL", "STRONG", "EM", "BLOCKQUOTE"}
	if len(textTagNames) != len(expectedTags) {
		t.Errorf("textTagNames length = %d, want %d", len(textTagNames), len(expectedTags))
	}
	for _, expected := range expectedTags {
		if !slices.Contains(textTagNames, expected) {
			t.Errorf("textTagNames missing %q", expected)
		}
	}
}

func TestPara(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with ID", []Opt{WithID("intro")}},
		{"create with classes", []Opt{WithClass("lead", "text-muted")}},
		{"create with combined options", []Opt{WithID("paragraph-1"), WithClass("fs-5")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Para()
			if p == nil {
				t.Fatal("Para() returned nil")
			}
			if len(tt.opts) > 0 {
				p.Opts(tt.opts...)
			}
			if p.Name() != ViewText {
				t.Errorf("Para().Name() = %v, want %v", p.Name(), ViewText)
			}
		})
	}
}

func TestDeleted(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with class", []Opt{WithClass("text-decoration-line-through")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Deleted()
			if len(tt.opts) > 0 {
				d.Opts(tt.opts...)
			}
			if d == nil {
				t.Fatal("Deleted() returned nil")
			}
			if d.Name() != ViewText {
				t.Errorf("Deleted().Name() = %v, want %v", d.Name(), ViewText)
			}
		})
	}
}

func TestHighlighted(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with ID", []Opt{WithID("highlight-1")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Highlighted()
			if len(tt.opts) > 0 {
				h.Opts(tt.opts...)
			}
			if h == nil {
				t.Fatal("Highlighted() returned nil")
			}
			if h.Name() != ViewText {
				t.Errorf("Highlighted().Name() = %v, want %v", h.Name(), ViewText)
			}
		})
	}
}

func TestSmall(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with muted class", []Opt{WithClass("text-muted")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Small()
			if len(tt.opts) > 0 {
				s.Opts(tt.opts...)
			}
			if s == nil {
				t.Fatal("Small() returned nil")
			}
			if s.Name() != ViewText {
				t.Errorf("Small().Name() = %v, want %v", s.Name(), ViewText)
			}
		})
	}
}

func TestStrong(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with fw-bold class", []Opt{WithClass("fw-bold")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Strong()
			if len(tt.opts) > 0 {
				s.Opts(tt.opts...)
			}
			if s == nil {
				t.Fatal("Strong() returned nil")
			}
			if s.Name() != ViewText {
				t.Errorf("Strong().Name() = %v, want %v", s.Name(), ViewText)
			}
		})
	}
}

func TestEm(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with fst-italic class", []Opt{WithClass("fst-italic")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Em()
			if len(tt.opts) > 0 {
				e.Opts(tt.opts...)
			}
			if e == nil {
				t.Fatal("Em() returned nil")
			}
			if e.Name() != ViewText {
				t.Errorf("Em().Name() = %v, want %v", e.Name(), ViewText)
			}
		})
	}
}

func TestBlockquote(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{"create without options", nil},
		{"create with ID", []Opt{WithID("quote-1")}},
		{"create with additional classes", []Opt{WithClass("blockquote-footer")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bq := Blockquote()
			if len(tt.opts) > 0 {
				bq.Opts(tt.opts...)
			}
			if bq == nil {
				t.Fatal("Blockquote() returned nil")
			}
			if bq.Name() != ViewText {
				t.Errorf("Blockquote().Name() = %v, want %v", bq.Name(), ViewText)
			}
		})
	}
}

func TestTextViewInterface(t *testing.T) {
	elements := []struct {
		name    string
		element *text
	}{
		{"Para", Para()},
		{"Deleted", Deleted()},
		{"Highlighted", Highlighted()},
		{"Small", Small()},
		{"Strong", Strong()},
		{"Em", Em()},
		{"Blockquote", Blockquote()},
	}
	for _, elem := range elements {
		t.Run(elem.name, func(t *testing.T) {
			if elem.element.Name() != ViewText {
				t.Errorf("%s.Name() = %v, want %v", elem.name, elem.element.Name(), ViewText)
			}
			root := elem.element.Root()
			if root == nil {
				t.Errorf("%s.Root() returned nil", elem.name)
			}
			_ = elem.element.ID()
		})
	}
}

func TestTextChaining(t *testing.T) {
	t.Run("chaining pattern compiles", func(t *testing.T) {
		_ = func(txt *text) View { return txt.Opts(WithClass("test")) }
	})
}

func TestNewTextFromElement(t *testing.T) {
	t.Run("function exists", func(t *testing.T) {
		var _ ViewConstructorFunc = newTextFromElement
	})
}

func TestTextEmbedding(t *testing.T) {
	p := Para()
	var v View = p
	if v.Name() != ViewText {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), ViewText)
	}
}

func TestTextMultipleInstances(t *testing.T) {
	p1 := Para()
	p1.Opts(WithID("p1"))
	p2 := Para()
	p2.Opts(WithID("p2"))
	if p1 == p2 {
		t.Error("Each text element call should create a distinct instance")
	}
	if p1.ID() == p2.ID() {
		t.Error("Text elements with different IDs should maintain separate state")
	}
}

func TestTextWithOptions(t *testing.T) {
	t.Run("with ID", func(t *testing.T) {
		p := Para()
		p.Opts(WithID("my-para"))
		if p.ID() != "my-para" {
			t.Errorf("Para().Opts(WithID('my-para')).ID() = %v, want %v", p.ID(), "my-para")
		}
	})
	t.Run("with classes", func(t *testing.T) {
		p := Para()
		p.Opts(WithClass("lead", "text-center"))
		if p == nil {
			t.Error("Para() with classes should not return nil")
		}
	})
	t.Run("with combined options", func(t *testing.T) {
		s := Span()
		s.Opts(WithID("highlight"), WithClass("badge", "bg-success"), WithAttr("data-type", "status"))
		if s == nil {
			t.Fatal("Span() with combined options should not return nil")
		}
		if s.ID() != "highlight" {
			t.Errorf("Span().ID() = %v, want %v", s.ID(), "highlight")
		}
	})
}

func TestBootstrapTypographyClasses(t *testing.T) {
	tests := []struct {
		name    string
		factory func() *text
		classes []string
	}{
		{"lead paragraph", func() *text { return Para() }, []string{"lead"}},
		{"display heading in paragraph", func() *text { return Para() }, []string{"display-1"}},
		{"font weight bold", func() *text { return Strong() }, []string{"fw-bold"}},
		{"font style italic", func() *text { return Em() }, []string{"fst-italic"}},
		{"text decoration line-through", func() *text { return Deleted() }, []string{"text-decoration-line-through"}},
		{"line height", func() *text { return Para() }, []string{"lh-1", "lh-sm", "lh-base", "lh-lg"}},
		{"text alignment", func() *text { return Para() }, []string{"text-start", "text-center", "text-end"}},
		{"text wrapping", func() *text { return Para() }, []string{"text-wrap", "text-nowrap"}},
		{"text break", func() *text { return Para() }, []string{"text-break"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elem := tt.factory()
			elem.Opts(WithClass(tt.classes...))
			if elem == nil {
				t.Errorf("%s factory returned nil with classes %v", tt.name, tt.classes)
			}
		})
	}
}

func TestBlockquoteBootstrapClasses(t *testing.T) {
	t.Run("blockquote with footer", func(t *testing.T) {
		bq := Blockquote()
		bq.Opts(WithClass("blockquote-footer"))
		if bq == nil {
			t.Error("Blockquote with footer class should not return nil")
		}
	})
	t.Run("blockquote with text alignment", func(t *testing.T) {
		bq := Blockquote()
		bq.Opts(WithClass("text-center"))
		if bq == nil {
			t.Error("Blockquote with text-center should not return nil")
		}
	})
	t.Run("blockquote with text end", func(t *testing.T) {
		bq := Blockquote()
		bq.Opts(WithClass("text-end"))
		if bq == nil {
			t.Error("Blockquote with text-end should not return nil")
		}
	})
}

func TestTextColorUtilities(t *testing.T) {
	colors := []string{
		"text-primary", "text-secondary", "text-success", "text-danger",
		"text-warning", "text-info", "text-light", "text-dark",
		"text-body", "text-muted", "text-white", "text-black-50", "text-white-50",
	}
	for _, color := range colors {
		t.Run(color, func(t *testing.T) {
			p := Para()
			p.Opts(WithClass(color))
			if p == nil {
				t.Errorf("Para with %s class should not return nil", color)
			}
		})
	}
}

func TestTextOpacityUtilities(t *testing.T) {
	opacities := []string{"text-opacity-25", "text-opacity-50", "text-opacity-75", "text-opacity-100"}
	for _, opacity := range opacities {
		t.Run(opacity, func(t *testing.T) {
			s := Span()
			s.Opts(WithClass(opacity))
			if s == nil {
				t.Errorf("Span with %s class should not return nil", opacity)
			}
		})
	}
}

func TestAllTextElementsAreDistinct(t *testing.T) {
	elements := []*text{Para(), Deleted(), Highlighted(), Small(), Strong(), Em(), Blockquote()}
	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements); j++ {
			if elements[i] == elements[j] {
				t.Errorf("Text elements at indices %d and %d are the same instance", i, j)
			}
		}
	}
}
