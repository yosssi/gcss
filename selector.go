package gcss

// selector represents a selector.
type selector struct {
	elementBase
}

// newSelector creates and returns a selector.
func newSelector(ln *line, parent element) *selector {
	return &selector{
		elementBase: newElementBase(ln, parent),
	}
}
