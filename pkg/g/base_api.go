package g

// api struct
type A struct {
	Name   string
	Input  interface{}
	Output interface{}
}

func Apis(apis ...*A) []*A {
	return apis
}

func Api(name string, input interface{}, output interface{}) *A {
	return &A{
		Name:   name,
		Input:  input,
		Output: output,
	}
}
