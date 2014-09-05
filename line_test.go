package gcss

import "testing"

func Test_newLine(t *testing.T) {
	no := 1
	s := "html"

	ln := newLine(no, s)

	if ln.no != no {
		t.Errorf("ln.no should be %d [actual: %d]", no, ln.no)
	}

	if ln.s != s {
		t.Errorf("ln.s should be %s [actual: %s]", s, ln.s)
	}
}
