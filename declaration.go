package gcss

import (
	"fmt"
	"io"
	"strings"
)

// declaration represents a declaration of CSS.
type declaration struct {
	elementBase
	property string
	value    string
}

// WriteTo writes the declaration to the writer.
func (dec *declaration) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// AppendChild does nothing.
func (dec *declaration) AppendChild(child element) error {
	return nil
}

// declarationPV extracts a declaration property and value
// from the line.
func declarationPV(ln *line) (string, string, error) {
	pv := strings.SplitN(strings.TrimSpace(ln.s), space, 2)

	if len(pv) < 2 {
		return "", "", fmt.Errorf("declaration's property and value should be divided by a space [line: %d]", ln.no)
	}

	if !strings.HasSuffix(pv[0], colon) {
		return "", "", fmt.Errorf("property should end with a colon [line: %d]", ln.no)
	}

	return pv[0], pv[1], nil
}

// newDeclaration creates and returns a declaration.
func newDeclaration(ln *line, parent element) (*declaration, error) {
	property, value, err := declarationPV(ln)

	if err != nil {
		return nil, err
	}

	return &declaration{
		elementBase: newElementBase(ln, parent),
		property:    property,
		value:       value,
	}, nil
}
