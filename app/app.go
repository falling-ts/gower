/*
Copyright (c) 2023 Falling TS

该源码使用 MIT 授权协议,
你可以在根目录下找到 MIT 授权协议.
*/

package app

import (
	"fmt"

	"gower/app/services"
	_ "gower/resources"
	_ "gower/routes"
)

// Run 运行系统
func Run() {
	if err := services.Route.Run(); err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
