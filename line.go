package gcss

// line represents a line of codes.
type line struct {
	no int
	s  string
}

// newLine creates and returns a line.
func newLine(no int, s string) *line {
	return &line{
		no: no,
		s:  s,
	}
}
