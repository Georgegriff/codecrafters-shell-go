package command

type UserDefined struct {
	Name string
	Path string
}

func (u UserDefined) CommandType() CommandType {
	return USER_DEFINED
}

func (u UserDefined) Run(args []string) {

}

func (u UserDefined) String() string {
	return u.Name
}
