package main

import (
	"kvstore/commands"
	"kvstore/repl"
	"os"
)

func main() {
	// Configure Evaluator
	commands := GetConfiguredCommands()
	exitCommand := "QUIT"
	evaluator := repl.NewEvaluator(exitCommand, commands...)

	// Start command repl (blocks)
	in := repl.ReadLinesFrom(os.Stdin)
	out, done := repl.WriteLinesTo(os.Stdout)
	evaluator.Run(in, out)

	<-done // ensure writer finishes
}

func GetConfiguredCommands() []commands.Command {
	return []commands.Command{
		&commands.WriteCommand{},
		&commands.ReadCommand{},
		&commands.DeleteCommand{},
		&commands.StartCommand{},
		&commands.CommitCommand{},
		&commands.AbortCommand{},
	}
}
