package generated

import (
	"encoding/json"
	"reflect"
)

func hasElem(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func getFields(data []byte) []string {
	properties := []string{}
	var genericInterface interface{}
	err := json.Unmarshal(data, &genericInterface)
	if err != nil { //It would be weird to get an error here, since we've unmarshaled once already
		return properties
	}

	switch t := genericInterface.(type) {
	case map[string]interface{}:
		for propertyName := range t {
			properties = append(properties, propertyName)
		}

	}

	return properties
}
