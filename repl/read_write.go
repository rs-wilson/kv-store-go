package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLinesFrom(source io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(source)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()
	return out
}

func WriteLinesTo(destination io.Writer) (chan<- string, <-chan struct{}) {
	in := make(chan string)
	done := make(chan struct{})
	go func() {
		defer close(done)
		writer := bufio.NewWriter(destination)
		writer.WriteString("> ")
		writer.Flush()

		// Loop on all lines until the channel is closed by the writer
		for line := range in {
			if len(line) != 0 {
				if _, err := writer.WriteString(line + "\n"); err != nil {
					fmt.Fprintf(os.Stderr, "Failed to write line to destination: %s\n", err.Error())
					return
				}
			}

			// write next carrot if we're not exiting
			if line != ExitMessage {
				writer.WriteString("> ")
			}

			// Ensure lines get immediately written
			if err := writer.Flush(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to flush write destination: %s\n", err.Error())
				return
			}
		}
	}()
	return in, done
}
