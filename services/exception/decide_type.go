package exception

func decideType(arg any, e *Exception) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		e.Exceptions.SetMsg(err.Error())
		e.RawErr = err.(error)
	case string:
		e.Exceptions.SetMsg(arg.(string))
	default:
		e.Exceptions.SetData(arg)
	}
}
