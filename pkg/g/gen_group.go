package g

import (
	_ "embed"
	"fmt"
	"log"
	"path"
)

//go:embed tpl/handler/g.tpl
var groupT string

const groupTemplate = `
	
t{{.index}} := en.Group("{{.group}}")
	{
    {{.apis}}
	}
`

func genGroup(baseDir string, pkgName string, groups []*G) {
	file := NewFile(path.Join(baseDir, handlerDir, pkgName, "a.go"))

	info := ""
	for index, group := range groups {
		tg := NewTemplate(pkgName, groupTemplate, map[string]interface{}{
			"index": index + 1,
			"group": pkgName,
			"apis":  genApiRouterInfos(group, index+1),
		})

		str, err := tg.ToFormattedString()
		if err != nil {
			log.Fatal(err)
		}

		info = fmt.Sprintf("%s\n%s", info, str)

	}

	t := NewTemplate(pkgName, groupT, map[string]interface{}{
		"pkgName": pkgName,
		"info":    info,
	})
	file.Write(t)

}

func genApiRouterInfos(group *G, index int) string {
	info := ""
	for _, api := range group.Apis {
		info = fmt.Sprintf("%s\n\t\t%s", info, genApiRouterInfoItem(api, index))
	}
	return info
}

func genApiRouterInfoItem(api *A, index int) string {
	return fmt.Sprintf("t%d.%s(\"%s\",%s)", index, api.Method, api.Path, api.Name)
}
