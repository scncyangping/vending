package util

import (
	"github.com/globalsign/mgo/bson"
	"reflect"
)

const (
	Nil = iota
	Bool
	Int
	Float
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	Undefined
)

func IsExpectType(i any) int {
	if i == nil {
		return Nil
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.String:
		return String
	case reflect.Map:
		return Map
	case reflect.Array:
		return Array
	case reflect.Struct:
		return Struct
	case reflect.Slice:
		return Slice
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Int
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.Interface:
		return Interface
	case reflect.Bool:
		return Bool
	default:
		return Undefined
	}
}

// Contains 判断obj是否在target中，target支持的类型array,slice,map
func Contains(obj any, target any) bool {
	if target == nil {
		return false
	}
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func DeepCopy(value any) any {
	if valueMap, ok := value.(map[string]any); ok {
		newMap := make(map[string]any)
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]any); ok {
		newSlice := make([]any, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueMap, ok := value.(bson.M); ok {
		newMap := make(bson.M)
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
	}
	return value
}

func FiltrationData(data map[string]any, field []string) map[string]any {
	resultMap := make(map[string]any)
	if data == nil || field == nil || len(data) < 1 || len(field) < 1 {
		return nil
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(data); j++ {
			if Contains(field[i], data) {
				resultMap[field[i]] = data[field[i]]
			}
		}
	}
	return resultMap
}
