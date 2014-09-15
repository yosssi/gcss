package gcss

import (
	"io/ioutil"
	"testing"
)

func Test_mixinDeclaration_WriteTo(t *testing.T) {
	ln := newLine(1, "$test()")

	md, err := newMixinDeclaration(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}

	md.WriteTo(ioutil.Discard)
}

func Test_mixinDeclaration_mixinDeclarationNP_errPrefixDollar(t *testing.T) {
	ln := newLine(1, "test()")

	_, _, err := mixinDeclarationNP(ln)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration must start with \"$\" [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_mixinDeclaration_mixinDeclarationNP_errNoOpenParenthesis(t *testing.T) {
	ln := newLine(1, "$test")

	_, _, err := mixinDeclarationNP(ln)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration's format is invalid [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_mixinDeclaration_mixinDeclarationNP_errNoCloseParenthesis(t *testing.T) {
	ln := newLine(1, "$test(")

	_, _, err := mixinDeclarationNP(ln)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration must end with \")\" [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_mixinDeclaration_mixinDeclarationNP_errMultiCloseParentheses(t *testing.T) {
	ln := newLine(1, "$test())")

	_, _, err := mixinDeclarationNP(ln)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration's format is invalid [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_mixinDeclaration_mixinDeclarationNP_noParamNames(t *testing.T) {
	ln := newLine(1, "$test()")

	_, _, err := mixinDeclarationNP(ln)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func Test_mixinDeclaration_mixinDeclarationNP(t *testing.T) {
	ln := newLine(1, "$test($param1)")

	_, _, err := mixinDeclarationNP(ln)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func Test_mixinDeclaration_mixinDeclarationNP_errInvalidParamNames(t *testing.T) {
	ln := newLine(1, "$test(param1)")

	_, _, err := mixinDeclarationNP(ln)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration's parameter must start with \"$\" [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_newMixinDeclaration(t *testing.T) {
	ln := newLine(1, "$test($param1, $param2)")

	_, err := newMixinDeclaration(ln, nil)

	if err != nil {
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func Test_newMixinDeclaration_errInvalidParamNames(t *testing.T) {
	ln := newLine(1, "$test(param1)")

	_, err := newMixinDeclaration(ln, nil)

	if err == nil {
		t.Error("error should occur")
	}

	if expected, actual := "mixin declaration's parameter must start with \"$\" [line: 1]", err.Error(); actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_newMixinDeclaration_fromFile(t *testing.T) {
	pathc, errc := Compile("test/0014.gcss")

	select {
	case <-pathc:
	case err := <-errc:
		t.Error("error occurred [error: %q]", err.Error())
	}
}
