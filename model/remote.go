package model

type RemoteCommand struct {
	ButtonName    string
	RawSignalData string
}

type RemoteControl struct {
	Name    string
	Buttons []RemoteCommand
}

func (remote *RemoteControl) Add(name string, command RemoteCommand) {
	remote.Buttons = append(remote.Buttons, command)
}

//func (remote *RemoteControl) Remove(name string) {
//	delete(remote.Buttons, name)
//}
