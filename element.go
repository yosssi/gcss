package gcss

import "io"

// element represents an element of GCSS source codes.
type element interface {
	io.WriterTo
	AppendChild(child element) error
	Base() *elementBase
	SetContext(*context)
	Context() *context
}

// newElement creates and returns an element.
func newElement(ln *line, parent element) (element, error) {
	var e element
	var err error

	switch {
	case ln.isAtRule():
		e = newAtRule(ln, parent)
	case ln.isVariable():
		e, err = newVariable(ln, parent)
	case ln.isDeclaration():
		e, err = newDeclaration(ln, parent)
	default:
		e, err = newSelector(ln, parent)
	}

	return e, err
}
