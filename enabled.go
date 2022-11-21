//go:build extlog_enabled

package extlog

import (
	"io"
	"log"
	"runtime"
)

// ExtLog is an extended log facility. It uses the standard log package under
// the hood. Its main feature is that it is a lot more expensive than the
// standard one but allows 2 interesting things:
//
//  1. Implements "...EveryN" method for each logging function
//  2. Allows for almost complete disabling of logs (as much as is) possible
//     without a macro-like facility being available.
//
// Note that most methods (except for the "...EveryN" ones are a direct)
// passthrough to the underlying log.Logger, which only adds an extra
// indirection to each call (which is not that bad).
type ExtLog struct {
	*log.Logger

	// Counts per PC.
	counters map[uintptr]int
}

// New returns a new ExtLog instance.
func New(out io.Writer, prefix string, flags int) *ExtLog {
	return &ExtLog{
		Logger:   log.New(out, prefix, flags),
		counters: make(map[uintptr]int),
	}
}

// PrintEveryN calls log.Print every Nth call of it in a specific call site.
func (l *ExtLog) PrintEveryN(n int, v ...any) {
	if l.updateCounter(n) == 0 {
		l.Print(v...)
	}
}

// PrintfEveryN calls log.Printf every Nth call of it in a specific call site.
func (l *ExtLog) PrintfEveryN(n int, format string, v ...any) {
	if l.updateCounter(n) == 0 {
		l.Printf(format, v...)
	}
}

// PrintlnEveryN calls log.Println every Nth call of it in a specific call site.
func (l *ExtLog) PrintlnEveryN(n int, v ...any) {
	if l.updateCounter(n) == 0 {
		l.Println(v...)
	}
}

func (l *ExtLog) updateCounter(n int) int {
	// Get relevant caller pc by skipping the top 2 stack frames (this function
	// and the actual logging function).
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		panic("Can not get Callers. I DO NOT KNOW WHAT TO DO!!!!1111one")
	}

	// Update counter for the specific PC. Should generally map to a specific
	// call site in a program (inlining probably breaks this).
	//
	// TODO(bga): This needs to be done atomically.
	l.counters[pc] = (l.counters[pc] + 1) % n

	return l.counters[pc]
}
