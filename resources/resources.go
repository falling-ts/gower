package resources

import (
	"embed"
	"html/template"
	"reflect"

	"gitee.com/falling-ts/gower/app"
)

var (
	route = app.Route()
	Tmpl  *embed.FS
)

func init() {
	route.SetFuncMap(template.FuncMap{
		"assertAnySlice": func(v any) []any {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Slice:
				val := reflect.ValueOf(v)
				result := make([]interface{}, val.Len())
				for i := 0; i < val.Len(); i++ {
					result[i] = val.Index(i).Interface()
				}
				return result
			default:
				return []any{}
			}
		},
		"until": func(n int) []int {
			res := make([]int, n)
			for i := range res {
				res[i] = i + 1
			}
			return res
		},
	})

	err := route.LoadHTMLGlobs(
		"resources/views/*.tmpl",
		"resources/views/**/*.tmpl",
		"resources/views/**/**/*.tmpl",
		"resources/views/**/**/**/*.tmpl",
		"resources/views/**/**/**/**/**/*.tmpl",
		"resources/views/**/**/**/**/**/**/*.tmpl",
		"resources/views/**/**/**/**/**/**/**/*.tmpl",
		"resources/views/**/**/**/**/**/**/**/**/*.tmpl",
		"resources/views/**/**/**/**/**/**/**/**/**/*.tmpl",
	)
	if err != nil {
		if Tmpl == nil {
			panic("没有模板内容可加载")
		}
		route.LoadHTMLFS(Tmpl,
			"views/*.tmpl",
			"views/**/*.tmpl",
			"views/**/**/*.tmpl",
			"views/**/**/**/*.tmpl",
			"views/**/**/**/**/*.tmpl",
			"views/**/**/**/**/**/**/*.tmpl",
			"views/**/**/**/**/**/**/**/*.tmpl",
			"views/**/**/**/**/**/**/**/**/*.tmpl",
			"views/**/**/**/**/**/**/**/**/**/*.tmpl")
	}
}
