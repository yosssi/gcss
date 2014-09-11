package gcss

import (
	"bytes"
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
func (sel *selector) WriteTo(w io.Writer) (int64, error) {
	bf := new(bytes.Buffer)

	// Write the declarations.
	if len(sel.decs) > 0 {
		bf.WriteString(sel.names())
		bf.WriteString(openBrace)

		for _, dec := range sel.decs {
			bf.WriteString(dec.property)
			bf.WriteString(colon)
			bf.WriteString(dec.value)
			bf.WriteString(semicolon)
		}

		bf.WriteString(closeBrace)
	}

	// Write the child selectors.
	for _, childSel := range sel.sels {
		// Writing to the bytes.Buffer never returns an error.
		childSel.WriteTo(bf)
	}

	n, err := w.Write(bf.Bytes())

	return int64(n), err
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

// names returns the selector names.
func (sel *selector) names() string {
	bf := new(bytes.Buffer)

	switch sel.parent.(type) {
	case nil:
		for _, name := range strings.Split(sel.name, comma) {
			if bf.Len() > 0 {
				bf.WriteString(comma)
			}

			bf.WriteString(strings.TrimSpace(name))
		}
	case *selector:
		for _, parentS := range strings.Split(sel.parent.(*selector).names(), comma) {
			for _, s := range strings.Split(sel.name, comma) {
				if bf.Len() > 0 {
					bf.WriteString(comma)
				}

				bf.WriteString(parentS)

				s = strings.TrimSpace(s)

				if strings.HasPrefix(s, ampersand) {
					bf.WriteString(strings.TrimPrefix(s, ampersand))
				} else {
					bf.WriteString(space)
					bf.WriteString(s)
				}
			}
		}
	}

	return bf.String()
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
		name:        strings.TrimSpace(ln.s),
	}
}
