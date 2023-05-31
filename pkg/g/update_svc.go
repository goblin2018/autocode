package g

import (
	"auto/pkg/log"
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
	"strings"
	"unicode"

	"golang.org/x/tools/go/ast/astutil"
)

const serviceContextFile = "svc/context.go"

func updateServiceContext(baseDir string, pkgName string) error {
	// 解析 service.go 文件
	fset := token.NewFileSet()
	path := path.Join(baseDir, serviceContextFile)

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	modelName := fmt.Sprintf("%sModel", toUpperCamel(pkgName))
	astutil.AddImport(fset, node, fmt.Sprintf("auto/app/model/%s", pkgName))

	// 查找并更新 ServiceContext 结构体和 NewServiceContext 函数
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == "ServiceContext" {
				st, ok := x.Type.(*ast.StructType)
				if !ok {
					return false
				}
				log.Info("add model to ServiceContext")

				// 添加新的模型到 ServiceContext 结构体
				addModelToStruct(st, pkgName, modelName)
			}
		case *ast.FuncDecl:
			if x.Name.Name == "NewServiceContext" {
				// 添加新的模型到 NewServiceContext 函数
				addModelToFunc(x, pkgName, modelName)
			}
		}
		return true
	})

	// 格式化抽象语法树；将 node 内容写入 buf
	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		return err
	}

	// 将格式化后的内容写入磁盘文件
	formattedSrc := buf.String()
	err = ioutil.WriteFile(path, []byte(formattedSrc), 0644)
	if err != nil {
		return err
	}

	return nil
}

// 向 ServiceContext 结构体添加新的模型
func addModelToStruct(st *ast.StructType, pkgName string, modelName string) {
	modelType := fmt.Sprintf("*%s.Model", pkgName)

	for _, field := range st.Fields.List {
		if field.Names[0].Name == modelName {
			// 模型已存在，不需要添加
			return
		}
	}

	// 添加新的模型
	st.Fields.List = append(st.Fields.List, &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(modelName)},
		Type:  ast.NewIdent(modelType),
	})
}

// 向 NewServiceContext 函数添加新的模型初始化
func addModelToFunc(fn *ast.FuncDecl, pkgName, modelName string) {
	// 查找 return 语句
	for _, stmt := range fn.Body.List {
		retStmt, ok := stmt.(*ast.ReturnStmt)
		if !ok {
			continue
		}

		// 查找 return &ServiceContext{...} 语句
		for _, expr := range retStmt.Results {
			ce, ok := expr.(*ast.UnaryExpr)
			if !ok {
				continue
			}

			if ident, ok := ce.X.(*ast.CompositeLit); ok {
				if structIdent, ok := ident.Type.(*ast.Ident); ok && structIdent.Name == "ServiceContext" {
					// 检查模型是否已经存在
					for _, elt := range ident.Elts {
						kve := elt.(*ast.KeyValueExpr)
						if kve.Key.(*ast.Ident).Name == modelName {
							// 模型已存在，不需要添加
							return
						}
					}

					// 添加新的模型初始化
					ident.Elts = append(ident.Elts, &ast.KeyValueExpr{
						Key:   ast.NewIdent("\n\t\t" + modelName),
						Value: ast.NewIdent(fmt.Sprintf("%s.NewModel(mongoModel),\n", pkgName)),
					})
				}
			}
		}
	}
}

func toUpperCamel(s string) string {
	var buf strings.Builder
	underScore := true

	for _, r := range s {
		if r == '_' {
			underScore = true
		} else {
			if underScore {
				buf.WriteRune(unicode.ToUpper(r))
				underScore = false
			} else {
				buf.WriteRune(r)
			}
		}
	}
	return buf.String()
}
