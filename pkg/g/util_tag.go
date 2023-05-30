package g

import (
	"fmt"
	"reflect"
	"strings"
)

// getApiTags extracts json and form tags from a struct field tag.
func getApiTags(tag reflect.StructTag) string {
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
