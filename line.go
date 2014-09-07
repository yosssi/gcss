package gcss

import (
	"fmt"
	"strings"
)

const unicodeSpace = 32

const indentTop = 0

// line represents a line of codes.
type line struct {
	no     int
	s      string
	indent int
}

// isEmpty returns true if the line's s is zero value.
func (ln *line) isEmpty() bool {
	return strings.TrimSpace(ln.s) == ""
}

// isTopIndent returns true if the line's indent is the top level.
func (ln *line) isTopIndent() bool {
	return ln.indent == indentTop
}

func (ln *line) childOf(parent element) (bool, error) {
	var ok bool
	var err error

	switch pIndent := parent.Base().ln.indent; {
	case ln.indent == pIndent+1:
		ok = true
	case ln.indent > pIndent+1:
		err = fmt.Errorf("indent is invalid [line: %d]", ln.no)
	}

	return ok, err
}

// newLine creates and returns a line.
func newLine(no int, s string) *line {
	return &line{
		no:     no,
		s:      s,
		indent: indent(s),
	}
}

// indent returns the string's indent.
func indent(s string) int {
	var i int

	for _, b := range s {
		if b != unicodeSpace {
			break
		}
		i++
	}

	return i / 2
}
