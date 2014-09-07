package gcss

// element represents an element of GCSS source codes.
type element interface {
}

func newElement(ln *line, parent element) element {
	var e element

	switch {
	default:
		e = newSelector(ln, parent)
	}

	return e
}
