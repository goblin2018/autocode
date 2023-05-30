package g

import (
	"fmt"
	"reflect"
	"strings"
)

type Builder struct {
	strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{
		Builder: strings.Builder{},
	}
}

func (b *Builder) AddTypeStruct(s S) {
	structType := reflect.TypeOf(s)
	fmt.Fprintf(b, "\n")
	fmt.Fprintf(b, "type %s struct {\n", structType.Name())

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Only retain JSON or Form tags.
		tags := getApiTags(field.Tag)
		t := field.Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		tname := parseType(t)
		fmt.Fprintf(b, "\t%s %s %s\n", field.Name, tname, tags)
	}

	fmt.Fprint(b, "}\n")
}

func parseType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + parseType(t.Elem())
	case reflect.Slice:
		return "[]" + parseType(t.Elem())
	case reflect.Struct:
		strs := strings.Split(t.Name(), ".")
		return strs[len(strs)-1]
	default:
		return t.String()
	}
}
