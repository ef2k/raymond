package raymond

import (
	"testing"
)

func TestRemovePartial(t *testing.T) {
	RegisterPartial("foo", "<h1>foo</h1>")
	if _, ok := partials["foo"]; !ok {
		t.Error("Expected partial to be registered")
	}

	RemovePartial("foo")
	if _, ok := partials["foo"]; ok {
		t.Error("Expected partial to be removed")
	}
}
