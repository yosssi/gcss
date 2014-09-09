package gcss

import "io"

// selector represents a selector of CSS.
type selector struct {
	elementBase
}

// WriteTo write the selector to the writer.
func (sel *selector) WriteTo(w io.Writer) (n int64, err error) {
	for _, e := range sel.children {
		e.WriteTo(w)
	}
	return 0, nil
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
	}
}
