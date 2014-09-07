package gcss

import (
	"io/ioutil"
	"testing"
)

func Test_selector_WriteTo(t *testing.T) {
	ln := newLine(1, "html")

	sel := newSelector(ln, nil)

	_, err := sel.WriteTo(ioutil.Discard)

	if err != nil {
		t.Errorf("err should be nil [err: %s]", err.Error())
	}
}

func Test_newSelector(t *testing.T) {
	ln := newLine(1, "html")

	sel := newSelector(ln, nil)

	if sel.ln != ln {
		t.Errorf("sel.ln should be %+v [actual: %+v]", ln, sel.ln)
	}
}
