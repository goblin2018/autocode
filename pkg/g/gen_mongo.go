package g

import (
	"auto/pkg/log"
	_ "embed"
	"fmt"
	"path"
	"strings"
)

//go:embed tpl/mongo/base.tpl
var mongoBase string

//go:embed tpl/mongo/model.tpl
var mongoModel string

//go:embed tpl/mongo/list.tpl
var mongoListFunc string

func GenMongo(baseDir string, name string, model *M) (err error) {
	genBase(baseDir, name, model)
	r, _ := genModel(baseDir, name, model)
	genFuncs(baseDir, name, r["StructName"].(string), model)
	err = updateServiceContext(baseDir, name)
	if err != nil {
		log.Errorf("update service context failed: %v", err)
	}
	return
}

// genBase
func genBase(baseDir string, name string, model *M) {
	file := NewFile(path.Join(baseDir, modelDir, name, "base.go"))
	builder := NewBuilder()
	// 第一个作为 base model
	for i, s := range model.Structs {
		st := NewStruct(s)
		builder.AddMongoStruct(st, i == 0)
	}

	t := NewTemplate("mongoBase", mongoBase, map[string]interface{}{
		"types":   builder.String(),
		"package": name,
	})
	err := file.Write(t)
	if err != nil {
		fmt.Println(err)
	}
}

func genModel(baseDir, name string, model *M) (r map[string]interface{}, err error) {
	file := NewFile(path.Join(baseDir, modelDir, name, "model.go"))

	r, err = parseMongoModel(name, model)
	if err != nil {
		return
	}

	t := NewTemplate("mongoModel", mongoModel, r)
	err = file.Write(t)
	return
}

func genFuncs(baseDir, name string, structName string, model *M) (err error) {
	for _, mf := range model.MFs {
		file := NewFile(path.Join(baseDir, modelDir, name, toLowerAndUnderscore(mf.Name)+".go"))
		var t *Template
		if mf.Type == "list" {
			r, err := parseMongoListFunc(name, structName, mf)
			if err != nil {
				return err
			}
			t = NewTemplate(mf.Name, mongoListFunc, r)
		}

		err = file.Write(t)
	}

	return
}

func parseMongoListFunc(name string, structName string, mf *MF) (r map[string]interface{}, err error) {
	r = make(map[string]interface{})
	r["package"] = name
	r["funcName"] = mf.Name
	r["StructName"] = structName

	req := NewStruct(mf.Input)
	r["reqName"] = req.Name

	builder := strings.Builder{}
	for _, f := range req.Fields {
		if f.Name == "Page" {
			r["hasPage"] = true
		}

		// 忽略分页参数
		if f.Name == "Page" || f.Name == "Size" || f.Name == "All" {
			continue
		}

		zero := getZero(f.Type)

		fmt.Fprintf(&builder, "\tif req.%s != %s {\n", f.Name, zero)
		fmt.Fprintf(&builder, "\t\tfilter[\"%s\"] = req.%s\n", f.GetTagWithoutOmit(), f.Name)
		fmt.Fprintf(&builder, "\t}\n")
	}

	r["filter"] = builder.String()

	return
}

func getZero(t string) string {
	switch t {
	case "string":
		return `""`
	case "int64":
		return "0"
	default:
		return "unknown"
	}
}

func parseMongoModel(name string, model *M) (r map[string]interface{}, err error) {
	r = make(map[string]interface{})
	r["package"] = name
	base := NewStruct(model.Structs[0])
	r["StructName"] = base.Name

	r["UniKeysWithColon"] = ""
	r["UniKeysWithComma"] = ""
	r["UniKeysWithDataComma"] = ""
	r["UniKeysWithAnd"] = ""
	r["UniKeysWithType"] = ""
	r["UniKeysWithoutType"] = ""
	r["UniKeysBsonFilter"] = ""

	// Todo 目前只有一组唯一Key
	uniKeys := base.UniKeys[0]
	log.Infof("uniKeys: %+v", uniKeys)
	for i, k := range uniKeys.Columns {
		first := i == 0
		r["UniKeysWithColon"] = fmt.Sprintf("%s:%s", r["UniKeysWithColon"], k.Json)
		r["UniKeysWithComma"] = fmt.Sprintf(`%s, "%s"`, r["UniKeysWithComma"], k.Json)
		if first {
			r["UniKeysWithDataComma"] = fmt.Sprintf(`data.%s`, k.Name)
		} else {
			r["UniKeysWithDataComma"] = fmt.Sprintf(`%s, data.%s`, r["UniKeysWithDataComma"], k.Name)
		}
		if first {
			r["UniKeysWithAnd"] = k.Name
		} else {
			r["UniKeysWithAnd"] = fmt.Sprintf(`%sAnd%s`, r["UniKeysWithAnd"], k.Name)
		}

		if first {
			r["UniKeysWithType"] = fmt.Sprintf(`%s %s`, k.Json, k.Type)
			r["UniKeysWithoutType"] = k.Json
		} else {
			r["UniKeysWithType"] = fmt.Sprintf(`%s, %s %s`, r["UniKeysWithType"], k.Json, k.Type)
			r["UniKeysWithoutType"] = fmt.Sprintf(`%s, %s`, r["UniKeysWithoutType"], k.Json)
		}

		if first {
			r["UniKeysBsonFilter"] = fmt.Sprintf(`"%s": %s`, k.Json, k.Json)
		} else {
			r["UniKeysBsonFilter"] = fmt.Sprintf(`%s, "%s": %s`, r["UniKeysBsonFilter"], k.Json, k.Json)
		}
	}

	return
}

// 只有第一个需要使用 _id updateAt createAt
func (b *Builder) AddMongoStruct(st *Struct, isBase bool) {
	fmt.Fprintf(b, "\n")
	fmt.Fprintf(b, "type %s struct {\n", st.Name)
	if isBase {
		fmt.Fprintf(b, "\tId string `json:\"id\" bson:\"_id\"`\n")
		fmt.Fprintf(b, "\tUpdateAt int64 `json:\"updateAt\" bson:\"updateAt\"`\n")
		fmt.Fprintf(b, "\tCreateAt int64 `json:\"createAt\" bson:\"createAt\"`\n")
	}

	for _, f := range st.Fields {
		if isBase {
			// 过滤掉 _id updateAt createAt
			if f.Name == "Id" || f.Name == "UpdateAt" || f.Name == "CreateAt" {
				continue
			}
		}

		// 只有API文件需要
		if f.ApiOnly {
			continue
		}

		if isBase {
			fmt.Fprintf(b, "\t%s\n", f.ToMongo())
		} else {
			fmt.Fprintf(b, "\t%s\n", f.ToJsonWithoutOmit())
		}
	}
	fmt.Fprint(b, "}\n")
}
