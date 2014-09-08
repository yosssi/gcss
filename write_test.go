package gcss

import "testing"

func Test_writeErr(t *testing.T) {
	done, errc := write("not_exist_dir/not_exit_file", nil)

	select {
	case <-done:
		t.Error("error should be occurred")
	case err := <-errc:
		if expected, actual := "open not_exist_dir/not_exit_file: no such file or directory", err.Error(); expected != actual {
			t.Errorf("err should be %q [actual: %q]", expected, actual)
		}
	}
}
