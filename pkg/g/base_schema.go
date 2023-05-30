package g

type Sc struct {
	Structs []S
	Groups  []*G
	Model   *M
	Name    string
}

func Schema(name string, structs []S, groups []*G, model *M) Sc {
	return Sc{
		Name:    name,
		Structs: structs,
		Groups:  groups,
		Model:   model,
	}
}
