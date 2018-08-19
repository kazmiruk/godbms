package command

type Command interface {
	Process() (response string)
}
