package response

import (
	"net/http"
)

// DecideType 判定类型
func (s *Service) decideType(arg any) {
	switch arg.(type) {
	case int:
		code := arg.(int)
		if code >= http.StatusOK && code < http.StatusMultipleChoices {
			s.HttpStatus = code
		} else {
			s.Response.Set(arg)
		}
	case string:
		s.Response.Set(arg)
		s.config.HTMLName = arg.(string)
	default:
		s.Response.Set(arg)
		s.config.HTMLData = arg
	}

	s.config.Data = s.Response
}
