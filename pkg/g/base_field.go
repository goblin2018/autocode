package g

import "fmt"

const (
	Int      = "int64"
	Str      = "string"
	Bool     = "bool"
	Arr      = "[]"
	Inerface = "interface{}"

	OmitEmpty = "omitempty"
)

var BaseFields = []F{
	Field("Id", "string"),
}

func WithBaseFields(fields ...F) []F {
	return Fields(append(fields, BaseFields...)...)
}

func Fields(fields ...F) []F {
	return fields
}

type F struct {
	Name     string
	Type     string
	Json     string
	UniKey   string
	Form     string
	Key      string
	Desc     string
	ApiOnly  bool
	Optional bool
}

func Field(name string, t string, tags ...string) F {
	return F{
		Name: name,
		Type: t,
	}
}

func (f F) ToMongo() string {
	mongoTag := fmt.Sprintf(`bson:"%s" json:"%s"`, f.Json, f.Json)
	return fmt.Sprintf("%s %s `%s`", f.Name, f.Type, mongoTag)
}
