package route

import "github.com/jaevor/go-nanoid"

func getKey() string {
	genKey, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}
	return genKey()
}
