package exception

func decideType(arg any, e *Exception) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		e.Exceptions.Set(err.Error())
		e.RawErr = err.(error)
	case string:
		e.Exceptions.Set(arg.(string))
	default:
		e.Exceptions.Set(arg)
	}
}
