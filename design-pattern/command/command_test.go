package command

import "testing"

func TestNewCommand(t *testing.T) {
	//创建一个Receiver
	tTV := TV{}

	//创建一个ConcreteCommand
	tCommand := NewCommand("open", tTV)
	//创建一个调用者
	tInvoke := Invoke{tCommand}
	tInvoke.Execute()

	tCommand = NewCommand("close", tTV)
	tInvoke = Invoke{tCommand}
	tInvoke.Execute()

	tCommand = NewCommand("changechannel", tTV)
	tInvoke = Invoke{tCommand}
	tInvoke.Execute()
}










