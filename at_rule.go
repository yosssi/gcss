package gcss

import "io"

// atRule represents an at-rule of CSS.
type atRule struct {
	elementBase
}

// WriteTo writes the at-rule to the writer.
func (ar *atRule) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// newAtRule creates and returns a at-rule.
func newAtRule(ln *line, parent element) *atRule {
	return &atRule{
		elementBase: newElementBase(ln, parent),
	}
}
