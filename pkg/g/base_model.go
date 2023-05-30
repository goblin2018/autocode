package g

type M struct {
	DB      string // mongo | mysql | postgres
	Structs []S    // 对应的结构体
	MFs     []*MF  // 对应的方法
	UniKeys []*K   // 唯一索引
	Keys    []*K   // 普通索引
}

type K struct {
	Columns []F
	Name    string
}

func UniKeys(keys ...*K) []*K {
	return keys
}

func Keys(keys ...*K) []*K {
	return keys
}

func Models(models ...*M) []*M {
	return models
}

func Model(db string, structs []S, mfs []*MF) *M {
	return &M{
		DB:      db,
		Structs: structs,
		MFs:     mfs,
	}
}

type MF struct {
	Name   string
	Input  S
	Output S
}

func ModelFuncs(mfs ...*MF) []*MF {
	return mfs
}

func ModelFunc(name string, input S, output S) *MF {
	return &MF{
		Name:   name,
		Input:  input,
		Output: output,
	}
}
