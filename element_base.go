package gcss

// elementBase holds the common fields of an element.
type elementBase struct {
	ln     *line
	parent element
	sels   []*selector
	decs   []*declaration
	mixins []*mixinInvocation
	ctx    *context
}

// AppendChild appends a child element to the element.
func (eBase *elementBase) AppendChild(child element) {
	switch child.(type) {
	case *mixinInvocation:
		eBase.mixins = append(eBase.mixins, child.(*mixinInvocation))
	case *declaration:
		eBase.decs = append(eBase.decs, child.(*declaration))
	case *selector:
		eBase.sels = append(eBase.sels, child.(*selector))
	}
}

// Base returns the element base.
func (eBase *elementBase) Base() *elementBase {
	return eBase
}

// SetContext sets the context to the element.
func (eBase *elementBase) SetContext(ctx *context) {
	eBase.ctx = ctx
}

// Context returns the top element's context.
func (eBase *elementBase) Context() *context {
	if eBase.parent != nil {
		return eBase.parent.Context()
	}

	return eBase.ctx
}

// hasMixinDecs returns true if the selector has a mixin
// which has declarations.
func (eBase *elementBase) hasMixinDecs() bool {
	for _, mi := range eBase.mixins {
		if decs, _ := mi.decsParams(); len(decs) > 0 {
			return true
		}
	}

	return false
}

// newElementBase creates and returns an element base.
func newElementBase(ln *line, parent element) elementBase {
	return elementBase{
		ln:     ln,
		parent: parent,
	}
}
