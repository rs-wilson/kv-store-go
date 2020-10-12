package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestMain_Simple(t *testing.T) {
	main := exec.Command("./kv-store")

	input := `WRITE a hello
READ a
WRITE a hello-again
READ a
DELETE a
READ a
WRITE a once-more
READ a
QUIT
	`

	wanted := `> > hello
> > hello-again
> > Key not found: a
> > once-more
> Exiting...
`

	// Run
	main.Stdin = strings.NewReader(input)
	output, err := main.Output()
	if err != nil {
		fmt.Printf("Received error running main: %s\n", err.Error())
		t.Fail()
	}

	strOutput := string(output)
	if strOutput != wanted {
		fmt.Printf("Failed to match wanted output exactly. Got:\n%s\n", strOutput)
		t.Fail()
	}
}

func TestMain_Transactions(t *testing.T) {
	main := exec.Command("./kv-store")

	input := `WRITE a hello
	READ a
	START
	WRITE a hello-again
	READ a
	START
	DELETE a
	READ a
	COMMIT
	READ a
	WRITE a once-more
	READ a
	ABORT
	READ a
	QUIT
`

	wanted := `> > hello
> > > hello-again
> > > Key not found: a
> > Key not found: a
> > once-more
> > hello
> Exiting...
`

	// Run
	main.Stdin = strings.NewReader(input)
	output, err := main.Output()
	if err != nil {
		fmt.Printf("Received error running main: %s\n", err.Error())
		t.Fail()
	}

	strOutput := string(output)
	if strOutput != wanted {
		fmt.Printf("Failed to match wanted output exactly. Got:\n%s\n", strOutput)
		t.Fail()
	}
}
