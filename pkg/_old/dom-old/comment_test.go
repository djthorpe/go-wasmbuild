package dom_test

import (
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Comment_001(t *testing.T) {
	comment := GetWindow().Document().CreateComment("test")
	if comment.NodeType() != dom.COMMENT_NODE {
		t.Error("Expected COMMENT_NODE")
	} else if comment.NodeName() != "#comment" {
		t.Errorf("Expected #comment, got %q", comment.NodeName())
	} else {
		t.Log(comment)
	}
}
