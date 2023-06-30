package g

type Sc struct {
	Structs []S
	Groups  []*G
	Model   *M
	Name    string
}

// name 数据表的名称
func Schema(name string, structs []S, groups []*G, model *M) Sc {
	return Sc{
		Name:    name,
		Structs: structs,
		Groups:  groups,
		Model:   model,
	}
}
