package g

import (
	_ "embed"
	"fmt"
	"path"
)

//go:embed tpl/handler/g.tpl
var groupT string

func genGroup(baseDir string, pkgName string, group *G) {
	file := NewFile(path.Join(baseDir, handlerDir, group.Name, "a.go"))

	t := NewTemplate(group.Name, groupT, map[string]interface{}{
		"group": group.Name,
		"apis":  genApiRouterInfos(group),
	})

	file.Write(t)

}

func genApiRouterInfos(group *G) string {
	info := ""
	for _, api := range group.Apis {
		info = fmt.Sprintf("%s\n\t\t%s", info, genApiRouterInfoItem(api))
	}
	return info
}

func genApiRouterInfoItem(api *A) string {
	methodName := "GET"
	switch api.Name {
	case "add":
		methodName = "POST"
	case "list":
		methodName = "GET"
	case "update":
		methodName = "PUT"
	case "del":
		methodName = "DELETE"
	}

	return fmt.Sprintf("t.%s(\"%s\",%s)", methodName, api.Path, api.Name)
}
