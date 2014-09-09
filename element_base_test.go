package gcss

import "testing"

func Test_elementBase_AppendChild(t *testing.T) {
	parent, err := newElement(newLine(1, "html"), nil)
	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	child, err := newElement(newLine(2, "  font-size: 12px;"), parent)
	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	parent.Base().AppendChild(child)

	if expected, actual := 1, len(parent.Base().children); actual != expected {
		t.Errorf("len(parent.Base().children) should be %d [actual: %d]", expected, actual)
	}

	if parent.Base().children[0] != child {
		t.Errorf("parent.Base().children[0] should be %+v [actual: %+v]", child, parent.Base().children[0])
	}
}

func Test_elementBase_Base(t *testing.T) {
	e, err := newElement(newLine(1, "html"), nil)
	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	if e.Base() == nil {
		t.Error("e.Base() should not be nil")
	}
}

func Test_newElementBase(t *testing.T) {
	ln := newLine(1, "html")

	eBase := newElementBase(ln, nil)

	if eBase.ln != ln {
		t.Errorf("eBase.ln should be %+v [actual: %+v]", ln, eBase.ln)
	}
}
