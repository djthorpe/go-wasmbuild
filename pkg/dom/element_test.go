package dom_test

import (
	"bytes"
	"testing"

	// Package imports
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
	assert "github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestElementCreate(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}

	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	assert.False(element.IsConnected())
	assert.Equal("DIV", element.TagName())
	assert.Equal("", element.ID())
	assert.Equal("", element.ClassName())
	assert.Zero(len(element.Attributes()))
	assert.False(element.HasAttributes())

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal("<div></div>", b.String())
}

func TestElementAppend_001(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	child := doc.CreateTextNode("hello, world")
	element.AppendChild(child)
	assert.True(element.HasChildNodes())
	assert.True(element.FirstChild().Equals(child))

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal("<div>hello, world</div>", b.String())
}

func TestElementAppend_002(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	child1 := element.AppendChild(doc.CreateTextNode("hello,"))
	assert.True(element.HasChildNodes())
	assert.True(element.FirstChild().Equals(child1))

	child2 := element.AppendChild(doc.CreateTextNode(" world"))
	assert.True(element.HasChildNodes())
	assert.True(element.LastChild().Equals(child2))

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal("<div>hello, world</div>", b.String())
}

func TestElementRemove_001(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	child1 := element.AppendChild(doc.CreateTextNode("hello,"))
	assert.True(element.HasChildNodes())
	assert.True(element.FirstChild().Equals(child1))

	child2 := element.AppendChild(doc.CreateTextNode(" world"))
	assert.True(element.HasChildNodes())
	assert.True(element.LastChild().Equals(child2))

	element.RemoveChild(child1)
	element.RemoveChild(child2)

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal("<div></div>", b.String())
}

func TestElementID(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	element.SetID("test")
	assert.Equal("test", element.ID())
	assert.Equal("test", element.GetAttribute("id"))

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal(`<div id="test"></div>`, b.String())
}

func TestElementClass(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	element.SetClassName("test")
	assert.Equal("test", element.ClassName())

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal(`<div class="test"></div>`, b.String())
}

func TestElementSetAttribute(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	element.SetAttribute("name", "value")
	assert.Equal("value", element.GetAttribute("name"))
	assert.ElementsMatch([]string{"name"}, element.GetAttributeNames())

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal(`<div name="value"></div>`, b.String())
}

func TestElementRemoveAttribute(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	element.SetAttribute("name", "value")
	assert.Equal("value", element.GetAttribute("name"))
	element.RemoveAttribute("name")

	var b bytes.Buffer
	n, err := element.Write(&b)
	assert.NoError(err)
	assert.Greater(n, 0)
	assert.Equal(`<div></div>`, b.String())
}

func TestElementInnerHTML(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	assert.Equal("", element.InnerHTML())

	element.SetInnerHTML("<!-- comment -->")
	assert.Equal("<!-- comment -->", element.InnerHTML())

	element.SetInnerHTML("")
	assert.Equal("", element.InnerHTML())

	element.SetInnerHTML("comment")
	assert.Equal("comment", element.InnerHTML())

	element.SetInnerHTML("")
	assert.Equal("", element.InnerHTML())

	element.SetInnerHTML("<!-- first --> comment <!-- second -->")
	assert.Equal("<!-- first --> comment <!-- second -->", element.InnerHTML())

	element.SetInnerHTML("")
	assert.Equal("", element.InnerHTML())

	element.SetInnerHTML("<span><b>bold</b></span>")
	assert.Equal("<span><b>bold</b></span>", element.InnerHTML())

	element.SetInnerHTML("<span>comment!</span>")
	assert.Equal("<span>comment!</span>", element.InnerHTML())

	element.SetInnerHTML(`<span class="comment">comment!</span>`)
	assert.Equal(`<span class="comment">comment!</span>`, element.InnerHTML())
}

func TestElementChildren(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	element.SetInnerHTML("<a></a><b></b><c></c><!-- some text --> some other text")
	children := element.Children()
	assert.Len(children, 3)
	assert.Equal("A", children[0].TagName())
	assert.Equal("B", children[1].TagName())
	assert.Equal("C", children[2].TagName())
	assert.Equal(3, element.ChildElementCount())
}
