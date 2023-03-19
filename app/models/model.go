package models

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"gower/app"
	"gower/app/http/requests"
	"gower/utils/str"

	"gorm.io/gorm"
)

const (
	fnField = `^(\w+?)\(([\w,\s?]+?\w)\)$`
)

type Rule map[string]any

type Interface interface {
	In(request requests.Request, r Rule) (Interface, error)
	Out(data any, r Rule) (any, error)
	SetInterface(i Interface)
}

type Model struct {
	Interface `gorm:"-"`
	gorm.Model
}

var db = app.GormDB()

// In 数据进来
func (m *Model) In(request requests.Request, r Rule) (Interface, error) {
	if err := trans(reflect.ValueOf(m.Interface), reflect.ValueOf(request), r); err != nil {
		return nil, err
	}
	return m.Interface, nil
}

// Out 数据出去
func (m *Model) Out(data any, r Rule) (any, error) {
	if err := trans(reflect.ValueOf(data), reflect.ValueOf(m.Interface), r); err != nil {
		return nil, err
	}
	return data, nil
}

// SetInterface 设置具体模型
func (m *Model) SetInterface(i Interface) {
	m.Interface = i
}

func trans(dest reflect.Value, src reflect.Value, r Rule) error {
	for k, v := range r {
		var fnParams []string
		fnReg := regexp.MustCompile(fnField)
		matchs := fnReg.FindStringSubmatch(k)
		if len(matchs) > 0 {
			k = matchs[1]
			fnParams = strings.Split(matchs[2], ",")
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

func migrate(args ...any) {
	if err := db.AutoMigrate(args...); err != nil {
		panic(err)
	}
}
