package gcss

import (
	"bufio"
	"io"
	"os"
)

const extCSS = ".css"

// WriteFlusher is the interface that groups the basic Write and Flush methods.
type WriteFlusher interface {
	io.Writer
	Flush() error
}

var newBufWriter = func(w io.Writer) WriteFlusher {
	return bufio.NewWriter(w)
}

// write writes the string "s" to the CSS file.
func write(path string, bc <-chan []byte, berrc <-chan error) (<-chan struct{}, <-chan error) {
	done := make(chan struct{})
	errc := make(chan error)

	go func() {
		f, err := os.Create(path)

		if err != nil {
			errc <- err
			return
		}

		defer f.Close()

		w := newBufWriter(f)

		for {
			select {
			case b, ok := <-bc:
				if _, err := w.Write(b); err != nil {
					errc <- err
					return
				}

				if !ok {
					if err := w.Flush(); err != nil {
						errc <- err
						return
					}

					done <- struct{}{}

					return
				}
			case err := <-berrc:
				errc <- err
				return
			}
		}
	}()

	return done, errc
}
