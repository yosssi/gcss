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
