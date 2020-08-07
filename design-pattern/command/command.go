package command


import (
	"fmt"
)

//实现receiver
type TV struct{}

func (t *TV)open() {
	fmt.Println("开机")
}

func (tv TV) Close() {
	fmt.Println("关机")
}

func (tv TV) ChangeChannel() {
	fmt.Println("换台")
}


//实现Command和ConcreteCommand
type Command interface {
	Execute()
}

type OpenCommand struct {
	receiver TV
}

func (oc OpenCommand) Execute() {
	oc.receiver.Open()
}

type CloseCommand struct {
	receiver TV
}

func (cc CloseCommand) Execute() {
	cc.receiver.Close()
}

type ChangeChannelCommand struct {
	receiver TV
}

func (ccc ChangeChannelCommand) Execute() {
	ccc.receiver.ChangeChannel()
}


//实现Invoke
type Invoke struct {
	Command
}

//定义工厂方法创建Command
func NewCommand(t string, tv TV) Command {
	switch t {
	case "open":
		return &OpenCommand{
			receiver: tv,
		}
	case "close":
		return CloseCommand{
			receiver: tv,
		}
	case "changechannel":
		return ChangeChannelCommand{
			receiver: tv,
		}
	default:
		return nil
	}
}


