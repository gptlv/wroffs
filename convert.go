package main

import "reflect"

func StructToMap(s interface{}) map[string]string {
	result := make(map[string]string)

	val := reflect.ValueOf(s)
	typ := val.Type()

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	if val.Kind() != reflect.Struct {
		panic("expected a struct or pointer to a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("csv")

		if tag != "" {
			result[tag] = val.Field(i).String()
		}
	}

	return result
}
