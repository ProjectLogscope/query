package jsonmap

import (
	"log/slog"
	"reflect"
)

type JSONMap map[string]interface{}

func (j *JSONMap) LogValue() slog.Value {
	if j == nil {
		return slog.Value{}
	}

	attrs := []slog.Attr{}
	for k, v := range *j {
		attrs = append(attrs, slog.Any(k, v))
	}
	return slog.GroupValue(attrs...)
}

func Parse(message interface{}) JSONMap {
	jsonMap := make(map[string]interface{})
	v := reflect.ValueOf(message)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Tag.Get("json")
		if len(fieldName) == 0 {
			continue
		}
		if field.IsNil() {
			continue
		}
		jsonMap[fieldName] = field.Interface()
	}
	return jsonMap
}
