package g

// group struct
type G struct {
	MiddleWares []interface{}
	Apis        []*A
}

// name api group 名称
func Group(name string, apis ...*A) *G {
	return &G{
		Apis: apis,
	}
}

func Groups(groups ...*G) []*G {
	return groups
}
