package gcss

import "io"

// selector represents a selector.
type selector struct {
	elementBase
}

func (sel *selector) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
	}
}
