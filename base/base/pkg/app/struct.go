package app

import (
	"base/tools"
	"context"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	json "github.com/json-iterator/go"
	"reflect"
	"strings"
)

func SetContext(Name string, structName interface{}, filter ...string) (res map[string]string) {
	res = make(map[string]string)
	if Name == "elect" {
		// 传递查询的数据库字段
		res[Name] = GetStructName(structName, filter)
	}
	return
}

func GetContext(ctx context.Context, Name string) string {
	req, _ := current.GetRequestContext(ctx)
	if req[Name] != "" {
		return req[Name]
	}
	return ""
}

// 获取 proto 结构体 JSON 名称
// filter 读取结构体 JSON 后排除的字段名
func GetStructName(structName interface{}, filter []string) string {
	t := reflect.TypeOf(structName)
	result := make([]string, 0)
	if t.Kind() == reflect.Ptr {
		var t interface{}
		_ = json.Unmarshal(Struct2Json(structName), &t)
		if len(filter) > 0 {
			for k, _ := range t.(map[string]interface{}) {
				if tools.ContainsString(filter, k) == -1 {
					result = append(result, k)
				}
			}
		} else {
			for k, _ := range t.(map[string]interface{}) {
				result = append(result, k)
			}
		}
	} else if t.Kind() == reflect.Slice {
		s := reflect.ValueOf(structName)
		if len(filter) > 0 {
			for i := 0; i < s.Len(); i++ {
				jsonName := s.Index(i).Interface().(string)
				if tools.ContainsString(filter, jsonName) == -1 {
					result = append(result, jsonName)
				}
			}
		} else {
			for i := 0; i < s.Len(); i++ {
				result = append(result, s.Index(i).Interface().(string))
			}
		}
	} else if t.Kind() == reflect.Struct {
		fieldNum := t.NumField()
		if len(filter) > 0 {
			for i := 0; i < fieldNum; i++ {
				jsonName := t.Field(i).Tag.Get("json")
				if jsonName != "-" && jsonName != "" {
					comma := strings.Index(jsonName, ",")
					if comma > 1 {
						jsonName = jsonName[:comma]
					}
					if tools.ContainsString(filter, jsonName) == -1 {
						result = append(result, jsonName)
					}
				}
			}
		} else {
			for i := 0; i < fieldNum; i++ {
				jsonName := t.Field(i).Tag.Get("json")
				if jsonName != "-" && jsonName != "" {
					comma := strings.Index(jsonName, ",")
					if comma > 1 {
						jsonName = jsonName[:comma]
					}
					result = append(result, jsonName)
				}
			}
		}
	}
	return strings.Join(result, ", ")
}

func Struct2Json(form interface{}) []byte {
	jsonx, err := json.Marshal(form)
	if err != nil {
		jsonx, err = json.Marshal(map[string]string{"err": err.Error()})
		if err != nil {
			return []byte(err.Error())
		}
	}
	return jsonx
}

func UnmarshalJson(req interface{}, v ...interface{}) (err error) {
	jsonb := Struct2Json(req)
	for _, val := range v {
		err = json.Unmarshal(jsonb, &val)
		if err != nil {
			return err
		}
	}
	return nil
}

// 读取值转换
func Unmarshal(req interface{}, v interface{}) error {
	return json.Unmarshal(Struct2Json(req), &v)
}

// 读取值转换
func UnmarshalElect(req interface{}, v interface{}) (string, error) {
	return GetStructName(req, nil), json.Unmarshal(Struct2Json(req), &v)
}