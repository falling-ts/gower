[TOC]

# Go/Gin Gower Web 单体模式

![](storage/app/public/images/logo.png)

[中文](README.md)|[English](README_EN.md)

[![license](https://img.shields.io/badge/license-MIT-green?style=flat-square&logo=MIT)](LICENSE) [![benchmark](https://img.shields.io/badge/gower-benchmark-red?style=flat-square&logo=Sencha)](tests/benchmarks/benchmark) [![actions](https://img.shields.io/badge/github-actions-green?style=flat-square&logo=GitHub)](https://github.com/falling-ts/gower/actions) [![version](https://img.shields.io/badge/version-0.7.1-yellow?style=flat-square&logo=V)]()

---

Gower 是基于 [Go/Gin](https://github.com/gin-gonic/gin) 的 Web 快速启动框架, 架构核心思想主要借鉴 [Laravel](https://github.com/laravel/laravel) 的设计理念. 目录结构与 Laravel 大同小异, 功能基本类似. 基于 Go/Gin 的路由设计, 在保证基本性能的前提下, 尽量提高代码开发的优雅性, 借助 Go 的反射与类型断言机制, 通过实现依赖注入的功能, 将参数校验/模型初始化放在逻辑之前, 很好的简化了代码.

单体模式, 只能占用一个端口来提供一组服务.

Go/Gin Gower Workspace 【[gower-work](https://github.com/falling-ts/gower-work)】 能够组装多个单体模式, 实现微服务架构

主要特性:

- 命令即本体, 命令行工具与构建的程序相结合
- 服务与服务提供者, 通过 key 与函数绑定, 实现动态服务构建
- 服务间通过依赖注入, 避免环形依赖陷阱
- 业务即核心, 核心内容通过 app 提供服务能力, app 通过服务提供者获得服务能力
- Gin 路由函数包装, 实现控制器方法参数与返回值自定义功能, 即自由控制器
- 控制器方法注入请求结构体, 实现自动验证请求参数
- 非前后分离, 也可用作前后分离
- 前端借助 Vite 实现库打包模式, 为模板提供 css 与 js
- 整体环境, 分开发/测试/生产, 前后端各有自己的环境文件
- 整体发布, 主要由 Docker 提供容器化运行, 主要好处是避免了环境不一致带来的问题

系统要求:

> go >= v1.23
>
> nodejs >= v16.13
>
> pnpm >= v7.0
>
> docker >= v20.10 [非必要]
>
> docker compose >= v2.0 [非必要]
>
> git >= 2.39
>
> gradle == 8.10.2
>
> jdk >= 23

## 快速开始[单体模式]

### 源码安装[推荐]

#### 1.执行远程编译安装

```shell
$ go install -tags cli github.com/falling-ts/gower@latest
```

> 验证结果: `$ gower --version`
>
> 当前安装在系统全局环境中

#### 2.创建项目, 自动初始化

```shell
$ gower create my-project
```

### 运行项目

> 将执行创建项目, 初始化文件, 环境, 仓库, 前后端依赖, 执行基准测试

#### 1.使用 Docker

```shell
$ cd my-project
$ ./docker/run-dev
```

#### 2.使用 GoLand 进行 Debug

用 GoLand 打开项目后, 找到 Gower Run 运行配置, 修改工作目录, 并选择模块, 最后 Debug 运行, 就可以进行断点调试了.

#### 3.使用 gradle 运行

- 单项目模式需要将 `build.gradle` 开头的 `"org.hidetake.ssh"` 插件注释去掉
- 提前在 GoLand 中安装好 gradle 插件
- 修改 `build.gradle`, 去掉开头插件引用的注释
- 第一次使用 GoLand 打开 my-project 时, 会提醒 `找到Gradle 'my-project' 构建脚本`, 然后点击 `加载 Gradle 项目`, 会初始化 gradle 构建体系
- 最后, 在右侧 gradle 任务中找到 dev 下的 Run, 运行它即可.

### 使用 Git 安装

#### 1.下载

```shell
$ git clone -b v0.7.1 --single-branch --depth 1 https://github.com/falling-ts/gower.git
或
$ git clone -b v0.7.1 --single-branch --depth 1 https://gitee.com/falling-ts/gower.git
```

> 切换完成, 你可以删除 `.git` 目录, 自行创建自己的仓库

#### 2.安装前后端依赖

```shell
$ pnpm install
$ go mod tidy
$ go install -tags cli
```

> 注意: 先到 [goproxy.cn](https://goproxy.cn) 配置加速代理, 再使用 `go mod tidy`

#### 4.初始化环境

在根目录下, 复制出 `.env.test` 和 `.env.prod` 两个前端环境文件

在 `envs/` 目录下, 复制出 `.env.test` 和 `.env.prod` 两个后端环境文件

生成 APP 和 JWT 的密钥


```shell
$ gower init key
$ gower jwt key
```

#### 5.通过 Docker 运行 dev 开发环境

```shell
$ ./docker/run-dev
```

> windows 已测试通过, 其它系统有问题, 请提 issues

#### 6.不使用 Docker

##### 构建前端

```shell
$ npm run dev
```

> 将在 `public/static` 下构建出 js 和 css 以及 images 内容

##### 构建后端与运行

```shell
$ go test
$ go install
$ gower run # 要在项目根目录下执行, 记得把 $GOPATH/bin 加入环境变量
```

> 如果需要打包静态资源请执行 `go install -tags tmpl,static`

###### tags:

```
test: 打包测试环境的程序文件
prod: 打包生成环境的程序文件
tmpl: 打包模板
static: 打包静态资源
cli: 命令行模式
```

> 打包这些内容好处是无需关心程序迁移时, 需要携带的内容, 因为都打包进程序了, 一个文件就是整套系统, 灵活性极高

## 快速开发

### 创建控制器

```shell
$ gower make --controller Hello
```

`app/http/controllers/hello_controller.go`

```shell
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

// Index 获取页面
func (*HelloController) Index(req *requests.HelloRequest) (services.Response, error) {
    return res.Ok("home/hello", app.Data{
        "name": req.Name,
    }), nil
}

```

### 创建请求

```shell
$ gower make --request Hello
```

`app\http\requests\hello_request.go`

```shell
package requests

import "gower/app"

type HelloRequest struct {
    app.Request

    Name *string `form:"name" json:"name" binding:"required"`
}
```

### 创建模型

```shell
$ gower make --model Hello
```

`app\models\hello.go`

```shell
package models

func init() {
    migrate(new(Hello))
}

type Hello struct {
    Model

    Name *string `gorm:"type:string;default:'';comment:名称"`
}
```

> 注: 如果命令输出很多 Debug 内容, 那是因为 envs/.env.dev 的 APP_MODE 是开发模式, 修改为测试模式就可以了

### 添加路由

`routes/web.go`

```shell
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

### 执行请求

```shell
$ curl -i http://localhost:8080/hello?name=Gower
```

## 使用的第三方库和内容, 同时表达对开源的感谢

```

github.com/alexedwards/argon2id v1.0.0
github.com/caarlos0/env/v7 v7.1.0
github.com/gin-contrib/cors v1.7.2
github.com/gin-gonic/gin v1.10.0
github.com/glebarez/sqlite v1.11.0
github.com/go-playground/locales v0.14.1
github.com/go-playground/universal-translator v0.18.1
github.com/go-playground/validator/v10 v10.22.1
github.com/go-sql-driver/mysql v1.8.1
github.com/golang-jwt/jwt/v5 v5.2.1
github.com/jaevor/go-nanoid v1.4.0
github.com/joho/godotenv v1.5.1
github.com/patrickmn/go-cache v2.1.0+incompatible
github.com/stretchr/testify v1.9.0
github.com/urfave/cli/v2 v2.27.5
go.uber.org/zap v1.27.0
golang.org/x/crypto v0.28.0
gorm.io/driver/mysql v1.5.7
gorm.io/gorm v1.25.12

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
"jssha": "^3.3.0",
"postcss": "8.4.31",
"resize-observer-polyfill": "^1.5.1",
"simplebar": "^6.2.1",
"stylus": "^0.59.0",
"tailwindcss": "^3.2.7",
"vue": "^3.2.47"
"@rollup/plugin-replace": "^5.0.2",
"@types/crypto-js": "^4.1.1",
"@types/jquery": "^3.5.16",
"@types/js-cookie": "^3.0.3",
"@types/node": "^18.15.10",
"@types/vue": "^2.0.0",
"@vitejs/plugin-vue": "^4.0.0",
"cross-env": "^7.0.3",
"vite": "5.4.6"
```

## 文档

[文档地址](https://falling-ts.github.io/gower-docs)

## LICENSE

[MIT License](LICENSE)

## 主页页面

![](.github/images/home.png)

## 示例主题

通过修改 `.env.xxx` 的 `VIEW_THEME`, 详情见 [DaisyUI](https://daisyui.com/docs/themes/)

### cupcake

![](.github/images/cupcake.png)

### forest

![](.github/images/forest.png)

### halloween

![](.github/images/halloween.png)

### lofi

![](.github/images/lofi.png)

### synthwave

![](.github/images/synthwave.png)
