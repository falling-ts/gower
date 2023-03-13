package route

import (
	"fmt"
	"gower/services"
	"net/http"
	"path"
	"reflect"

	"github.com/gin-gonic/gin"
)

func transHandler(handler services.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		setWriter(c)
		if handle, ok := handler.(func(*gin.Context)); ok {
			handle(c)
			return
		}
		if handle, ok := handler.(func(Context)); ok {
			handle(c)
			return
		}

		handleValue := reflect.ValueOf(handler)
		handleType := handleValue.Type()

		args := make([]reflect.Value, handleType.NumIn())
		for i := 0; i < handleType.NumIn(); i++ {
			argType := handleType.In(i)
			var argValue reflect.Value

			switch argType.Kind() {
			case reflect.Struct:
				pkgPath := argType.PkgPath()
				pkg := path.Base(pkgPath)

				switch pkg {
				case "requests":
					argValue = reflect.New(argType).Elem()
					if SetContext := argValue.MethodByName("SetContext"); SetContext.IsValid() {
						SetContext.Call([]reflect.Value{
							reflect.ValueOf(c),
						})
					}
					if Validate := argValue.MethodByName("Validate"); Validate.IsValid() {
						excp := Validate.Call([]reflect.Value{
							reflect.ValueOf(c),
							argValue,
						})

						fmt.Println(excp)
					}
				default:
					handleException(exceptions.New(http.StatusBadRequest, "参数声明错误"), c)
				}
			case reflect.Ptr:
				argType = argType.Elem()
				switch argType.Kind() {
				case reflect.Struct:
					pkgPath := argType.PkgPath()
					pkg := path.Base(pkgPath)

					switch pkg {
					case "requests":
						argValue = reflect.New(argType)
						//if SetContext := argValue.MethodByName("SetContext"); SetContext.IsValid() {
						//	SetContext.Call([]reflect.Value{
						//		reflect.ValueOf(c),
						//	})
						//}
						//if Validate := argValue.MethodByName("Validate"); Validate.IsValid() {
						//	excp := Validate.Call([]reflect.Value{
						//		reflect.ValueOf(c),
						//		argValue,
						//	})
						//
						//	fmt.Println(excp)
						//}
					default:
						handleException(exceptions.New(http.StatusBadRequest, "参数声明错误"), c)
					}
				}
			default:
				panic("控制器方法设计错误")
			}

			args[i] = argValue
		}

		handleValue.Call(args)
	}
}
