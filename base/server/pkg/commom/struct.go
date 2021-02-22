package commom

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strings"
)

func ToBson(r interface{}) bson.M {
	result := make(bson.M)
	v := reflect.ValueOf(r)
	t := reflect.TypeOf(r)
	for i := 0; i < v.NumField(); i++ {
		filed := v.Field(i)
		tag := t.Field(i).Tag
		key := tag.Get("bson")
		if key == "" || key == "-" {
			continue
		}
		keys := strings.Split(key, ",")
		if len(keys) > 0 {
			key = keys[0]
		}
		// TODO: 处理字段嵌套问题
		switch filed.Kind() {
		case reflect.String:
			v := filed.String()
			if v != "" {
				result[key] = v
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := filed.Int()
			if v != 0 {
				result[key] = v
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			v := filed.Uint()
			if v != 0 {
				result[key] = v
			}
		case reflect.Float32, reflect.Float64:
			v := filed.Float()
			if v != 0 {
				result[key] = v
			}
		case reflect.Bool:
			v := filed.Bool()
			if v == true {
				result[key] = v
			}
		case reflect.Ptr:
		default:
		}
	}
	return result
}
