//go:build !(js && wasm)

package dom

import (
	"testing"

	// Packages
	"github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestNode_001(t *testing.T) {
	assert := assert.New(t)

	// Create a node
	node := newNode(nil, nil, "#node", UNKNOWN_NODE, "cdata")
	assert.NotNil(node)
	assert.Nil(node.OwnerDocument())
	assert.Nil(node.ParentNode())
	assert.Equal("#node", node.NodeName())
	assert.Equal(UNKNOWN_NODE, node.NodeType())
	assert.Equal("cdata", node.TextContent())
}

func TestNode_002(t *testing.T) {
	assert := assert.New(t)

	// Create a node
	node := newNode(nil, nil, "#node", UNKNOWN_NODE, "")

	// Add a child
	child := newNode(nil, nil, "#child", UNKNOWN_NODE, "child")
	node.AppendChild(&child)

	// Check child is added to parent
	assert.True(node.HasChildNodes())
	assert.True(node.Contains(&child))
	assert.Len(node.ChildNodes(), 1)

	children := node.ChildNodes()
	assert.True(child.Equals(children[0]))
	assert.True(child.Equals(node.FirstChild()))
	assert.True(child.Equals(node.LastChild()))
	assert.True(child.IsConnected())
	assert.True(node.Equals(child.ParentNode()))
	assert.Nil(child.NextSibling())
	assert.Nil(child.PreviousSibling())
}

func TestNode_003(t *testing.T) {
	assert := assert.New(t)

	// Create two parents
	node1 := newNode(nil, nil, "#node1", UNKNOWN_NODE, "")
	node2 := newNode(nil, nil, "#node2", UNKNOWN_NODE, "")

	// Add a child to node1
	child := newNode(nil, nil, "#child", UNKNOWN_NODE, "child")
	node1.AppendChild(&child)

	// Check child is added to node1
	assert.True(node1.HasChildNodes())
	assert.True(node1.Contains(&child))
	assert.Len(node1.ChildNodes(), 1)

	// Now add child to node2
	node2.AppendChild(&child)
	assert.True(node2.HasChildNodes())
	assert.True(node2.Contains(&child))
	assert.Len(node2.ChildNodes(), 1)

	// Make sure node1 no longer has children
	assert.False(node1.HasChildNodes())
	assert.False(node1.Contains(&child))
	assert.Len(node1.ChildNodes(), 0)
}

func TestNode_004(t *testing.T) {
	assert := assert.New(t)

	// Create a parent
	node := newNode(nil, nil, "#node1", UNKNOWN_NODE, "")

	// Add two children to node
	child1 := newNode(nil, nil, "#child1", UNKNOWN_NODE, "child1")
	node.AppendChild(&child1)

	child2 := newNode(nil, nil, "#child2", UNKNOWN_NODE, "child2")
	node.AppendChild(&child2)

	// Check children are added
	assert.True(node.HasChildNodes())
	assert.True(node.Contains(&child1))
	assert.True(node.Contains(&child2))
	assert.Len(node.ChildNodes(), 2)

	// Remove child1
	node.RemoveChild(&child1)
	assert.True(node.HasChildNodes())
	assert.False(node.Contains(&child1))
	assert.True(node.Contains(&child2))
	assert.Len(node.ChildNodes(), 1)

	// Append child1
	node.AppendChild(&child1)
	assert.True(node.HasChildNodes())
	assert.True(node.Contains(&child1))
	assert.True(node.Contains(&child2))
	assert.Len(node.ChildNodes(), 2)
	assert.True(node.FirstChild().Equals(&child2))
	assert.True(node.LastChild().Equals(&child1))

	// Append child2
	node.AppendChild(&child2)
	assert.True(node.HasChildNodes())
	assert.True(node.Contains(&child1))
	assert.True(node.Contains(&child2))
	assert.Len(node.ChildNodes(), 2)
	assert.True(node.FirstChild().Equals(&child1))
	assert.True(node.LastChild().Equals(&child2))
}

func TestNode_005(t *testing.T) {
	assert := assert.New(t)

	// Create a parent
	node := newNode(nil, nil, "#node1", UNKNOWN_NODE, "")

	// Add three children to node
	child1 := newNode(nil, nil, "#child1", UNKNOWN_NODE, "child1")
	node.AppendChild(&child1)

	child2 := newNode(nil, nil, "#child2", UNKNOWN_NODE, "child2")
	node.AppendChild(&child2)

	child3 := newNode(nil, nil, "#child3", UNKNOWN_NODE, "child3")
	node.AppendChild(&child3)

	// Check children are added
	assert.True(node.HasChildNodes())
	assert.True(node.Contains(&child1))
	assert.True(node.Contains(&child2))
	assert.True(node.Contains(&child3))
	assert.Len(node.ChildNodes(), 3)

	// Insert child1 at the end
	assert.Equal(node.InsertBefore(&child1, nil), &child1)
	assert.True(node.LastChild().Equals(&child1))
	assert.True(child1.PreviousSibling().Equals(&child3))
	assert.Nil(child1.NextSibling())
	assert.True(child3.NextSibling().Equals(&child1))

	// Insert child2 at the end
	assert.Equal(node.InsertBefore(&child2, nil), &child2)
	assert.True(node.LastChild().Equals(&child2))
	assert.True(child2.PreviousSibling().Equals(&child1))
	assert.Nil(child2.NextSibling())
	assert.True(child1.NextSibling().Equals(&child2))

	// Insert child3 at the end
	assert.Equal(node.InsertBefore(&child3, nil), &child3)
	assert.True(node.LastChild().Equals(&child3))
	assert.True(child3.PreviousSibling().Equals(&child2))
	assert.Nil(child3.NextSibling())
	assert.True(child2.NextSibling().Equals(&child3))

	// Insert child3 at the beginning
	assert.Equal(node.InsertBefore(&child3, node.FirstChild()), &child3)
	assert.True(node.FirstChild().Equals(&child3))
	assert.True(child3.NextSibling().Equals(&child1))
	assert.Nil(child3.PreviousSibling())
	assert.True(child1.PreviousSibling().Equals(&child3))

	// Insert child2 at the beginning
	assert.Equal(node.InsertBefore(&child2, node.FirstChild()), &child2)
	assert.True(node.FirstChild().Equals(&child2))
	assert.True(child2.NextSibling().Equals(&child3))
	assert.Nil(child2.PreviousSibling())
	assert.True(child2.NextSibling().Equals(&child3))

	// Insert child1 at the beginning
	assert.Equal(node.InsertBefore(&child1, node.FirstChild()), &child1)
	assert.True(node.FirstChild().Equals(&child1))
	assert.True(child1.NextSibling().Equals(&child2))
	assert.Nil(child1.PreviousSibling())
	assert.True(child1.NextSibling().Equals(&child2))
}
