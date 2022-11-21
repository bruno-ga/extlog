//go:build extlog_enabled

package extlog

import (
	"os"
	"testing"
)

func TestAll(t *testing.T) {
	l := New(os.Stderr, "Test", 0)

	l.PrintlnEveryN(5, " should not print")

	for i := 0; i < 100; i++ {
		l.PrintlnEveryN(20, " should print", i)
	}

	l.PrintlnEveryN(1, " should print")
}
