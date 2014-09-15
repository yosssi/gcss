package gcss

import "io"

// mixinInvocation represents a mixin invocation.
type mixinInvocation struct {
	elementBase
	name        string
	paramValues []string
}

// WriteTo writes the selector to the writer.
func (mi *mixinInvocation) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newMixinInvocation creates and returns a mixin invocation.
func newMixinInvocation(ln *line, parent element) (*mixinInvocation, error) {
	name, paramValues, err := mixinNP(ln, false)

	if err != nil {
		return nil, err
	}

	return &mixinInvocation{
		elementBase: newElementBase(ln, parent),
		name:        name,
		paramValues: paramValues,
	}, nil
}
