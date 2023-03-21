package app

import (
	"errors"
	"gower/utils/str"
	"reflect"
)

const fnField = `^(\w+?)\(([\w,\s?]+?\w)\)$`

type Rule map[string]any

type Model interface {
	In(request RequestIFace, r Rule) (Model, error)
	Out(data any, r Rule) (any, error)
	SetModel(i Model)
}

type ModelHandle struct {
	Model `gorm:"-"`
}

// In 数据进来
func (m *ModelHandle) In(request RequestIFace, r Rule) (Model, error) {
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

func trans(dest reflect.Value, src reflect.Value, r map[string]any) error {
	if dest.Kind() == reflect.Ptr {
		dest = dest.Elem()
	}
	if src.Kind() == reflect.Ptr {
		src = src.Elem()
	}

	_, ok := r["_other"]
	if ok {
		destType := dest.Type()
		switch destType.Kind() {
		case reflect.Struct:
			for i := 0; i < destType.NumField(); i++ {
				fieldType := destType.Field(i)
				fieldName := fieldType.Name
				field := dest.FieldByName(fieldName)
				if _, ok = r[fieldName]; ok {
					continue
				}
				if _, ok = r[str.Conv(fieldName).Lowercase()]; ok {
					continue
				}

				switch src.Kind() {
				case reflect.Struct:
					srcField := src.FieldByName(fieldName)
					if !srcField.IsValid() {
						continue
					}
					if field.Kind() != srcField.Kind() {
						continue
					}
					if field.Kind() == reflect.Ptr && field.Elem().Kind() != srcField.Elem().Kind() {
						continue
					}
				case reflect.Map:
					result := src.MapIndex(reflect.ValueOf(fieldName))
					if result.IsValid() {
						r[fieldName] = fieldName
						continue
					}
					result = src.MapIndex(reflect.ValueOf(str.Conv(fieldName).Lowercase()))
					if result.IsValid() {
						r[fieldName] = fieldName
						continue
					}

					continue
				}

				r[fieldName] = fieldName
			}
		case reflect.Map:
			keys := dest.MapKeys()
			if len(keys) == 0 {
				switch src.Kind() {
				case reflect.Struct:
					srcType := src.Type()
					for i := 0; i < srcType.NumField(); i++ {
						fieldName := srcType.Field(i).Name
						if _, ok = r[fieldName]; ok {
							continue
						}
						if _, ok = r[str.Conv(fieldName).Lowercase()]; ok {
							continue
						}

						r[fieldName] = fieldName
					}
				case reflect.Map:
					srcKeys := src.MapKeys()
					for _, key := range srcKeys {
						fieldName := key.String()
						if _, ok = r[fieldName]; ok {
							continue
						}
						if _, ok = r[str.Conv(fieldName).Lowercase()]; ok {
							continue
						}

						r[fieldName] = fieldName
					}
				}
			} else {
				for _, key := range keys {
					fieldName := key.String()
					if _, ok = r[fieldName]; ok {
						continue
					}
					if _, ok = r[str.Conv(fieldName).Lowercase()]; ok {
						continue
					}

					switch src.Kind() {
					case reflect.Struct:
						srcField := src.FieldByName(fieldName)
						if !srcField.IsValid() {
							continue
						}
					}

					r[fieldName] = fieldName
				}
			}
		}

		delete(r, "_other")
	}

	var (
		destValue reflect.Value
		srcValue  reflect.Value
		argValue  reflect.Value
		err       error
	)
	for k, v := range r {
		destValue, err = valueByKey(dest, k)
		if err != nil {
			return err
		}

		rule := reflect.ValueOf(v)
		switch rule.Kind() {
		case reflect.Func:
			argValue, err = valueByKey(src, k)
			if err != nil {
				return err
			}
			arg := []reflect.Value{argValue}
			results := rule.Call(arg)
			for _, result := range results {
				res := result.Interface()
				if res == nil {
					continue
				}
				switch res.(type) {
				case error:
					return res.(error)
				default:
					destValue.Set(results[0])
				}
			}
		case reflect.Struct:
			argsValue := rule.FieldByName("Args")
			funcValue := rule.FieldByName("Func")
			if !argsValue.IsValid() {
				break
			}
			if !funcValue.IsValid() {
				break
			}

			var args []string
			args, ok = argsValue.Interface().([]string)
			if !ok {
				break
			}
			if funcValue.Kind() != reflect.Func {
				break
			}

			realArgs := make([]reflect.Value, len(args))
			for i, arg := range args {
				argValue, err = valueByKey(src, arg)
				if err != nil {
					realArgs[i] = argValue
					continue
				}
				argValue = reflect.ValueOf(arg)

				realArgs[i] = argValue
			}

			results := funcValue.Call(realArgs)
			for _, result := range results {
				res := result.Interface()
				if res == nil {
					continue
				}
				switch res.(type) {
				case error:
					return res.(error)
				default:
					destValue.Set(results[0])
				}
			}
		case reflect.Map:
			srcValue, err = valueByKey(src, v.(string))
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
					if err = trans(destValue.Index(i), srcValue.Index(i), r); err != nil {
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
				makeSlice := reflect.MakeSlice(elemType, srcValue.Len(), srcValue.Len())
				destValue.Set(makeSlice)
				for i := 0; i < srcValue.Len(); i++ {
					if err = trans(destValue.Index(i), srcValue.Index(i), r); err != nil {
						return err
					}
				}
			case reflect.Map:
				if err = trans(destValue, srcValue, r); err != nil {
					return err
				}
			}
		case reflect.String:
			srcValue, err = valueByKey(src, v.(string))
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

	switch v.Kind() {
	case reflect.Map:
		result = v.MapIndex(reflect.ValueOf(k))
		if result.IsValid() {
			break
		}
		result = v.MapIndex(reflect.ValueOf(str.Conv(k).Lowercase()))
		if result.IsValid() {
			break
		}
		result = v.MapIndex(reflect.ValueOf(str.Conv(k).Uppercase()))
		if result.IsValid() {
			break
		}
		result = reflect.ValueOf(new(any))
		v.SetMapIndex(reflect.ValueOf(str.Conv(k).Lowercase()), result)
	case reflect.Struct:
		result = v.FieldByName(str.Conv(k).Uppercase())
	}
	if result.IsValid() {
		return result, nil
	}

	return result, errors.New("类型错误")
}
