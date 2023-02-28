/*
Copyright (c) 2023 Falling TS

This source code uses the MIT license,
You can find the MIT license in the root directory.
*/

package app

import (
	"fmt"

	"gower/app/services"
	_ "gower/resources"
	_ "gower/routes"
)

// Run app
func Run() {
	if err := services.Route.Run(); err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
