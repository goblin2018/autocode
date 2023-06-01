package g

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"

	"golang.org/x/tools/go/ast/astutil"
)

const routerFile = "handler/router.go"

func updateRouter(baseDir string, groupName string) (err error) {
	path := path.Join(baseDir, routerFile)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return
	}
	astutil.AddImport(fset, node, fmt.Sprintf("auto/app/handler/%s", groupName))
	addCaller(node, groupName)
	// 格式化抽象语法树；将 node 内容写入 buf
	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		return
	}

	// 将格式化后的内容写入磁盘文件
	formattedSrc := buf.String()
	err = ioutil.WriteFile(path, []byte(formattedSrc), 0644)
	if err != nil {
		return
	}

	return

}

func addCaller(node *ast.File, groupName string) {
	var registerRoutersFunc *ast.FuncDecl
	for _, decl := range node.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok && funcDecl.Name.Name == "RegisterRouters" {
			registerRoutersFunc = funcDecl
			break
		}
	}

	exists := false
	ast.Inspect(registerRoutersFunc, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		if selectorExpr.X.(*ast.Ident).Name == groupName && selectorExpr.Sel.Name == "RegisterTo" {
			exists = true
			return false
		}
		return true
	})

	if !exists {
		callExpr := &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent(groupName),
				Sel: ast.NewIdent("RegisterTo"),
			},
			Args: []ast.Expr{
				ast.NewIdent("g"),
				ast.NewIdent("svc"),
			},
		}
		registerRoutersFunc.Body.List = append(registerRoutersFunc.Body.List, &ast.ExprStmt{X: callExpr})
	}

}
