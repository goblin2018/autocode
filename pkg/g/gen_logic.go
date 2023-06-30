package g

import (
	"auto/pkg/log"
	"bytes"
	_ "embed"
	"fmt"
	"path"
	"strings"
	"text/template"
	"unicode"
)

//go:embed tpl/logic/logic.tpl
var logicT string

func GenLogics(baseDir string, pkgName string, sc Sc) error {
	baseStruct := NewStruct(sc.Structs[0])
	for _, group := range sc.Groups {
		for _, api := range group.Apis {
			err := genLogic(baseDir, pkgName, group, api, baseStruct)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func genLogic(baseDir string, pkgName string, group *G, api *A, baseStruct *Struct) error {
	file := NewFile(path.Join(baseDir, logicDir, group.Name, api.Name+".go"))
	logic := FirstCharToUpper(api.Name) + "Logic"
	function := FirstCharToUpper(api.Name)

	req := NewStruct(api.Input)
	request := ""
	if len(req.Fields) > 0 {
		request = fmt.Sprintf("req *api.%s", req.Name)
	}
	response := ""
	resp := NewStruct(api.Output)
	if len(resp.Fields) > 0 {
		response = fmt.Sprintf("resp *api.%s, ", resp.Name)
	}

	t := NewTemplate(logic, logicT, map[string]interface{}{
		"package":  pkgName,
		"logic":    logic,
		"function": function,
		"request":  request,
		"response": response,
		"body":     genBody(api, pkgName, baseStruct),
		"imports":  genLogicImports(api, pkgName),
	})

	return file.Write(t)

}

const addTemplate = `
md := new({{.package}}.{{.name}})
copier.Copy(md, req)
err = l.{{.modelName}}.Create(md)
`

const updateTemplate = `
md := new({{.package}}.{{.name}})
copier.Copy(md, req)
err = l.{{.modelName}}.Update(md)
`

const delTemplate = `
err = l.{{.modelName}}.Delete(req.Id)
`

const listTemplate = `
resp = new(api.{{.respName}})
	opt := new({{.package}}.{{.reqName}})
	copier.Copy(opt, req)
	items, _ := l.{{.modelName}}.List(l.ctx, opt)
	for _, it := range items {
		td := new(api.{{.name}})
		copier.Copy(td, it)
		resp.Items = append(resp.Items, td)
	}

	{{if .usePage}}
	if opt.Size != 0 {
		resp.Total, _ = l.{{.modelName}}.Count(l.ctx, opt)
	}
	{{end}}
`

func genBody(api *A, pkgName string, baseStruct *Struct) string {
	body := ""

	req := NewStruct(api.Input)
	resp := NewStruct(api.Output)
	modelName := fmt.Sprintf("%sModel", toUpperCamel(pkgName))
	opt := map[string]interface{}{
		"package":   pkgName,
		"name":      baseStruct.Name,
		"modelName": modelName,
	}

	// 增加增删改查的基础逻辑
	switch api.Name {
	case "add":
		body = genTmpl("add", addTemplate, opt)
	case "update":
		body = genTmpl("update", updateTemplate, opt)
	case "del":
		body = genTmpl("del", delTemplate, opt)
	default:
		opt["reqName"] = req.Name
		opt["respName"] = resp.Name
		if strings.HasPrefix(api.Name, "list") {
			for _, f := range req.Fields {
				if f.Name == "Page" {
					opt["usePage"] = true
					continue
				}
			}

			body = genTmpl(api.Name, listTemplate, opt)
		}
	}

	return body
}

// 通用的模板生成逻辑
func genTmpl(name string, templateStr string, opt map[string]interface{}) string {

	log.Infof("genTmpl: %s %+v", name, opt)
	tmpl, err := template.New(name).Parse(templateStr)
	if err != nil {
		log.Errorf("genTmpl: %s", err)
		return ""
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, opt)

	return buf.String()
}

func FirstCharToUpper(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s) // 将字符串转为rune slice
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// genLogicImport

func genLogicImports(api *A, pkgName string) string {

	if api.Name != "del" {
		return fmt.Sprintf(`
		"auto/app/model/%s"
		"github.com/jinzhu/copier"
		`, pkgName)
	}
	return ""
}
