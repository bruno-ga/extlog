//go:build !extlog_enabled

package extlog

import (
	"io"
)

// ExtLog is a disabled logger. It should be ellided by the compiler.
type ExtLog struct{}

// New returns a new disabled logger.
func New(out io.Writer, prefix string, flags int) *ExtLog {
	return nil
}

// Empty implementations of all log.Logger methods.

func (l *ExtLog) Fatal(v ...any)                 {}
func (l *ExtLog) Fatalf(format string, v ...any) {}
func (l *ExtLog) Fatalln(v ...any)               {}
func (l *ExtLog) Flags() int {
	return 0
}
func (l *ExtLog) Output(calldepth int, s string) error {
	return nil
}
func (l *ExtLog) Panic(v ...any)                 {}
func (l *ExtLog) Panicf(format string, v ...any) {}
func (l *ExtLog) Panicln(v ...any)               {}
func (l *ExtLog) Prefix() string {
	return ""
}
func (l *ExtLog) Print(v ...any)                 {}
func (l *ExtLog) Printf(format string, v ...any) {}
func (l *ExtLog) Println(v ...any)               {}
func (l *ExtLog) SetFlags(flag int)              {}
func (l *ExtLog) SetOutput(w io.Writer)          {}
func (l *ExtLog) SetPrefix(prefix string)        {}
func (l *ExtLog) Writer() io.Writer {
	return nil
}

// Empty implementation of all Extlog methods.

func (l *ExtLog) PrintEveryN(n int, v ...any)                 {}
func (l *ExtLog) PrintfEveryN(n int, format string, v ...any) {}
func (l *ExtLog) PrintlnEveryN(n int, v ...any)               {}
