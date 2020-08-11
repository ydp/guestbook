package http

import (
	"reflect"
	"testing"
)

func Test_createGuestbook(t *testing.T) {
	g := createGuestbook([]Signature{})
	if g.SignatureCount != 0 {
		t.Error("signature count should be 0")
	}
	if reflect.DeepEqual(g.Signatures, []string{}) {
		t.Error("signatures should be empty")
	}
}
