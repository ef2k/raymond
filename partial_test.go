package raymond

import (
	"testing"
)

func TestRemovePartial(t *testing.T) {
	RegisterPartial("foo", "<h1>foo</h1>")
	if _, ok := partials["foo"]; !ok {
		t.Error("Expected partial to be registered")
	}

	if err := RegisterPartial("foo", "<h1>foo</h1>"); err == nil {
		t.Error("Expected registering the same partial to return an error")
	}

	RemovePartial("foo")
	if _, ok := partials["foo"]; ok {
		t.Error("Expected partial to be removed")
	}
}

func TestRegisterPartials(t *testing.T) {
	partials := map[string]string{
		"foo": "<h1>foo</h1>",
	}
	if err := RegisterPartials(partials); err != nil {
		t.Error("Expected partial to be registered without error")
	}
	if err := RegisterPartials(partials); err == nil {
		t.Error("Expected registering the same partial to return an error")
	}
}
