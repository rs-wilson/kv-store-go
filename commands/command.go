package commands

type Command interface {
	Name() string
	NewInstance(...string) (CommandInstance, error)
}

type CommandInstance interface {
	Evaluate(*CommandContext) (*CommandContext, string, error)
}

type CommandContext struct {
	Parent  *CommandContext
	KVStore map[string]string
}

// TODO: Write unit tests for all commands
