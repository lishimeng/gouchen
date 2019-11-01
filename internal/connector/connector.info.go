package connector

type Info struct {
	Id    string
	Name  string
	Type  string
	State bool
}

func (info Info) GetID() string {
	return info.Id
}

func (info Info) GetName() string {
	return info.Name
}

func (info Info) GetType() string {
	return info.Type
}

func (info Info) GetState() bool {
	return info.State
}
