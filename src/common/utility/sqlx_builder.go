package utility

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	strField strings.Builder
	strValue strings.Builder
)

func GetNamedStruct(data interface{}) []string {
	value := []string{}
	val := reflect.ValueOf(data)

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			fmt.Println(val.Index(i))
		}
	default:
		for i := 0; i < val.Type().NumField(); i++ {
			if val.Type().Field(i).Tag.Get("db") == "" {
				continue
			}
			value = append(value, val.Type().Field(i).Tag.Get("db"))
		}
	}

	return value
}

func SetInsert(data interface{}, unsetColumn []string) (query string) {
	fieldItem := GetNamedStruct(data)
	strField.Reset()
	strValue.Reset()
	for _, v := range fieldItem {
		if UnsetColumn(unsetColumn, v) {
			continue
		}
		strField.WriteString(v + ",")
		strValue.WriteString(":" + v + ",")
	}

	query = "(" + strings.TrimSuffix(strField.String(), ",") + ")" + " VALUES(" + strings.TrimSuffix(strValue.String(), ",") + ")"
	return
}

func SetUpdate(data interface{}, unsetColumn []string) (query string) {
	fieldItem := GetNamedStruct(data)
	strField.Reset()
	for _, v := range fieldItem {
		if UnsetColumn(unsetColumn, v) {
			continue
		}

		strField.WriteString(v + "=:" + v + ", ")
	}

	query = strings.TrimSuffix(strField.String(), ", ")

	return
}

func UnsetColumn(data []string, column string) (unset bool) {
	for _, v := range data {
		if v == column {
			unset = true
			break
		}
	}
	return
}
