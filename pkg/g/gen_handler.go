package g

import (
	"auto/pkg/log"
	_ "embed"
	"path"
)

//go:embed tpl/handler/handler.tpl
var handlerT string

func GenHandlers(baseDir string, pkgName string, sc Sc) error {
	baseStruct := NewStruct(sc.Structs[0])
	for _, group := range sc.Groups {
		genGroup(baseDir, pkgName, group)
		updateRouter(baseDir, group.Name)
		for _, api := range group.Apis {
			err := genHandler(baseDir, pkgName, group, api, baseStruct)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}

func genHandler(baseDir string, pkgName string, group *G, api *A, baseStruct *Struct) error {
	file := NewFile(path.Join(baseDir, handlerDir, group.Name, api.Name+".go"))
	logicFuncName := FirstCharToUpper(api.Name)

	req := NewStruct(api.Input)
	resp := NewStruct(api.Output)

	t := NewTemplate(api.Name, handlerT, map[string]interface{}{
		"package":       pkgName,
		"group":         group.Name,
		"reqName":       req.Name,
		"hasResp":       len(resp.Fields) > 0,
		"logicFuncName": logicFuncName,
		"api":           api.Name,
	})

	return file.Write(t)
}
