package gcss

import "io"

var newDeclarationFunc = newDeclaration

// element represents an element of GCSS source codes.
type element interface {
	io.WriterTo
	AppendChild(child element) error
	Base() *elementBase
}

// newElement creates and returns an element.
func newElement(ln *line, parent element) (element, error) {
	var e element
	var err error

	switch {
	case ln.isAtRule():
		e = newAtRule(ln, parent)
	case ln.isDeclaration():
		e, err = newDeclarationFunc(ln, parent)
		if err != nil {
			return nil, err
		}
	default:
		e = newSelector(ln, parent)
	}

	return e, nil
}
