package view

type IView interface {
	GetName() string
	GetHelp() string
	Handle(args []string) error
}
