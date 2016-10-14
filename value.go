package m2rcache

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func fullValue(dst interface{}, fields []string, value []interface{}) {
	v := reflect.ValueOf(dst)
	t := reflect.TypeOf(dst)
	if t.Kind() != reflect.Ptr {
		panic("dst参数必须是地址")
	}
	ve := v.Elem()
	switch ve.Kind() {
	case reflect.Slice:
		vFullArr(ve, value)
	case reflect.Map:
		vFullMap(ve, fields, value)
	case reflect.Struct:
		vFullStruct(ve, fields, value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vFullInt(ve, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vFullUint(ve, value)
	case reflect.Float32, reflect.Float64:
		vFullFloat(ve, value)

	case reflect.String:
		vFullString(ve, value)

	default:
		panic("不支持该类型" + ve.Kind().String())
	}

}
func vFullArr(v reflect.Value, value []interface{}) {

	// Grow slice if necessary
	l := len(value)
	newv := reflect.MakeSlice(v.Type(), l, l)
	reflect.Copy(newv, reflect.ValueOf(value))
	v.Set(newv)

}
func vFullMap(v reflect.Value, fields []string, value []interface{}) {

	newv := reflect.MakeMap(v.Type())
	for i, v := range fields {
		newv.SetMapIndex(reflect.ValueOf(v), reflect.ValueOf(value[i]))
	}
	v.Set(newv)

}
func vFullStruct(v reflect.Value, fields []string, value []interface{}) {

	for i, f := range fields {
		uf := strings.ToUpper(f)
		mv := vFindFieldValue(v, uf)
		vFullField(mv, value[i])
	}

}
func vFullString(v reflect.Value, value interface{}) {
	v.SetString(fmt.Sprintf("%v", value))
}
func vFullInt(m reflect.Value, value interface{}) {
	switch value.(type) {
	case string:
		tmp := value.(string)
		iv, err := strconv.ParseInt(tmp, 10, 64)
		if err == nil {
			m.SetInt(iv)
		}
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value)
		m.SetInt(v.Int())

	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value)
		m.SetInt(int64(v.Uint()))
	case float32, float64:
		v := reflect.ValueOf(value)
		m.SetInt(int64(v.Float()))
	}
}
func vFullUint(m reflect.Value, value interface{}) {
	switch value.(type) {
	case string:
		tmp := value.(string)
		iv, err := strconv.ParseUint(tmp, 10, 64)
		if err == nil {
			m.SetUint(iv)
		}
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value)
		m.SetUint(uint64(v.Int()))

	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value)
		m.SetUint(v.Uint())
	case float32, float64:
		v := reflect.ValueOf(value)
		m.SetUint(uint64(v.Float()))
	}
}
func vFullFloat(m reflect.Value, value interface{}) {
	switch value.(type) {
	case string:
		tmp := value.(string)
		iv, err := strconv.ParseFloat(tmp, 64)
		if err == nil {
			m.SetFloat(iv)
		}
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value)
		m.SetFloat(float64(v.Int()))

	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value)
		m.SetFloat(float64(v.Uint()))
	case float32, float64:
		v := reflect.ValueOf(value)
		m.SetFloat(v.Float())
	}
}
func vFindFieldValue(v reflect.Value, field string) reflect.Value {
	return v.FieldByNameFunc(func(m string) bool {
		return strings.Compare(strings.ToUpper(m), field) == 0
	})
}
func vFullField(m reflect.Value, value interface{}) {

	switch m.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vFullInt(m, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vFullUint(m, value)
	case reflect.Float32, reflect.Float64:
		vFullFloat(m, value)

	case reflect.String:
		vFullString(m, value)
	default:
		panic("不支持该类型" + m.Kind().String())
	}
}
