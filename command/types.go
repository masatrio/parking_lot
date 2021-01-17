package command

type CommandManager interface {
	Start(filePath string)
	Handle(cmd Command) bool
}

type Command interface {
	GetInput() []string
	GetOutput() string
	GetErr() error
	Error(err error)
	Output(str string)
	Print()
}
