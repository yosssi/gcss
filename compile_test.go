package gcss

import "testing"

func TestCompile_readFileErr(t *testing.T) {
	pathc, errc := Compile("not_exist_file")

	select {
	case <-pathc:
		t.Error("error should be occurred")
	case err := <-errc:
		if expected, actual := "open not_exist_file: no such file or directory", err.Error(); expected != actual {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
		}
	}
}

func TestCompile_compileStringErr(t *testing.T) {
	pathc, errc := Compile("test/4.gcss")

	select {
	case <-pathc:
		t.Error("error should be occurred")
	case err := <-errc:
		if expected, actual := "indent is invalid [line: 5]", err.Error(); expected != actual {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
		}
	}
}

func TestCompile_writeErr(t *testing.T) {
	cssFileBack := cssFilePath

	cssFilePath = func(_ string) string {
		return "not_exist_dir/not_exit_file"
	}

	pathc, errc := Compile("test/3.gcss")

	select {
	case <-pathc:
		t.Error("error should be occurred")
	case err := <-errc:
		if expected, actual := "open not_exist_dir/not_exit_file: no such file or directory", err.Error(); expected != actual {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
		}
	}

	cssFilePath = cssFileBack
}

func TestCompile(t *testing.T) {
	pathc, errc := Compile("test/3.gcss")

	select {
	case path := <-pathc:
		if expected := "test/3.css"; expected != path {
			t.Errorf("path should be %q [actual: %q]", expected, path)
		}
	case err := <-errc:
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func TestCompile_pattern2(t *testing.T) {
	pathc, errc := Compile("test/7.gcss")

	select {
	case path := <-pathc:
		if expected := "test/7.css"; expected != path {
			t.Errorf("path should be %q [actual: %q]", expected, path)
		}
	case err := <-errc:
		t.Errorf("error occurred [error: %q]", err.Error())
	}
}

func TestComplieBytes(t *testing.T) {
	CompileBytes([]byte(""))
}
