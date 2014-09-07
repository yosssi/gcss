package gcss

func generate(elems []element) (<-chan string, <-chan error) {
	sc := make(chan string)
	errc := make(chan error)

	go func() {
		sc <- ""
	}()
	return sc, errc
}
