package gcss

import (
	"io/ioutil"
	"testing"
)

func Test_parse(t *testing.T) {
	data, err := ioutil.ReadFile("./test/1.gcss")
	if err != nil {
		t.Errorf("error occurred [error: %s]", err.Error())
	}

	elemsc, errc := parse(string(data))

	select {
	case <-elemsc:
	case <-errc:
		t.Errorf("error occurred [error: %s]", err.Error())
	}
}

func Test_formatLF(t *testing.T) {
	s := cr + crlf + lf + "a" + crlf + lf + cr + "b" + lf + cr + crlf
	expectedS := lf + lf + lf + "a" + lf + lf + lf + "b" + lf + lf + lf

	if formatLF(s) != expectedS {
		t.Errorf("formatLF(s) should be %s [actual: %s]", expectedS, formatLF(s))
	}
}
