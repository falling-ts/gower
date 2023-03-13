package route

import (
	"fmt"
	"net/http"
	"path"
	"reflect"

	"gower/services"

	"github.com/gin-gonic/gin"
)

func transHandler(handler services.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		setWriter(c)
		if notUseReflect(handler, c) {
			return
		}
		if useReflect(handler, c) {
			return
		}
		handleException(exceptions.New(http.StatusBadRequest, "控制器方法错误."), c)
	}
}

func notUseReflect(handler services.Handler, c *gin.Context) bool {
	switch handler.(type) {
	case func(*gin.Context):
		handler.(func(*gin.Context))(c)
		return true
	case func(Context):
		handler.(func(Context))(c)
		return true
	case func(*gin.Context) error:
		if err := handler.(func(*gin.Context) error)(c); err != nil {
			handleError(err, c)
		}
		return true
	case func(Context) error:
		if err := handler.(func(Context) error)(c); err != nil {
			handleError(err, c)
		}
		return true
	case func(*gin.Context) services.Response:
		handleResponse(handler.(func(*gin.Context) services.Response)(c), c)
		return true
	case func(Context) services.Response:
		handleResponse(handler.(func(Context) services.Response)(c), c)
		return true
	case func(*gin.Context) (services.Response, error):
		response, err := handler.(func(*gin.Context) (services.Response, error))(c)
		if err != nil {
			handleError(err, c)
		} else {
			handleResponse(response, c)
		}
		return true
	case func(Context) (services.Response, error):
		response, err := handler.(func(Context) (services.Response, error))(c)
		if err != nil {
			handleError(err, c)
		} else {
			handleResponse(response, c)
		}
		return true
	}

	return false
}

func useReflect(handler services.Handler, c *gin.Context) bool {
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
				return false
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
					if SetContext := argValue.Elem().MethodByName("SetContext"); SetContext.IsValid() {
						SetContext.Call([]reflect.Value{
							reflect.ValueOf(c),
						})
					}
					if Validate := argValue.Elem().MethodByName("Validate"); Validate.IsValid() {
						excp := Validate.Call([]reflect.Value{
							reflect.ValueOf(c),
							argValue,
						})

						fmt.Println(excp)
					}
				default:
					return false
				}
			}
		default:
			return false
		}

		args[i] = argValue
	}

	results := handleValue.Call(args)
	for _, result := range results {
		res := result.Interface()
		if res == nil {
			continue
		}
		switch res.(type) {
		case services.Response:
			handleResponse(res.(services.Response), c)
		case error:
			handleError(res.(error), c)
		default:
			c.String(http.StatusOK, result.String())
		}
	}
	return true
}

func handleError(err error, c *gin.Context) {
	if e, ok := err.(services.Exceptions); ok {
		handleException(e, c)
	} else {
		handleException(exceptions.New(http.StatusBadRequest, err), c)
	}
}
