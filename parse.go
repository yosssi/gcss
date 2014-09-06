package gcss

import (
	"fmt"
	"strings"
)

// Special characters
const (
	cr   = "\r"
	lf   = "\n"
	crlf = "\r\n"
)

// parse parses the string, generates the elements
// and returns the two channels: the first one returns
// the generated elements and the last one returns
// an error when it occurs.
func parse(s string) (<-chan []element, <-chan error) {
	elemsc := make(chan []element)
	errc := make(chan error)

	go func() {
		var elements []element

		lines := strings.Split(formatLF(s), lf)

		i := 0
		l := len(lines)

		for i < l {
			// Fetch a line.
			ln := newLine(i+1, lines[i])
			i++

			// Ignore the empty line.
			if strings.TrimSpace(ln.s) == "" {
				continue
			}

			fmt.Println(ln.s)
		}

		elemsc <- elements
	}()

	return elemsc, errc
}

// formatLF replaces the line feed codes with LF and
// returns the result string.
func formatLF(s string) string {
	return strings.Replace(strings.Replace(s, crlf, lf, -1), cr, lf, -1)
}
