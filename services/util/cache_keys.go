package util

import "strings"

// ExcpKey 异常缓存 key
func (s *Service) ExcpKey() string {
	return strings.Join([]string{"excp", s.Nanoid()}, " ")
}

// BlackTokenKey Token 黑名单缓存 key
func (s *Service) BlackTokenKey(nanoid string) string {
	return strings.Join([]string{"black-token", nanoid}, " ")
}
