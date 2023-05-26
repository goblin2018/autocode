package g

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type empty struct {
}

type Struct interface{}

var Empty = empty{}

func Structs(structs ...Struct) []Struct {
	return structs
}

func (s Struct) ToString() string {
	return ""
}

// WriteStructToFile 将结构体定义写入带有给定文件名的文件。如果 appendToFile 为 true，则追加到现有文件。
func WriteStructToFile(s interface{}, fileName string, appendToFile bool) {
	structContent := BuildStructString(s)

	if appendToFile {
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Failed to open file: %s\n", err)
			return
		}
		defer file.Close()

		// 为了可读性，在结构体之间添加一个空行
		structContent = "\n" + structContent
	} else {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Failed to create file: %s\n", err)
			return
		}
		defer file.Close()
	}

	_, err := file.WriteString(structContent)
	if err != nil {
		fmt.Printf("Failed to write struct to file: %s\n", err)
	}
}

// BuildStructString 构建并返回结构体的定义字符串
func BuildStructString(s interface{}) string {
	structType := reflect.TypeOf(s)
	builder := strings.Builder{}

	fmt.Fprintf(&builder, "type %s struct {\n", structType.Name())

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		// Only retain JSON or Form tags.
		tags := extractJSONOrFormTags(field.Tag)
		t := field.Type
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		fmt.Fprintf(&builder, "    %s %s %s\n", field.Name, t, tags)
	}

	fmt.Fprint(&builder, "}\n")

	return builder.String()
}

// extractJSONOrFormTags extracts json and form tags from a struct field tag.
func extractJSONOrFormTags(tag reflect.StructTag) string {
	var jsonString, formString string

	jsonTag := tag.Get("json")
	if jsonTag != "" {
		jsonString = fmt.Sprintf(`json:"%s"`, jsonTag)
	}

	formTag := tag.Get("form")
	if formTag != "" {
		formString = fmt.Sprintf(`form:"%s"`, formTag)
	}

	result := strings.TrimSpace(jsonString + " " + formString)
	if result != "" {
		result = "`" + result + "`"
	}

	return result
}
