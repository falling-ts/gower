/*
Copyright (c) 2023 Falling TS

This source code uses the MIT license,
You can find the MIT license in the root directory.
*/

package app

import (
	"fmt"
	"gower/app/services"
)

type app struct {
	// Provider service providers
	*services.Services

	// App name. example: `gower`
	Name string

	// App Version. example: `v1.0.0`
	Version string
}

// App public open up
var App = &app{
	Services: services.Get(),
	Name:     "",
	Version:  "",
}

// Get service
func Get(key string) services.Service {
	return App.Get(key)
}

func Run() {
	route := App.Get("route").(*services.Route)
	if err := route.Run(); err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
