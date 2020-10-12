package commands

import (
	"fmt"
)

type ReadCommand struct{}

func (me *ReadCommand) Name() string {
	return "READ"
}

func (me *ReadCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("%s command requires exactly one arguments: KEY", me.Name())
	}
	return &ReadCommandInstance{
		key: args[0],
	}, nil
}

type ReadCommandInstance struct {
	key string
}

func (me *ReadCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	value, found := context.KVStore[me.key]
	if !found {
		return context, fmt.Sprintf("Key not found: %s", me.key), nil
	}
	return context, value, nil
}
