package g

import (
	"auto/pkg/log"
	_ "embed"
	"path"
)

//go:embed tpl/handler/handler.tpl
var handlerT string

func GenHandlers(baseDir string, sc Sc) error {
	baseStruct := NewStruct(sc.Structs[0])
	genGroup(baseDir, sc.Name, sc.Groups)
	for _, group := range sc.Groups {
		updateRouter(baseDir, sc.Name)
		for _, api := range group.Apis {
			err := genHandler(baseDir, sc.Name, group, api, baseStruct)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}

func genHandler(baseDir string, pkgName string, group *G, api *A, baseStruct *Struct) error {
	file := NewFile(path.Join(baseDir, handlerDir, pkgName, api.Name+".go"))
	logicFuncName := FirstCharToUpper(api.Name)

	req := NewStruct(api.Input)
	resp := NewStruct(api.Output)

	t := NewTemplate(api.Name, handlerT, map[string]interface{}{
		"package":       pkgName,
		"group":         pkgName,
		"reqName":       req.Name,
		"hasResp":       len(resp.Fields) > 0,
		"logicFuncName": logicFuncName,
		"api":           api.Name,
	})

	return file.Write(t)
}
