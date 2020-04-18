package command

import (
	"context"
)

type Command interface {
	Execute(ctx context.Context)
}

type Invoker struct {
	cmdList []Command
}

func (invoke *Invoker) AddCmd(cmd Command) {
	if invoke.cmdList == nil {
		invoke.cmdList = make([]Command, 0, 2)
	}

	invoke.cmdList = append(invoke.cmdList, cmd)
}

func (invoke *Invoker) SetRunEnv() {

}

func (invoke *Invoker) ExecuteCommand(ctx context.Context) {
	invoke.SetRunEnv()
	for _, cmd := range invoke.cmdList {
		cmd.Execute(ctx)
	}
}
