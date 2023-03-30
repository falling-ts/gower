package app

import (
	"errors"
	"reflect"
	"time"

	"github.com/falling-ts/gower/utils/slice"
	"github.com/falling-ts/gower/utils/str"
)

type Rule map[string]any
type Skips []string

type Model interface {
	In(request RequestIFace, r Rule) (Model, error)
	Out(r Rule) (any, error)
	SetModel(i Model) Model
}

type ModelHandle struct {
	Model `gorm:"-"`
}

var config = Config()

// In 数据进来
func (m *ModelHandle) In(request RequestIFace, r Rule) (Model, error) {
	model := reflect.Indirect(reflect.ValueOf(m.Model))
	if !model.IsValid() {
		return nil, errors.New("未设置原模型")
	}

	req := reflect.Indirect(reflect.ValueOf(request))
	if err := trans(model, req, r); err != nil {
		return nil, err
	}

	return m.Model, nil
}

// Out 数据出去
func (m *ModelHandle) Out(r Rule) (any, error) {
	mapType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(new(any)).Elem())
	data := reflect.MakeMap(mapType)

	model := reflect.Indirect(reflect.ValueOf(m.Model))
	if !model.IsValid() {
		return nil, errors.New("未设置原模型")
	}

	if err := trans(data, model, r); err != nil {
		return nil, err
	}
	return data.Interface(), nil
}

// SetModel 设置具体模型
func (m *ModelHandle) SetModel(i Model) Model {
	m.Model = i
	return i
}

func trans(dest reflect.Value, src reflect.Value, r map[string]any) error {
	dest = reflect.Indirect(dest)
	src = reflect.Indirect(src)

	s, ok := r["_skips"].([]string)
	if !ok {
		s, ok = r["_skips"].(Skips)
	}
	if ok {
		rawSkips := slice.Strings(s)
		skips := rawSkips.Map(func(s string) string {
			return str.Conv(s).UpCamel()
		})

		destType := dest.Type()
		switch destType.Kind() {
		case reflect.Struct:
			destNoSkips(destType, dest, src, skips, r)
		case reflect.Map:
			keys := dest.MapKeys()
			if len(keys) == 0 {
				switch src.Kind() {
				case reflect.Struct:
					srcNoSkips(src.Type(), src, skips, r)
				case reflect.Map:
					srcKeys := src.MapKeys()
					for _, key := range srcKeys {
						fieldName := key.String()
						if isContinue(fieldName, rawSkips, r) {
							continue
						}

						r[fieldName] = fieldName
					}
				}
			} else {
				for _, key := range keys {
					fieldName := key.String()
					if isContinue(fieldName, rawSkips, r) {
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

		delete(r, "_skips")
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
			arg := make([]reflect.Value, 0)
			if rule.Type().NumIn() == 1 {
				arg = append(arg, src)
			}

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
					setValue(dest, destValue, k, results[0])
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
					setValue(dest, destValue, k, results[0])
				}
			}
		case reflect.Map:
			srcValue, err = valueByKey(src, k)
			if err != nil {
				setValue(dest, destValue, k, rule)
				break
			}

			destValue = reflect.Indirect(destValue)
			srcValue = reflect.Indirect(srcValue)

			switch destValue.Kind() {
			case reflect.Array:
				if srcValue.Kind() != reflect.Array {
					setValue(dest, destValue, k, rule)
					break
				}

				elemType := destValue.Type().Elem()
				for i := 0; i < srcValue.Len(); i++ {
					switch elemType.Kind() {
					case reflect.Ptr:
						destValue.Index(i).Set(reflect.New(elemType))
					case reflect.Map:
						destValue.Index(i).Set(reflect.MakeMap(elemType))
					default:
						destValue.Index(i).Set(reflect.New(elemType).Elem())
					}

					_r, ok := v.(map[string]any)
					if !ok {
						_r, ok = v.(Rule)
					}
					if !ok {
						return errors.New("规则类型错误")
					}
					if err = trans(destValue.Index(i), srcValue.Index(i), _r); err != nil {
						return err
					}
				}
			case reflect.Slice:
				if srcValue.Kind() != reflect.Slice {
					setValue(dest, destValue, k, rule)
					break
				}
				if srcValue.Len() == 0 {
					break
				}

				elemType := destValue.Type().Elem()
				destValue = reflect.MakeSlice(reflect.SliceOf(elemType), srcValue.Len(), srcValue.Len())
				for i := 0; i < srcValue.Len(); i++ {
					switch elemType.Kind() {
					case reflect.Ptr:
						destValue.Index(i).Set(reflect.New(elemType))
					case reflect.Map:
						destValue.Index(i).Set(reflect.MakeMap(elemType))
					default:
						destValue.Index(i).Set(reflect.New(elemType).Elem())
					}

					_r, ok := v.(map[string]any)
					if !ok {
						_r, ok = v.(Rule)
					}
					if !ok {
						return errors.New("规则类型错误")
					}
					if err = trans(destValue.Index(i), srcValue.Index(i), _r); err != nil {
						return err
					}
				}
			case reflect.Map, reflect.Struct:
				_r, ok := v.(map[string]any)
				if !ok {
					_r, ok = v.(Rule)
				}
				if !ok {
					return errors.New("规则类型错误")
				}
				if err = trans(destValue, srcValue, _r); err != nil {
					return err
				}
			default:
				switch srcValue.Kind() {
				case reflect.Array:
					if srcValue.Len() == 0 {
						break
					}

					elemType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(new(any)).Elem())
					array := reflect.ArrayOf(srcValue.Len(), elemType)
					setValue(dest, destValue, k, reflect.New(array).Elem())
					for i := 0; i < srcValue.Len(); i++ {
						destValue.Elem().Index(i).Set(reflect.MakeMap(elemType))

						_r, ok := v.(map[string]any)
						if !ok {
							_r, ok = v.(Rule)
						}
						if !ok {
							return errors.New("规则类型错误")
						}
						if err = trans(destValue.Elem().Index(i), srcValue.Index(i), _r); err != nil {
							return err
						}
					}
				case reflect.Slice:
					if srcValue.Len() == 0 {
						break
					}

					elemType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(new(any)).Elem())
					makeSlice := reflect.MakeSlice(reflect.SliceOf(elemType), srcValue.Len(), srcValue.Len())
					setValue(dest, destValue, k, makeSlice)
					for i := 0; i < srcValue.Len(); i++ {
						destValue.Elem().Index(i).Set(reflect.MakeMap(elemType))

						_r, ok := v.(map[string]any)
						if !ok {
							_r, ok = v.(Rule)
						}
						if !ok {
							return errors.New("规则类型错误")
						}
						if err = trans(destValue.Elem().Index(i), srcValue.Index(i), _r); err != nil {
							return err
						}
					}
				case reflect.Map, reflect.Struct:
					mapType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(new(any)).Elem())
					data := reflect.MakeMap(mapType)
					setValue(dest, destValue, k, data)

					_r, ok := v.(map[string]any)
					if !ok {
						_r, ok = v.(Rule)
					}
					if !ok {
						return errors.New("规则类型错误")
					}
					if err = trans(destValue.Elem(), srcValue, _r); err != nil {
						return err
					}
				}
			}
		case reflect.String:
			srcValue, err = valueByKey(src, v.(string))
			if err != nil {
				setValue(dest, destValue, k, rule)
				break
			}
			if Format := srcValue.MethodByName("Format"); Format.IsValid() {
				srcValue = Format.Call([]reflect.Value{
					reflect.ValueOf(time.DateTime),
				})[0]
			}
			setValue(dest, destValue, k, srcValue)
		default:
			setValue(dest, destValue, k, rule)
		}
	}

	return nil
}

func destNoSkips(destType reflect.Type, dest reflect.Value, src reflect.Value, skips slice.Strings, r Rule) {
	for i := 0; i < destType.NumField(); i++ {
		fieldType := destType.Field(i)
		fieldName := fieldType.Name
		field := dest.FieldByName(fieldName)
		typ := field.Type()

		if fieldType.Tag.Get("gorm") == "-" {
			continue
		}
		if fieldType.Anonymous {
			destNoSkips(typ, field, src, skips, r)
			continue
		}
		if isContinue(fieldName, skips, r) {
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
			if field.Kind() == reflect.Ptr {
				fieldElem := field.Type().Elem()
				srcFieldElem := srcField.Type().Elem()
				if fieldElem.Kind() != srcFieldElem.Kind() {
					continue
				}
			}
		case reflect.Map:
			convFieldName := str.Conv(fieldName)
			result := src.MapIndex(reflect.ValueOf(fieldName))
			if result.IsValid() {
				r[fieldName] = fieldName
				continue
			}
			result = src.MapIndex(reflect.ValueOf(convFieldName.Camel()))
			if result.IsValid() {
				r[fieldName] = fieldName
				continue
			}
			result = src.MapIndex(reflect.ValueOf(convFieldName.Snake()))
			if result.IsValid() {
				r[fieldName] = fieldName
				continue
			}

			continue
		}

		r[fieldName] = fieldName
	}
}

func srcNoSkips(srcType reflect.Type, src reflect.Value, skips slice.Strings, r Rule) {
	for i := 0; i < srcType.NumField(); i++ {
		fieldType := srcType.Field(i)
		fieldName := fieldType.Name
		field := src.FieldByName(fieldName)
		typ := field.Type()

		if fieldType.Tag.Get("gorm") == "-" {
			continue
		}
		if fieldType.Anonymous {
			srcNoSkips(typ, field, skips, r)
			continue
		}
		if isContinue(fieldName, skips, r) {
			continue
		}

		r[fieldName] = fieldName
	}
}

func isContinue(fieldName string, skips slice.Strings, r Rule) bool {
	var ok bool
	convFieldName := str.Conv(fieldName)

	if skips.Has(fieldName) {
		return true
	}
	if _, ok = r[fieldName]; ok {
		return true
	}
	if _, ok = r[convFieldName.Camel()]; ok {
		return true
	}
	if _, ok = r[convFieldName.Snake()]; ok {
		return true
	}

	return false
}

func valueByKey(v reflect.Value, k string) (reflect.Value, error) {
	var result reflect.Value
	k = str.Conv(k).UpCamel()

	switch v.Kind() {
	case reflect.Map:
		convK := str.Conv(k)
		result = v.MapIndex(reflect.ValueOf(k))
		if result.IsValid() {
			break
		}
		result = v.MapIndex(reflect.ValueOf(convK.Camel()))
		if result.IsValid() {
			break
		}
		result = v.MapIndex(reflect.ValueOf(convK.Snake()))
		if result.IsValid() {
			break
		}
		result = reflect.ValueOf(new(any)).Elem()
	case reflect.Struct:
		result = v.FieldByName(k)
	}
	if result.IsValid() {
		return result, nil
	}

	return result, errors.New("类型错误")
}

func setValue(dest reflect.Value, destValue reflect.Value, k string, v reflect.Value) {
	k = str.Conv(k).UpCamel()
	destValue.Set(v)

	if dest.Kind() == reflect.Map {
		switch config.Res.KeyType {
		case "CamelType":
			dest.SetMapIndex(reflect.ValueOf(k), destValue)
		case "camelType":
			dest.SetMapIndex(reflect.ValueOf(str.Conv(k).Camel()), destValue)
		default:
			dest.SetMapIndex(reflect.ValueOf(str.Conv(k).Snake()), destValue)
		}
	}
}
