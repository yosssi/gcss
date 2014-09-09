package gcss

import (
	"fmt"
	"io"
	"strings"
)

// selector represents a selector of CSS.
type selector struct {
	elementBase
	name string
	sels []*selector
	decs []*declaration
}

// WriteTo writes the selector to the writer.
func (sel *selector) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// AppendChild appends a selector or declaration to the selector.
func (sel *selector) AppendChild(child element) error {
	switch child.(type) {
	case *selector:
		sel.sels = append(sel.sels, child.(*selector))
	case *declaration:
		sel.decs = append(sel.decs, child.(*declaration))
	default:
		return fmt.Errorf("invalid child's type [line: %d]", sel.ln.no)
	}

	return nil
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
		name:        strings.TrimSpace(ln.s),
	}
}
