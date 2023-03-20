package app

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"gower/app/http"
	"gower/utils/str"
)

const fnField = `^(\w+?)\(([\w,\s?]+?\w)\)$`

type Rule map[string]any

type Model interface {
	In(request http.Request, r Rule) (Model, error)
	Out(data any, r Rule) (any, error)
	SetModel(i Model)
}

type ModelHandle struct {
	Model `gorm:"-"`
}

// In 数据进来
func (m *ModelHandle) In(request http.Request, r Rule) (Model, error) {
	if err := trans(reflect.ValueOf(m.Model), reflect.ValueOf(request), r); err != nil {
		return nil, err
	}
	return m.Model, nil
}

// Out 数据出去
func (m *ModelHandle) Out(data any, r Rule) (any, error) {
	if err := trans(reflect.ValueOf(data), reflect.ValueOf(m.Model), r); err != nil {
		return nil, err
	}
	return data, nil
}

// SetModel 设置具体模型
func (m *ModelHandle) SetModel(i Model) {
	m.Model = i
}

func trans(dest reflect.Value, src reflect.Value, r Rule) error {
	for k, v := range r {
		var fnParams []string
		fnReg := regexp.MustCompile(fnField)
		matches := fnReg.FindStringSubmatch(k)
		if len(matches) > 0 {
			k = matches[1]
			fnParams = strings.Split(matches[2], ",")
		}

		destValue, err := valueByKey(dest, k)
		if err != nil {
			return err
		}

		rule := reflect.ValueOf(v)
		switch rule.Kind() {
		case reflect.Func:
			args := make([]reflect.Value, len(fnParams))
			for i, param := range fnParams {
				argValue, err := valueByKey(src, strings.Trim(param, " "))
				if err != nil {
					args[i] = argValue
					continue
				}
				argValue = reflect.ValueOf(param)

				args[i] = argValue
			}

			results := rule.Call(args)
			for _, result := range results {
				res := result.Interface()
				if res == nil {
					continue
				}
				if err, ok := res.(error); ok {
					return err
				}
			}

			destValue.Set(results[0])
		case reflect.Map:
			srcValue, err := valueByKey(src, v.(string))
			if err != nil {
				destValue.Set(rule)
				break
			}

			if destValue.Kind() == reflect.Ptr {
				destValue = destValue.Elem()
			}
			if srcValue.Kind() == reflect.Ptr {
				srcValue = srcValue.Elem()
			}

			switch destValue.Kind() {
			case reflect.Array:
				if srcValue.Kind() != reflect.Array {
					destValue.Set(rule)
					break
				}
				for i := 0; i < srcValue.Len(); i++ {
					if err := trans(destValue.Index(i), srcValue.Index(i), r); err != nil {
						return err
					}
				}
			case reflect.Slice:
				if srcValue.Kind() != reflect.Slice {
					destValue.Set(rule)
					break
				}
				if srcValue.Len() == 0 {
					break
				}
				elemType := srcValue.Index(0).Type()
				slice := reflect.MakeSlice(elemType, srcValue.Len(), srcValue.Len())
				destValue.Set(slice)
				for i := 0; i < srcValue.Len(); i++ {
					if err := trans(destValue.Index(i), srcValue.Index(i), r); err != nil {
						return err
					}
				}
			case reflect.Map:
				if err := trans(destValue, srcValue, r); err != nil {
					return err
				}
			}
		case reflect.String:
			srcValue, err := valueByKey(src, v.(string))
			if err != nil {
				destValue.Set(rule)
				break
			}
			destValue.Set(srcValue)
		default:
			destValue.Set(rule)
		}
	}

	return nil
}

func valueByKey(v reflect.Value, k string) (reflect.Value, error) {
	var result reflect.Value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Map:
		result = v.MapIndex(reflect.ValueOf(str.Conv(k).Uppercase()))
	case reflect.Struct:
		result = v.FieldByName(str.Conv(k).Uppercase())
	}
	if result.IsValid() {
		return result, nil
	}

	return result, errors.New("类型错误")
}
