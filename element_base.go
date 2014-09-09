package gcss

// elementBase holds the common fields of an element.
type elementBase struct {
	ln     *line
	parent element
}

// Base returns the element base.
func (eBase *elementBase) Base() *elementBase {
	return eBase
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line, parent element) elementBase {
	return elementBase{
		ln:     ln,
		parent: parent,
	}
}
