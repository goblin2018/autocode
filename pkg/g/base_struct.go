package g

import (
	"reflect"
)

type empty struct {
}

var Empty = empty{}

type S interface{}

func Ss(structs ...S) []S {
	return structs
}

type Struct struct {
	Name    string
	Fields  []F
	Desc    string
	UniKeys []K
	Keys    []K
}

func NewStruct(s S) *Struct {
	structType := reflect.TypeOf(s)

	r := &Struct{
		Name: structType.Name(),
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag

		sf := F{
			Name:    field.Name,
			Type:    parseType(field.Type),
			Json:    tag.Get("json"),
			UniKey:  tag.Get("uni"),
			Key:     tag.Get("key"),
			ApiOnly: tag.Get("api") == "true",
			Form:    tag.Get("form"),
		}

		r.AddUniKey(sf)
		r.AddKey(sf)

		r.Fields = append(r.Fields, sf)
	}

	return r
}

func (s *Struct) AddUniKey(f F) {
	if f.UniKey == "" {
		return
	}
	for i, uni := range s.UniKeys {
		if uni.Name == f.UniKey {
			s.UniKeys[i].Columns = append(uni.Columns, f)
			return
		}
	}

	s.UniKeys = append(s.UniKeys, K{
		Name:    f.UniKey,
		Columns: []F{f},
	})

}

func (s *Struct) AddKey(f F) {

	if f.Key == "" {
		return
	}

	for _, k := range s.Keys {
		if k.Name == f.Key {
			k.Columns = append(k.Columns, f)
			return
		}
	}

	s.UniKeys = append(s.Keys, K{
		Name:    f.Key,
		Columns: []F{f},
	})

}
