package g

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
