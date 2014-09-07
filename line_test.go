package gcss

import "testing"

func Test_line_isEmpty_true(t *testing.T) {
	ln := newLine(1, "")

	if !ln.isEmpty() {
		t.Error("ln.Empty() should return true")
	}
}

func Test_line_isEmpty_false(t *testing.T) {
	ln := newLine(1, "html")

	if ln.isEmpty() {
		t.Error("ln.Empty() should return false")
	}
}

func Test_line_isTopIndent_true(t *testing.T) {
	ln := newLine(1, "html")

	if !ln.isTopIndent() {
		t.Error("ln.isTopIndent() should return true")
	}
}

func Test_line_isTopIndent_false(t *testing.T) {
	ln := newLine(1, "  html")

	if ln.isTopIndent() {
		t.Error("ln.isTopIndent() should return false")
	}
}

func Test_newLine(t *testing.T) {
	no := 1
	s := "  html"

	ln := newLine(no, s)

	if ln.no != no {
		t.Errorf("ln.no should be %d [actual: %d]", no, ln.no)
	}

	if ln.s != s {
		t.Errorf("ln.s should be %s [actual: %s]", s, ln.s)
	}

	if ln.indent != indent(s) {
		t.Errorf("ln.indent should be %d [actual: %d]", indent(s), ln.indent)
	}
}

func Test_indent_no_indent(t *testing.T) {
	s := "html"
	expected := 0
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}

func Test_indent_half_indent(t *testing.T) {
	s := "   html"
	expected := 1
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}

func Test_indent(t *testing.T) {
	s := "    html"
	expected := 2
	actual := indent(s)

	if actual != expected {
		t.Errorf("%q's indent should be %d [actual: %d]", s, expected, actual)
	}
}
