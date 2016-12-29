package utils

import (
	"reflect"
	"strings"
)

func Contains(s interface{}, elem interface{}) bool {
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

func Intersection(ss1 []string, ss2 []string) []string {
	var intersection []string
	for _, s := range ss1 {
		if Contains(ss2, s) {
			intersection = append(intersection, s)
		}
	}
	return intersection
}

func TitleArray(ss []string) []string {
	ret := []string{}
	for _, v := range ss {
		ret = append(ret, strings.Title(v))
	}
	return ret
}

func Extend(old interface{}, new interface{}, properties []string) {
	rOld := reflect.ValueOf(old).Elem()
	rNew := reflect.ValueOf(new).Elem()
	for _, propName := range properties {
		rOldProp := rOld.FieldByName(propName)
		rNewProp := rNew.FieldByName(propName)
		rOldProp.Set(rNewProp)
	}
}
