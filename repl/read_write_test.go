package repl

import (
	"fmt"
	"strings"
	"testing"
)

func TestLineReader(t *testing.T) {
	input := strings.NewReader("Hello World\nGoodbye\n")
	lines := ReadLinesFrom(input)

	first := <-lines
	if first != "Hello World" {
		fmt.Printf("Failed to read first line from line reader. Got: %s\n", first)
		t.Fail()
	}

	second := <-lines
	if second != "Goodbye" {
		fmt.Printf("Failed to read second line from line reader. Got: %s\n", second)
		t.Fail()
	}
}

func TestLineWriter(t *testing.T) {
	output := &strings.Builder{}
	lines, done := WriteLinesTo(output)

	lines <- "Hello World"
	lines <- "Goodbye"
	lines <- ExitMessage
	close(lines)

	<-done // wait for the writer to finish

	result := output.String()
	expectedResult := fmt.Sprintf("> Hello World\n> Goodbye\n> %s\n", ExitMessage)
	if result != expectedResult {
		fmt.Printf("Failed to write lines using line writer. Got: %s\n", result)
		t.Fail()
	}

}
