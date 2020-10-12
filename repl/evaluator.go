package repl

import (
	"fmt"
	"kvstore/commands"
	"strings"
)

const ExitMessage = "Exiting..."
const NoCommandMessage = "No command entered."
const CommandNotFoundMessage = "Command not recognized."
const FailedToInstanceCommandMessage = "Failed to parse command: %s"
const FailedToExecuteCommandMessage = "Failed to execute command: %s"

type Evaluator struct {
	ExitCommand string
	Commands    map[string]commands.Command
	Context     *commands.CommandContext
}

func NewEvaluator(exit string, commandList ...commands.Command) *Evaluator {
	commandMap := make(map[string]commands.Command)
	for _, command := range commandList {
		commandMap[command.Name()] = command
	}

	return &Evaluator{
		ExitCommand: exit,
		Commands:    commandMap,
		Context: &commands.CommandContext{
			Parent:  nil,
			KVStore: make(map[string]string),
		},
	}
}

func (me *Evaluator) Run(input <-chan string, output chan<- string) error {
	defer close(output)
	for line := range input {
		commandName, args := splitLine(line)

		// handle exit
		if commandName == me.ExitCommand {
			output <- ExitMessage
			break
		}

		// handle nothing
		if commandName == "" {
			output <- NoCommandMessage
			continue
		}

		// otherwise find & run command
		command, found := me.Commands[commandName]
		if !found {
			output <- CommandNotFoundMessage
			continue
		}

		// make an instance of the command to run it
		instance, err := command.NewInstance(args...)
		if err != nil {
			output <- fmt.Sprintf(FailedToInstanceCommandMessage, err.Error())
			continue
		}

		// evalutate
		newContext, outMessage, err := instance.Evaluate(me.Context)
		if err != nil {
			output <- fmt.Sprintf(FailedToExecuteCommandMessage, err.Error())
			continue
		}

		// ensure context is updated
		me.Context = newContext

		// send output of command along
		output <- outMessage
	}
	return nil
}

func splitLine(line string) (string, []string) {
	args := make([]string, 0)

	words := strings.Fields(line)
	if len(words) == 0 {
		return "", args
	}

	command := strings.ToUpper(words[0]) // Always upper-case commands
	args = words[1:]
	return command, args
}
