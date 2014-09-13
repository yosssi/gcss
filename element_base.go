package gcss

import "fmt"

// elementBase holds the common fields of an element.
type elementBase struct {
	ln     *line
	parent element
	sels   []*selector
	decs   []*declaration
}

// AppendChild appends a selector or declaration to the selector.
func (eBase *elementBase) AppendChild(child element) error {
	switch child.(type) {
	case *selector:
		eBase.sels = append(eBase.sels, child.(*selector))
	case *declaration:
		eBase.decs = append(eBase.decs, child.(*declaration))
	default:
		return fmt.Errorf("invalid child's type [line: %d]", eBase.ln.no)
	}

	return nil
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
