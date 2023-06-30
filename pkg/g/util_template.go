package g

import (
	"bytes"
	"text/template"
)

type Template struct {
	Info string                 // 模板文件
	Name string                 // 模板名称
	Data map[string]interface{} // 模板数据
}

func NewTemplate(name string, info string, data map[string]interface{}) *Template {
	return &Template{
		Info: info,
		Name: name,
		Data: data,
	}
}

func (t *Template) ToString() (str string, err error) {
	te := template.Must(template.New(t.Name).Parse(t.Info))
	buffer := new(bytes.Buffer)
	err = te.Execute(buffer, t.Data)
	if err != nil {
		return
	}
	str = buffer.String()
	return
}

func (t *Template) ToFormattedString() (str string, err error) {
	str, err = t.ToString()
	if err != nil {
		return
	}
	str = Format(str)
	return

}
