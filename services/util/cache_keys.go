package util

import "strings"

// ExcKey 异常缓存 key
func (s *Service) ExcKey() string {
	return strings.Join([]string{"exc", s.Nanoid()}, " ")
}

// BlackTokenKey Token 黑名单缓存 key
func (s *Service) BlackTokenKey(nanoid string) string {
	return strings.Join([]string{"black-token", nanoid}, " ")
}
