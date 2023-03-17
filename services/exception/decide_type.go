package exception

func (s *Service) decideType(arg any) {
	switch arg.(type) {
	case error:
		err := arg.(error)
		_ = s.Exception.Set(err.Error())
		s.RawErr = err.(error)
	case string:
		_ = s.Exception.Set(arg.(string))
	default:
		_ = s.Exception.Set(arg)
	}
}
