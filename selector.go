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
		names, err := sel.names()

		if err != nil {
			return 0, err
		}

		if _, err := bf.WriteString(names + openBrace); err != nil {
			return 0, err
		}

		for _, dec := range sel.decs {
			if _, err := bf.WriteString(dec.property + colon + dec.value + semicolon); err != nil {
				return 0, err
			}
		}

		if _, err := bf.WriteString(closeBrace); err != nil {
			return 0, err
		}
	}

	// Write the child selectors.
	for _, childSel := range sel.sels {
		if _, err := childSel.WriteTo(bf); err != nil {
			return 0, err
		}
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
func (sel *selector) names() (string, error) {
	if sel.parent == nil {
		names := strings.Split(sel.name, comma)

		for i, name := range names {
			names[i] = strings.TrimSpace(name)
		}

		return strings.Join(names, comma), nil
	}

	bf := new(bytes.Buffer)

	names, err := sel.parent.(*selector).names()

	if err != nil {
		return "", err
	}

	for _, parentS := range strings.Split(names, comma) {
		for _, s := range strings.Split(sel.name, comma) {
			s = strings.TrimSpace(s)

			if strings.HasPrefix(s, ampersand) {
				s = strings.TrimPrefix(s, ampersand)
			} else {
				s = space + s
			}

			if bf.Len() > 0 {
				if _, err := bf.WriteString(comma); err != nil {
					return "", err
				}
			}

			if _, err := bf.WriteString(parentS + s); err != nil {
				return "", err
			}
		}
	}

	return bf.String(), nil
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
		name:        strings.TrimSpace(ln.s),
	}
}
