package gcss

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// cssFilePath creates and returns the CSS file path.
var cssFilePath = func(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext) + extCSS
}

// Compile parses the GCSS file specified by the path parameter,
// generates a CSS file and returns the two channels: the first
// one returns the path of the generated CSS file and the last
// one returns an error when it occurs.
func Compile(path string) (<-chan string, <-chan error) {
	pathc := make(chan string)
	errc := make(chan error)

	go func() {
		data, err := ioutil.ReadFile(path)

		if err != nil {
			errc <- err
			return
		}

		bc, cErrc := CompileBytes(data)

		select {
		case b := <-bc:
			cssPath := cssFilePath(path)

			done, wErrc := write(cssPath, b)

			select {
			case <-done:
				pathc <- cssPath
			case err := <-wErrc:
				errc <- err
				return
			}
		case err := <-cErrc:
			errc <- err
			return
		}
	}()

	return pathc, errc
}

// CompileBytes parses the GCSS string passed as the s parameter,
// generates a CSS string and returns the two channels: the first
// one returns the CSS string and the last one returns an error
// when it occurs.
func CompileBytes(b []byte) (<-chan []byte, <-chan error) {
	bc := make(chan []byte)
	errc := make(chan error)

	go func() {
		elemc, pErrc := parse(string(b))

		bf := new(bytes.Buffer)

		for {
			select {
			case elem, ok := <-elemc:
				if elem != nil {
					elem.WriteTo(bf)
				}

				if ok {
					continue
				}

				bc <- bf.Bytes()

				return
			case err := <-pErrc:
				errc <- err
				return
			}
		}
	}()

	return bc, errc
}
