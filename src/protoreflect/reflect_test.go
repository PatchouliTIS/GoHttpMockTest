package protoreflect

import "testing"

func TestReflect(t *testing.T) {
	_, err := Reflect()
	if err != nil {
		t.Errorf("protoreflect error: %s", err)
	}
}
