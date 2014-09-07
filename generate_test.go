package gcss

import "testing"

func Test_generate(t *testing.T) {
	sc, errc := generate(nil)
	select {
	case <-sc:
	case <-errc:
	}
}
