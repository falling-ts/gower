package exception

func decideType(arg any, s *Service) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		s.Exception.Set(err.Error())
		s.RawErr = err.(error)
	case string:
		s.Exception.Set(arg.(string))
	default:
		s.Exception.Set(arg)
	}
}
