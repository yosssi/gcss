package gcss

import (
	"errors"
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

func Test_newElement_declaration_err(t *testing.T) {
	newDeclarationFuncBak := newDeclarationFunc

	errMsg := "test error"

	newDeclarationFunc = func(ln *line, parent element) (*declaration, error) {
		return nil, errors.New(errMsg)
	}

	ln := newLine(1, "color: #000")

	_, err := newElement(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
	}

	if err.Error() != errMsg {
		t.Errorf("err should be %q [actual: %q]", errMsg, err.Error())
	}

	newDeclarationFunc = newDeclarationFuncBak
}
