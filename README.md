# Go/Gin Gower Web 启动框架

![](storage/app/public/images/logo.png)

[中文](README.md)|[English](README_EN.md)

[![benchmark](https://img.shields.io/badge/gower-benchmark-red?style=flat-square&logo=Sencha)](tests/benchmarks/benchmark)
[![actions](https://img.shields.io/badge/github-actions-green?style=flat-square&logo=GitHub)](https://github.com/falling-ts/gower/actions)

---

Gower 是基于 [Go/Gin](https://github.com/gin-gonic/gin) 的 Web 快速启动框架， 架构核心思想主要借鉴 [Laravel](https://github.com/laravel/laravel) 的设计理念。目录结构与 Laravel 大同小异，功能基本类似。基于 Go/Gin 的路由设计，在保证基本性能的前提下，尽量提高代码开发的优雅性，借助 Go 的反射与类型断言机制，通过实现依赖注入的功能，将参数校验、模型初始化放在逻辑之前，很好的简化了代码。

主要特性：

- 命令即本体，命令行工具与构建的程序相结合
- 服务与服务提供者，通过 key 与函数绑定，实现动态服务构建
- 服务间通过依赖注入，避免环形依赖陷阱
- 业务即核心，核心内容通过 app 提供服务能力，app 通过服务提供者获得服务能力
- Gin 路由函数包装，实现控制器方法参数与返回值自定义功能，即自由控制器
- 控制器方法注入请求结构体，实现自动验证请求参数
- 非前后分离，也可用作前后分离
- 前端借助 Vite 实现库打包模式，为模板提供 css 与 js
- 整体环境，分开发、测试、生产，前后端各有自己的环境文件
- 整体发布，主要由 Docker 提供容器化运行，主要好处是避免了环境不一致带来的问题

系统要求:

> go >= v1.20
>
> nodejs >= v16.13
>
> pnpm >= v7.0
>
> docker >= v20.10
>
> docker compose >= v2.0

## 快速开始

### 1.安装前后端依赖

```shell
$ pnpm install
$ go mod tidy
```
> 注意: 先到 [goproxy.cn](https://goproxy.cn) 配置加速代理，再使用 `go mod tidy`

### 2.通过 Docker 运行 dev 开发环境

```shell
$ ./run-dev
```
> windows 已测试通过，其它系统有问题，请提 issues

### 3.不使用 Docker

- 构建前端

```shell
$ npm run dev
```
> 将在 `public/static` 下构建出 js 和 css 以及 images 内容

- 构建后端与运行

```shell
$ go test
$ go install
$ gower run # 要在项目根目录下执行，记得把 $GOPATH/bin 加入环境变量
```
> 如果需要打包静态资源请执行 `go install -tags tmpl,static`

#### tags:

```
test: 打包测试环境的程序文件
prod: 打包生成环境的程序文件
env: 把环境文件也打包进来
tmpl: 打包模板
static: 打包静态资源
```

> 打包这些内容好处是无需关心程序迁移时，需要携带的内容，因为都打包进程序了，一个文件就是整套系统，灵活性极高
>
> 建议不要打包 env，否则每次修改还需重新打包

## 快速开发

- 创建控制器

```shell
$ gower make --controller Hello
```

`app/http/controllers/hello_controller.go`
```go
package controllers

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/http/requests"
	"github.com/falling-ts/gower/services"
)

type HelloController struct {
	app.Controller
}

var Hello = new(HelloController)

// Index 获取页面
func (*HelloController) Index(req *requests.HelloRequest) (services.Response, error) {
	return res.Ok("home/hello", app.Data{
	    "name": req.Name,
	}), nil
}

```

- 创建请求

```shell
$ gower make --request Hello
```
`app\http\requests\hello_request.go`
```go
package requests

import "github.com/falling-ts/gower/app"

type HelloRequest struct {
	app.Request

	Name *string `form:"name" json:"name" binding:"required"`
}
```

- 创建模型

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

	Name *string `gorm:"type:string;default:'';comment:名称"`
}
```

> 注: 如果命令输出很多 Debug 内容，那是因为 envs/.env.development 的 APP_MODE 是开发模式, 修改为测试模式就可以了

- 添加路由

`routes/web.go`
```go
package routes

import (
	web "github.com/falling-ts/gower/app/http/controllers"
	mws "github.com/falling-ts/gower/app/http/middlewares"
	"github.com/falling-ts/gower/public"
)

func init() {
    // ...

	route.GET("/hello", web.Hello.Index)
}
```

- 执行请求

```shell
$ curl -i http://localhost:8080/hello?name=Gower
```

## 使用的第三方库和内容，同时表达对开源的感谢

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
