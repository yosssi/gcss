package gcss

import (
	"reflect"
	"testing"
)

func Test_newElement_selector(t *testing.T) {
	ln := newLine(1, "html")

	e := newElement(ln, nil)

	if _, ok := e.(*selector); !ok {
		t.Errorf(`e's type should be "*selector" [actual: %q]`, reflect.TypeOf(e))
	}
}
