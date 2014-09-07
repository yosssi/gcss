package gcss

import "testing"

func Test_newSelector(t *testing.T) {
	ln := newLine(1, "html")

	sel := newSelector(ln, nil)

	if sel.ln != ln {
		t.Errorf("sel.ln should be %+v [actual: %+v]", ln, sel.ln)
	}
}
