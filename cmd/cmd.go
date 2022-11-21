package main

import (
	"extlog"
	"os"
)

func main() {
	l := extlog.New(os.Stderr, "Test", 0)

	for i := 0; i < 10; i++ {
		l.PrintlnEveryN(5, "")
	}
}
