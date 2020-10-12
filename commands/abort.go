package commands

import "fmt"

type AbortCommand struct{}

func (me *AbortCommand) Name() string {
	return "ABORT"
}

func (me *AbortCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("%s command requires no arguments", me.Name())
	}
	return &AbortCommandInstance{}, nil
}

type AbortCommandInstance struct {
}

func (me *AbortCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	if context.Parent == nil {
		return context, "Not in a transaction.", nil
	}

	return context.Parent, "", nil
}
