package command

type CommandManager interface {
	Start(filePath string) error
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
