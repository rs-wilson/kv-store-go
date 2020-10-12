package commands

import (
	"fmt"
)

type WriteCommand struct{}

func (me *WriteCommand) Name() string {
	return "WRITE"
}

func (me *WriteCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("%s command requires exactly two arguments: KEY and VALUE", me.Name())
	}
	return &WriteCommandInstance{
		key:   args[0],
		value: args[1],
	}, nil
}

type WriteCommandInstance struct {
	key   string
	value string
}

func (me *WriteCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	context.KVStore[me.key] = me.value
	return context, "", nil
}
