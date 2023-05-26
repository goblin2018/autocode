package g

type Sc struct {
	Structs []Struct
	Groups  []G
}

func Schema(structs []Struct, groups []G) Sc {
	return Sc{
		Structs: structs,
		Groups:  groups,
	}
}
