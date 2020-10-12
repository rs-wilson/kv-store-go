package commands

import "fmt"

type StartCommand struct{}

func (me *StartCommand) Name() string {
	return "START"
}

func (me *StartCommand) NewInstance(args ...string) (CommandInstance, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("%s command requires no arguments", me.Name())
	}
	return &StartCommandInstance{}, nil
}

type StartCommandInstance struct {
}

func (me *StartCommandInstance) Evaluate(context *CommandContext) (*CommandContext, string, error) {
	newStore := make(map[string]string)
	for k, v := range context.KVStore {
		newStore[k] = v
	}
	newContext := &CommandContext{
		Parent:  context,
		KVStore: newStore,
	}
	return newContext, "", nil
}
