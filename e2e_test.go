package gcss

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func Test_e2e(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 1; i++ {
		idx := strconv.Itoa(i)

		gcssPath := "test/e2e/actual/" + strings.Repeat("0", 4-len(idx)) + idx + ".gcss"

		wg.Add(1)

		go func() {
			pc, errc := Compile(gcssPath)

			select {
			case actualCSSPath := <-pc:
				expectedCSSPath := strings.Replace(actualCSSPath, "actual", "expected", -1)

				actualB, err := ioutil.ReadFile(actualCSSPath)

				if err != nil {
					t.Errorf("error occurred [error: %q]", err.Error())
				}

				expectedB, err := ioutil.ReadFile(expectedCSSPath)

				if err != nil {
					t.Errorf("error occurred [error: %q]", err.Error())
				}

				if len(actualB) != len(expectedB) {
					t.Errorf("actual result does not match the expected result [path: %q]", gcssPath)
				}

				for i, b := range actualB {
					if b != expectedB[i] {
						t.Errorf("actual result does not match the expected result [path: %q]", gcssPath)
					}
				}

				wg.Done()
			case err := <-errc:
				t.Errorf("error occurred [error: %q]", err.Error())
			}
		}()

		wg.Wait()
	}
}
