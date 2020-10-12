package commands

import "fmt"

type CommitCommand struct{}

func (me *CommitCommand) Name() string {
	return "COMMIT"
}

func (me *CommitCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("%s command requires no arguments", me.Name())
	}
	return &CommitCommandInstance{}, nil
}

type CommitCommandInstance struct {
}

func (me *CommitCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	if context.Parent == nil {
		return context, "Not in a transaction.", nil
	}

	context.Parent.KVStore = context.KVStore
	return context.Parent, "", nil
}
