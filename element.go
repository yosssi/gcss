package gcss

import "io"

// element represents an element of GCSS source codes.
type element interface {
	io.WriterTo
	AppendChild(child element) error
	Base() *elementBase
}

// newElement creates and returns an element.
func newElement(ln *line, parent element) element {
	var e element

	switch {
	case ln.isAtRule():
		e = newAtRule(ln, parent)
	case ln.isDeclaration():
		// newDeclaration never returns an error in this context
		// because ln is a valid declaration by executing
		// `ln.isDeclaration()` beforehand.
		e, _ = newDeclaration(ln, parent)
	default:
		e = newSelector(ln, parent)
	}

	return e
}
