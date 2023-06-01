package g

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

func Api(name string, path string, input interface{}, output interface{}) *A {
	return &A{
		Name:   name,
		Path:   path,
		Input:  input,
		Output: output,
	}
}
