package g

// group struct
type G struct {
	Name string
	Apis []*A
}

func Group(name string, apis ...*A) *G {
	return &G{
		Name: name,
		Apis: apis,
	}
}

func Groups(groups ...*G) []*G {
	return groups
}
