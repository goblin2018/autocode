package g

import (
	_ "embed"
	"fmt"
	"path"
)

//go:embed tpl/type/type.tpl
var typeT string

func GenTypes(baseDir string, name string, structs []S) {
	file := NewFile(path.Join(baseDir, typeDir, name+".go"))
	builder := NewBuilder()
	for _, s := range structs {
		builder.AddTypeStruct(s)
	}

	t := NewTemplate("typeT", typeT, map[string]interface{}{
		"types": builder.String(),
	})

	err := file.Write(t)
	if err != nil {
		fmt.Println(err)
	}
}
