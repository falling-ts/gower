package exception

func decideType(arg any, e *Struct) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		e.Content.SetMsg(err.Error())
		e.RawErr = err.(error)
	case string:
		e.Content.SetMsg(arg.(string))
	default:
		e.Content.SetData(arg)
	}
}
