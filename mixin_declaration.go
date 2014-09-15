package gcss

import (
	"fmt"
	"io"
	"strings"
)

// mixinDeclaration represents a mixin declaration.
type mixinDeclaration struct {
	elementBase
	name       string
	paramNames []string
}

// WriteTo writes the selector to the writer.
func (md *mixinDeclaration) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}

// mixinDeclarationNP extracts a mixin name and parameter names
// from the line.
func mixinDeclarationNP(ln *line) (string, []string, error) {
	s := strings.TrimSpace(ln.s)

	if !strings.HasPrefix(s, dollarMark) {
		return "", nil, fmt.Errorf("mixin declaration must start with %q [line: %d]", dollarMark, ln.no)
	}

	s = strings.TrimPrefix(s, dollarMark)

	np := strings.Split(s, openParenthesis)

	if len(np) != 2 {
		return "", nil, fmt.Errorf("mixin declaration's format is invalid [line: %d]", ln.no)
	}

	paramNamesS := strings.TrimSpace(np[1])

	if !strings.HasSuffix(paramNamesS, closeParenthesis) {
		return "", nil, fmt.Errorf("mixin declaration must end with %q [line: %d]", closeParenthesis, ln.no)
	}

	paramNamesS = strings.TrimSuffix(paramNamesS, closeParenthesis)

	if strings.Index(paramNamesS, closeParenthesis) != -1 {
		return "", nil, fmt.Errorf("mixin declaration's format is invalid [line: %d]", ln.no)
	}

	var paramNames []string

	if paramNamesS != "" {
		paramNames = strings.Split(paramNamesS, comma)
	}

	for i, paramName := range paramNames {
		paramName = strings.TrimSpace(paramName)

		if !strings.HasPrefix(paramName, dollarMark) {
			return "", nil, fmt.Errorf("mixin declaration's parameter must start with %q [line: %d]", dollarMark, ln.no)
		}

		paramNames[i] = strings.TrimPrefix(paramName, dollarMark)
	}

	return np[0], paramNames, nil
}

// newMixinDeclaration creates and returns a mixin declaration.
func newMixinDeclaration(ln *line, parent element) (*mixinDeclaration, error) {
	name, paramNames, err := mixinDeclarationNP(ln)

	if err != nil {
		return nil, err
	}

	return &mixinDeclaration{
		elementBase: newElementBase(ln, parent),
		name:        name,
		paramNames:  paramNames,
	}, nil
}
