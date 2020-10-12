package commands

import (
	"fmt"
)

type DeleteCommand struct{}

func (me *DeleteCommand) Name() string {
	return "DELETE"
}

func (me *DeleteCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("%s command requires exactly one arguments: KEY", me.Name())
	}
	return &DeleteCommandInstance{
		key: args[0],
	}, nil
}

type DeleteCommandInstance struct {
	key string
}

func (me *DeleteCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	_, found := context.KVStore[me.key]
	if !found {
		return context, fmt.Sprintf("Key not found: %s", me.key), nil
	}
	delete(context.KVStore, me.key)
	return context, "", nil
}
