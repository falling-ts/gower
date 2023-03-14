package utils

type MapStrStr map[string]string

// MapKeys 迭代修改键
func (m MapStrStr) MapKeys(fn func(key string) string) MapStrStr {

	for k, v := range m {
		delete(m, k)
		m[fn(k)] = v
	}

	return m
}
