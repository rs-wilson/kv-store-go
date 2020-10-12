package repl

import (
	"fmt"
	"kvstore/commands"
	"testing"
	"time"
)

const TestExitCommand = "STOP"
const TestMockCommand = "MOCK"
const TestMockMessage = "mocked"

type mockCommand struct{}

func (me *mockCommand) Name() string {
	return TestMockCommand
}

func (me *mockCommand) NewInstance(...string) (commands.CommandInstance, error) {
	return &mockCommandInstance{}, nil
}

type mockCommandInstance struct{}

func (me *mockCommandInstance) Evaluate(context *commands.CommandContext) (*commands.CommandContext, string, error) {
	return context, TestMockMessage, nil
}

func GetTestEvaluator() *Evaluator {
	return NewEvaluator(TestExitCommand, &mockCommand{})
}

func SetupEvalTest(timeout time.Duration) (chan<- string, <-chan string, <-chan struct{}, <-chan struct{}) {
	// make
	eval := GetTestEvaluator()
	in := make(chan string)
	out := make(chan string)

	// run
	done := make(chan struct{})
	go func() {
		defer close(done)
		eval.Run(in, out)
	}()

	timer := make(chan struct{})
	go func() {
		defer close(timer)
		time.Sleep(timeout)
	}()

	return in, out, done, timer
}

// tests that we can quit in a reasonable amount of time
func TestEvaluator_Stop(t *testing.T) {
	in, out, done, timeout := SetupEvalTest(2 * time.Second)

	go func() {
		result := <-out
		if result != ExitMessage {
			fmt.Printf("Got an unexpected message from evaluator: %s\n", result)
			t.Fail()
		}
	}()

	in <- TestExitCommand

	select {
	case <-done:
		// Success!
	case <-timeout:
		fmt.Printf("Timed out waiting for evaluator to stop.\n")
		t.Fail()
	}
}

// test that we can run a mock command and get expected output
func TestEvaluator_Run(t *testing.T) {
	in, out, _, timeout := SetupEvalTest(2 * time.Second)

	// Test
	done := make(chan struct{})
	go func() {
		defer close(done)
		result := <-out
		if result != TestMockMessage {
			fmt.Printf("Got an unexpected message from evaluator: %s\n", result)
			t.Fail()
		}
	}()

	// Test command
	in <- TestMockCommand

	select {
	case <-done:
		// Success!
	case <-timeout:
		fmt.Printf("Timed out waiting for evaluator to evaluate the mock command.\n")
		t.Fail()
	}
}

// TODO test that we can input a bad command

// TODO test that we can input a good command with bad args
