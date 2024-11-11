# Go/Gin Gower Web Starter Framework

![](storage/app/public/images/logo.png)

[中文](README.md)|[English](README_EN.md)

[![benchmark](https://img.shields.io/badge/gower-benchmark-red?style=flat-square&logo=Sencha)](tests/benchmarks/benchmark)
[![actions](https://img.shields.io/badge/github-actions-green?style=flat-square&logo=GitHub)](https://github.com/falling-ts/gower/actions)
[![version](https://img.shields.io/badge/version-0.2.0-yellow?style=flat-square&logo=V)]()

***

Gower is a fast-starting web framework based on [Go/Gin](https://github.com/gin-gonic/gin), with its architectural core ideas mainly inspired by the design philosophy of [Laravel](https://github.com/laravel/laravel). Its directory structure is similar to Laravel's, and its features are mostly analogous. Based on Go/Gin's routing design, Gower ensures basic performance while improving the elegance of code development. Utilizing Go's reflection and type assertion mechanisms, Gower implements dependency injection to perform parameter validation and model initialization before executing the logic, simplifying the code.

Main Features:

*   Command is the essence, integrating command-line tools with the built program
*   Services and service providers, binding keys and functions to achieve dynamic service construction
*   Dependency injection between services, avoiding circular dependency traps
*   Business is the core, with the core content provided by app, and app obtaining service capabilities through service providers
*   Gin route function wrapping, implementing custom controller method parameters and return values, i.e., flexible controllers
*   Injecting request structures into controller methods, automatically validating request parameters
*   Can be used for both non-separated front-end and back-end and separated front-end and back-end
*   Vite is used for front-end library packaging, providing CSS and JS for templates
*   Separate environments for development, testing, and production, each with its own environment file
*   Overall deployment, mainly using Docker for containerized operation, avoiding issues caused by inconsistent environments

System Requirements:

> go >= v1.20
>
> nodejs >= v16.13
>
> pnpm >= v7.0
>
> docker >= v20.10
>
> docker compose >= v2.0
>
> git >= 2.39

## Quick Start

### Install from source \[Recommended]

#### 1. Install from remote compilation

```shell
$ go install -tags cli github.com/falling-ts/gower@latest
```

> Verify the installation: `$ gower --version`

#### 2. Create a project and automatically initialize it

```shell
$ gower create myproject
```

> This will create a project, initialize files, environments, repositories, frontend and backend dependencies, and run benchmark tests.

#### 3. Use Docker

```shell
$ cd myproject
$ ./run-dev
```

#### 4. Debug with Goland

In the `main.go` file, right-click on the green triangle and select "debug". The first time it runs, it will only print the command prompt without running. Then, select "Edit Configurations" above and add "run" to the Program arguments in the `go build gower` that was created. Save and execute the debug.


### Use git install

#### 1. Download

```shell
$ git clone https://github.com/falling-ts/gower.git
or
$ git clone https://gitee.com/falling-ts/gower.git

```

#### 2. Switch to the released version

```shell
$ git checkout v0.6.0
```

> After switching, you can delete the `.git` directory and create your own repository.


#### 3. Install front-end and back-end dependencies

```shell
$ pnpm install
$ go mod tidy
$ go install -tags cli
```

> Note: First, go to [goproxy.cn](https://goproxy.cn/) to configure the acceleration proxy, and then use `go mod tidy`


#### 4. Initialize environment

*   In the root directory, copy the two frontend environment files `.env.test` and `.env.prod`.
*   In the `envs/` directory, copy the two backend environment files `.env.test` and `.env.prod`.
*   Generate the APP and JWT keys.

```shell
$ gower init key
$ gower jwt key
```

#### 5. Run dev development environment with Docker

```shell
$ ./run-dev
```

> Tested on Windows, if there are problems with other systems, please raise issues

#### 6. Without Docker

*   Build front-end

```shell
$ npm run dev
```

> Will build js, css, and images in `public/static`

*   Build back-end and run

```shell
$ go test
$ go install
$ gower run # Execute in the project root directory and add $GOPATH/bin to the environment variable
```

> To package static resources, run `go install -tags tmpl,static`

##### tags:

```
test: Package test environment program files
prod: Package production environment program files
tmpl: Package templates
static: Package static resources
cli: Command line mode

```

> The advantage of packaging these contents is that there is no need to worry about the content to be carried during program migration, as it is all packaged into the program, making the system highly flexible.

## Rapid Development

*   Create a controller
```shell
$ gower make --controller Hello
```
`app/http/controllers/hello_controller.go`
```go
package controllers

import (
    "gower/app"
    "gower/app/http/requests"
    "gower/services"
)

type HelloController struct {
    app.Controller
}

var Hello = new(HelloController)

// Index Get the page
func (*HelloController) Index(req *requests.HelloRequest) (services.Response, error) {
    return res.Ok("home/hello", app.Data{
        "name": req.Name,
    }), nil
}
```
*   Create a request
```shell
$ gower make --request Hello
```
`app\http\requests\hello_request.go`
```go
package requests

import "gower/app"

type HelloRequest struct {
    app.Request

    Name *string `form:"name" json:"name" binding:"required"`
}
```
*   Create a model
```shell
$ gower make --model Hello
```
`app\models\hello.go`
```go
package models

func init() {
    migrate(new(Hello))
}

type Hello struct {
    Model

    Name *string `gorm:"type:string;default:'';comment:Name"`
}
```
> Note: If the command outputs a lot of Debug content, it's because the APP\_MODE in envs/.env.dev is in development mode. Change it to test mode to resolve.
*   Add a route
`routes/web.go`
```go
package routes

import (
    web "gower/app/http/controllers"
    mws "gower/app/http/middlewares"
    "gower/public"
)

func init() {
    // ...

    route.GET("/hello", web.Hello.Index)
}
```
*   Execute the request
```shell
$ curl -i http://localhost:8080/hello?name=Gower
```
## Third-party libraries and content used, and gratitude for open source
```
github.com/golang/go v1.20
github.com/alexedwards/argon2id v0.0.0-20230305115115-4b3c3280a736
github.com/caarlos0/env/v7 v7.0.0
github.com/gin-contrib/cors v1.4.0
github.com/gin-gonic/gin v1.9.0
github.com/go-playground/locales v0.14.1
github.com/go-playground/universal-translator v0.18.1
github.com/go-playground/validator/v10 v10.11.2
github.com/go-sql-driver/mysql v1.7.0
github.com/golang-jwt/jwt/v5 v5.0.0-rc.1
github.com/jaevor/go-nanoid v1.3.0
github.com/joho/godotenv v1.5.1
github.com/patrickmn/go-cache v2.1.0+incompatible
github.com/stretchr/testify v1.8.1
github.com/urfave/cli/v2 v2.25.0
go.uber.org/zap v1.24.0
golang.org/x/crypto v0.7.0
gorm.io/driver/mysql v1.4.7
gorm.io/gorm v1.24.6

github.com/rclone/rclone v1.62.2
github.com/laravel/laravel
github.com/moby/moby
github.com/docker/compose

FROM caddy:2.6
FROM grafana/grafana:9.4.3
FROM grafana/loki:main-0295fd4
FROM mysql/mysql-server:5.7.41
FROM grafana/promtail:main-0295fd4
FROM pingcap/tidb:v6.5.1

nodejs
pnpm
"animate.css": "^4.1.1",
"autoprefixer": "^10.4.13",
"daisyui": "^2.51.2",
"jquery": "^3.6.3",
"js-cookie": "^3.0.1",
"localforage": "^1.10.0",
"postcss": "^8.4.21",
"resize-observer-polyfill": "^1.5.1",
"simplebar": "^6.2.1",
"stylus": "^0.59.0",
"tailwindcss": "^3.2.7",
"vue": "^3.2.47"
"@rollup/plugin-replace": "^5.0.2",
"@types/jquery": "^3.5.16",
"@types/js-cookie": "^3.0.3",
"@types/node": "^18.15.10",
"@types/vue": "^2.0.0",
"@vitejs/plugin-vue": "^4.0.0",
"cross-env": "^7.0.3",
"vite": "^4.1.4"
```

## LICENSE

[MIT License](LICENSE)

