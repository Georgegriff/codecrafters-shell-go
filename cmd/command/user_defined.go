package command

type UserDefined struct {
	BaseCommand
	Name string
	Path string
}

func (u UserDefined) Run(args []string) {

}

func (u UserDefined) String() string {
	return u.Name
}
