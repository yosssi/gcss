package gcss

import (
	"io/ioutil"
	"testing"
)

func Test_atRule_WriteTo(t *testing.T) {
	ln := newLine(1, "html")

	ar := newAtRule(ln, nil)

	if _, err := ar.WriteTo(ioutil.Discard); err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func Test_atRule_WriteTo_fromFile(t *testing.T) {
	pathc, errc := Compile("test/9.gcss")

	select {
	case <-pathc:
	case err := <-errc:
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func Test_newAtRule(t *testing.T) {
	ln := newLine(1, "html")

	ar := newAtRule(ln, nil)

	if ar.ln != ln {
		t.Errorf("ar.ln should be %+v [actual: %+v]", ln, ar.ln)
	}
}
