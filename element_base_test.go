package gcss

import "testing"

func Test_newElementBase(t *testing.T) {
	ln := newLine(1, "html")

	eBase := newElementBase(ln, nil)

	if eBase.ln != ln {
		t.Errorf("eBase.ln should be %+v [actual: %+v]", ln, eBase.ln)
	}
}
