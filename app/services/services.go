package services

import (
	"fmt"
	"reflect"
	"strings"
)

type Service interface {
	Init()
}

type Map map[string]Service

type Services struct {
	Services Map
}

var services = &Services{
	Services: make(Map),
}

// Get services instance
func Get() *Services {
	return services
}

// Get service
func (s *Services) Get(key string) Service {
	if service, ok := services.Services[key]; ok {
		return service
	}

	panic(fmt.Sprintf("Error: service not exist, key is [%s]", key))
}

func (s *Services) Register(service Service) {
	service.Init()

	name := reflect.TypeOf(service).Elem().Name()
	key := strings.ToLower(name)

	services.Services[key] = service
}
