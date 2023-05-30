package g

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

func MakeDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func MakeFile(file string) (*os.File, error) {
	return os.Create(file)
}

// 封装文件定义

type File struct {
	Path string // 文件全名
	Dir  string // 全部目录
	Pkg  string // 文件包名,即最后一级目录
	Name string // 文件名
}

func NewFile(path string) *File {
	dir := filepath.Dir(path)
	pkg := filepath.Base(dir)
	os.MkdirAll(dir, 0755)

	return &File{
		Path: path,
		Dir:  dir,
		Pkg:  pkg,
	}
}

// 写文件 追加
func (f *File) AppendString(content string) {
	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}

// 写文件 覆盖
func (f *File) WriteString(content string) (err error) {
	file, err := os.Create(f.Path)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return
}

// 默认使用 template 写
func (f *File) Write(t *Template) (err error) {
	te := template.Must(template.New(t.Name).Parse(t.Info))
	buffer := new(bytes.Buffer)
	err = te.Execute(buffer, t.Data)
	if err != nil {
		return
	}
	r := Format(buffer.String())
	err = f.WriteString(r)
	return
}

// 按照go语法格式化代码
func Format(code string) string {
	ret, err := format.Source([]byte(code))
	if err != nil {
		return code
	}
	return string(ret)
}
