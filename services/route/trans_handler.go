package route

import (
	"gower/app/responses"
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
	case func() error:
		if err := handler.(func() error)(); err != nil {
			handleError(err, c)
		}
		return true
	case func() services.Response:
		handleResponse(handler.(func() services.Response)(), c)
		return true
	case func() (services.Response, error):
		response, err := handler.(func() (services.Response, error))()
		if err != nil {
			handleError(err, c)
		} else {
			handleResponse(response, c)
		}
		return true
	default:
		handlerType := reflect.TypeOf(handler).Kind()
		if handlerType != reflect.Func {
			if args, ok := handler.([]any); ok {
				handleResponse(response(new(responses.Responses), args...), c)
			} else {
				handleResponse(response(new(responses.Responses), handler), c)
			}
			return true
		}
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
				argValue = reflect.New(argType)
				if requestMethod(argValue, c) {
					return true
				}

				argValue = argValue.Elem()
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
					if requestMethod(argValue, c) {
						return true
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

	return handleResults(handleValue.Call(args), c)
}

func requestMethod(value reflect.Value, c *gin.Context) bool {
	if SetContext := value.MethodByName("SetContext"); SetContext.IsValid() {
		SetContext.Call([]reflect.Value{
			reflect.ValueOf(c),
		})
	}
	if Validate := value.MethodByName("Validate"); Validate.IsValid() {
		err := Validate.Call([]reflect.Value{
			reflect.ValueOf(c),
			value,
		})[0].Interface()
		if err != nil {
			handleException(err.(services.Exceptions), c)
			return true
		}
	}

	return false
}

func handleResults(results []reflect.Value, c *gin.Context) bool {
	r := make([]any, 0)
	for _, result := range results {
		res := result.Interface()
		if res == nil {
			continue
		}
		switch res.(type) {
		case services.Response:
			handleResponse(res.(services.Response), c)
			return true
		case error:
			handleError(res.(error), c)
			return true
		default:
			if _res, ok := res.([]any); ok {
				r = append(r, _res...)
			} else {
				r = append(r, res)
			}
		}
	}

	handleResponse(response(new(responses.Responses), r...), c)
	return true
}

func handleError(err error, c *gin.Context) {
	if e, ok := err.(services.Exceptions); ok {
		handleException(e, c)
	} else {
		handleException(exceptions.New(http.StatusBadRequest, err), c)
	}
}
