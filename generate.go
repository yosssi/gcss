package gcss

func generate(elem element) (<-chan string, <-chan error) {
	sc := make(chan string)
	errc := make(chan error)

	go func() {
		sc <- ""
	}()
	return sc, errc
}
