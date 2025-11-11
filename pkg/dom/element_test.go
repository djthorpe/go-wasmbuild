package dom_test

import (
	"bytes"
	"testing"

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
	assert.True(element.HasAttribute("id"))
	assert.True(element.HasAttributes())

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
	assert.True(element.HasAttribute("class"))

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
	assert.True(element.HasAttribute("name"))
	assert.True(element.HasAttributes())
	assert.ElementsMatch([]string{"name"}, element.GetAttributeNames())

	node := element.GetAttributeNode("name")
	if assert.NotNil(node) {
		owner := node.OwnerElement()
		if assert.NotNil(owner) {
			assert.True(owner.Equals(element))
		}
	}

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

	assert.Equal("", element.GetAttribute("name"))
	assert.Nil(element.GetAttributeNode("name"))
	assert.Empty(element.Attributes())
	assert.False(element.HasAttribute("name"))
	assert.False(element.HasAttributes())

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

func TestElementElementSiblingNavigation(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	list := doc.CreateElement("ul")
	if !assert.NotNil(list) {
		assert.FailNow("Expected element, got nil")
	}

	list.AppendChild(doc.CreateTextNode("preamble"))

	first := doc.CreateElement("li")
	first.SetAttribute("data-id", "first")
	list.AppendChild(first)

	list.AppendChild(doc.CreateComment("skip"))

	middle := doc.CreateElement("li")
	middle.SetAttribute("data-id", "middle")
	list.AppendChild(middle)

	list.AppendChild(doc.CreateTextNode("spacer"))

	last := doc.CreateElement("li")
	last.SetAttribute("data-id", "last")
	list.AppendChild(last)

	assert.True(list.FirstElementChild().Equals(first))
	assert.True(list.LastElementChild().Equals(last))

	assert.Nil(first.PreviousElementSibling())
	assert.True(first.NextElementSibling().Equals(middle))

	assert.True(middle.PreviousElementSibling().Equals(first))
	assert.True(middle.NextElementSibling().Equals(last))

	assert.True(last.PreviousElementSibling().Equals(middle))
	assert.Nil(last.NextElementSibling())
}

func TestElementSelectorQueries(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	container := doc.CreateElement("section")
	if !assert.NotNil(container) {
		assert.FailNow("Expected element, got nil")
	}

	articleOne := doc.CreateElement("article")
	articleOne.SetAttribute("data-id", "article-1")
	articleOne.SetClassName("featured highlight")
	container.AppendChild(articleOne)

	articleInner := doc.CreateElement("div")
	articleInner.SetAttribute("data-id", "feature-inner")
	articleInner.SetClassName("featured")
	articleOne.AppendChild(articleInner)

	articleTwo := doc.CreateElement("article")
	articleTwo.SetAttribute("data-id", "article-2")
	articleTwo.SetClassName("highlight")
	container.AppendChild(articleTwo)

	aside := doc.CreateElement("aside")
	aside.SetAttribute("data-id", "aside-1")
	container.AppendChild(aside)

	byClass := container.GetElementsByClassName("featured")
	ids := make(map[string]struct{}, len(byClass))
	for _, el := range byClass {
		ids[el.GetAttribute("data-id")] = struct{}{}
	}
	assert.Len(ids, 2)
	_, okArticle := ids["article-1"]
	_, okInner := ids["feature-inner"]
	assert.True(okArticle)
	assert.True(okInner)

	byTag := container.GetElementsByTagName("article")
	tagIDs := make(map[string]struct{}, len(byTag))
	for _, el := range byTag {
		tagIDs[el.GetAttribute("data-id")] = struct{}{}
	}
	assert.Len(tagIDs, 2)
	_, okFirst := tagIDs["article-1"]
	_, okSecond := tagIDs["article-2"]
	assert.True(okFirst)
	assert.True(okSecond)
}

func TestElementSetAttributeNodeLifecycle(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}
	element := doc.CreateElement("div")
	if !assert.NotNil(element) {
		assert.FailNow("Expected element, got nil")
	}

	first := doc.CreateAttribute("data-id")
	first.SetValue("first")
	replaced := element.SetAttributeNode(first)
	assert.Nil(replaced)
	assert.True(element.HasAttribute("data-id"))
	assert.True(first.OwnerElement().Equals(element))

	second := doc.CreateAttribute("data-id")
	second.SetValue("second")
	replaced = element.SetAttributeNode(second)
	if assert.NotNil(replaced) {
		assert.Nil(replaced.OwnerElement())
		assert.Equal("first", replaced.Value())
	}
	assert.True(second.OwnerElement().Equals(element))
	assert.Equal("second", element.GetAttribute("data-id"))

	element.RemoveAttributeNode(second)
	assert.False(element.HasAttribute("data-id"))
	assert.False(element.HasAttributes())
	assert.Equal("", element.GetAttribute("data-id"))
	assert.Nil(second.OwnerElement())
}

func TestElementAppendChildReparents(t *testing.T) {
	assert := assert.New(t)
	doc := dom.GetWindow().Document()
	if !assert.NotNil(doc) {
		assert.FailNow("Expected document, got nil")
	}

	parentA := doc.CreateElement("div")
	parentB := doc.CreateElement("div")
	child := doc.CreateElement("span")

	parentA.AppendChild(child)
	assert.True(child.ParentNode().Equals(parentA))
	assert.Len(parentA.Children(), 1)

	parentB.AppendChild(child)
	assert.True(child.ParentNode().Equals(parentB))
	assert.Len(parentA.Children(), 0)
	assert.Len(parentB.Children(), 1)
}
