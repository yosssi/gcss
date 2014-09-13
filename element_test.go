package gcss

import (
	"reflect"
	"testing"
)

func Test_newElement_selector(t *testing.T) {
	ln := newLine(1, "html")

	e, err := newElement(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	if _, ok := e.(*selector); !ok {
		t.Errorf(`e's type should be "*selector" [actual: %q]`, reflect.TypeOf(e))
	}
}
