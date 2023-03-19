package route

import (
	"errors"
	"gorm.io/gorm"
	"gower/services"
	"net/http"
	"path"
	"reflect"

	"github.com/gin-gonic/gin"
)

func transHandler(handler services.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if notUseReflect(handler, c) {
			return
		}
		if useReflect(handler, c) {
			return
		}
		exception.New(http.StatusBadRequest, "控制器方法错误.").Handle(c)
	}
}

func notUseReflect(handler services.Handler, c *gin.Context) bool {
	switch handler.(type) {
	case func(*gin.Context):
		handler.(func(*gin.Context))(c)
		return true
	case func(*gin.Context) error:
		if err := handler.(func(*gin.Context) error)(c); err != nil {
			return handleError(err, c)
		}
		return true
	case func(*gin.Context) services.Response:
		return handler.(func(*gin.Context) services.Response)(c).Handle(c)
	case func(*gin.Context) (services.Response, error):
		res, err := handler.(func(*gin.Context) (services.Response, error))(c)
		if err != nil {
			return handleError(err, c)
		} else {
			return res.Handle(c)
		}
	case func() error:
		if err := handler.(func() error)(); err != nil {
			return handleError(err, c)
		}
		return true
	case func() services.Response:
		return handler.(func() services.Response)().Handle(c)
	case func() (services.Response, error):
		res, err := handler.(func() (services.Response, error))()
		if err != nil {
			return handleError(err, c)
		} else {
			return res.Handle(c)
		}
	default:
		handlerType := reflect.TypeOf(handler).Kind()
		if handlerType != reflect.Func {
			if args, ok := handler.([]any); ok {
				return response.New(http.StatusOK, args...).Handle(c)
			} else {
				return response.New(http.StatusOK, handler).Handle(c)
			}
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
			case "models":
				argValue = reflect.New(argType)
				if injectDataById(reflect.New(argType), c) {
					return true
				}
				setModelInterface(argValue)

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
				case "models":
					argValue = reflect.New(argType)
					if injectDataById(reflect.New(argType), c) {
						return true
					}
					setModelInterface(argValue)
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
			return err.(services.Exception).Handle(c)
		}
	}

	return false
}

func injectDataById(value reflect.Value, c *gin.Context) bool {
	id := c.Param("id")
	if id != "" {
		result := db.First(value.Interface(), id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return exception.New(http.StatusNotFound, "没有找到资源.").Handle(c)
			} else {
				return exception.New(http.StatusBadRequest, result.Error).Handle(c)
			}
		}
	}
	return false
}

func setModelInterface(value reflect.Value) {
	if SetInterface := value.MethodByName("SetInterface"); SetInterface.IsValid() {
		if value.Kind() != reflect.Ptr {
			value = value.Addr()
		}
		SetInterface.Call([]reflect.Value{
			value,
		})
	}
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
			return res.(services.Response).Handle(c)
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

	return response.New(http.StatusOK, r...).Handle(c)
}

func handleError(err error, c *gin.Context) bool {
	if e, ok := err.(services.Exception); ok {
		return e.Handle(c)
	} else {
		return exception.New(http.StatusBadRequest, err).Handle(c)
	}
}
