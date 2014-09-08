package gcss

import (
	"io/ioutil"
	"os"
)

const extCSS = ".css"

// write writes the string "s" to the CSS file.
func write(path string, b []byte) (<-chan struct{}, <-chan error) {
	done := make(chan struct{})
	errc := make(chan error)

	go func() {
		if err := ioutil.WriteFile(path, b, os.ModePerm); err != nil {
			errc <- err
			return
		}

		done <- struct{}{}
	}()

	return done, errc
}
