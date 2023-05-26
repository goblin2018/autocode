package g

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

func Json(name string) string {
	return "json:" + name
}

func OmitJson(name string) string {
	return "json:" + name + ",omitempty"
}

func Uni(name string) string {
	return "uni:" + name
}

func Key(name string) string {
	return "key:" + name
}

func Fields(fields ...F) []F {
	return fields
}

type F struct {
	Name     string
	Type     string
	Json     string
	UniKey   string
	Key      string
	Desc     string
	Optional bool
}

func Field(name string, t string, tags ...string) F {
	return F{
		Name: name,
		Type: t,
	}
}
