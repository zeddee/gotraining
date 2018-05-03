package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger …
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New …
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

// Close …
func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

// Println
func (l *Logger) Println(v string) {
	select {
	// performs a channel send on a channel
	case l.ch <- v:
		// If you can send a message to a channel, ok
	default:
		// If cannot send to a channel, "DROP"
		fmt.Println("DROP")
	}
}
