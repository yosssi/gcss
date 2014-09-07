package gcss

// element represents an element of GCSS source codes.
type element interface {
	AppendChild(child element)
	Base() *elementBase
}

// newElement creates and returns an element.
func newElement(ln *line, parent element) element {
	var e element

	switch {
	default:
		e = newSelector(ln, parent)
	}

	return e
}
