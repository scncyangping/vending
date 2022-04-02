package util

import (
	"encoding/json"
	"reflect"
	"vending/config/log"
)

/**
 * 转换从redis获取的数据
 * @param   base {interface{}} 结构体参数
 * @returns d   {map[string]interface{}} 转换后的map
 */
func ConvertStringToMap(base map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for k, v := range base {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(v), &dat); err == nil {
			resultMap[k] = dat
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}

/**
 * 结构体转map
 * @param   obj {interface{}} 结构体参数
 * @returns d   {map[string]interface{}} 转换后的map
 * @returns err {error} 				 错误
 */
func ConvertStructToMap(obj interface{}) (d map[string]interface{}, err error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	d = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		d[t.Field(i).Name] = v.Field(i).Interface()
	}
	err = nil
	return
}

/**
 * date : 2019/5/7
 * author : yangping
 * desc : 结构体数据拷贝
 */
func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		log.Logger.Error("Fatal error:type of parameters must be Ptr of value")
		return
	}
	if srcv.IsNil() || dstv.IsNil() {
		log.Logger.Error("Fatal error:value of parameters should not be nil")
		return
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	fields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range fields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func DeepFields(baseType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < baseType.NumField(); i++ {
		v := baseType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}
