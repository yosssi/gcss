package gcss

import "testing"

func Test_newDeclaration(t *testing.T) {
	ln := newLine(1, "html")

	_, err := newDeclaration(ln, nil)

	if err == nil {
		t.Error("error should be occurred")
	}

	if expected := "declaration's property and value should be divided by a space [line: 1]"; expected != err.Error() {
		t.Errorf("err should be %q [actual: %q]", expected, err.Error())
	}
}
