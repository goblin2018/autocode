package g

const (
	GET  = "GET"
	POST = "POST"
	PUT  = "PUT"
	DEL  = "DELETE"
)

// api struct
type A struct {
	Name   string
	Path   string
	Input  interface{}
	Output interface{}
}

func Apis(apis ...*A) []*A {
	return apis
}

// Api 用于生成 api 的定义
// @param name api 名称
// path api 路径
// method POST PUT GET DELETE
// input api 输入参数
// output api 输出参数
func Api(name string, path string, method string, input interface{}, output interface{}) *A {
	return &A{
		Name:   name,
		Path:   path,
		Input:  input,
		Output: output,
	}
}
