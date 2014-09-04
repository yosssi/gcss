package gcss

// Compile parses the GCSS file specified by the path parameter,
// generates a CSS file and returns the two channels: the first
// one returns the path of the generated CSS file and the last
// one returns an error when it occurs.
func Compile(path string) (<-chan string, <-chan error) {
	pathc := make(chan string)
	errc := make(chan error)

	return pathc, errc
}

// CompileString parses the GCSS string passed as the s parameter,
// generates a CSS string and returns the two channels: the first
// one returns the CSS string and the last one returns an error
// when it occurs.
func CompileString(s string) (<-chan string, <-chan error) {
	sc := make(chan string)
	errc := make(chan error)

	return sc, errc
}
