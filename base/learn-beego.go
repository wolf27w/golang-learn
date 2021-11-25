//###  beego的使用以及配置

//####beego简介

//beego 是免费、开源的软件，这意味着任何人都可以为其开发和进步贡献力量。beego 源代码目前托管在 Github 上，Github 提供非常容易的途径 fork 项目和合并你的贡献。

//####beego安装

//1. beego 安装升级
// 2. beego 的安装
// beego 的安装是典型的 Go 安装包的形式：

// go get github.com/astaxie/beego
// 常见问题：

// git 没有安装，请自行安装不同平台的 git，如何安装请自行搜索。
// git https 无法获取，请配置本地的 git，关闭 https 验证：

//   git config --global http.sslVerify false
// 无法上网怎么安装 beego，目前没有好的办法，接下来我们会整理一个全包下载，每次发布正式版本都会提供这个全包下载，包含依赖包。

// 3. beego 的升级
// beego 升级分为 go 方式升级和源码下载升级：

// Go 升级,通过该方式用户可以升级 beego 框架，强烈推荐该方式：

//   go get -u github.com/astaxie/beego
// 源码下载升级，用户访问 https://github.com/astaxie/beego ,下载源码，然后覆盖到 $GOPATH/src/github.com/astaxie/beego 目录，然后通过本地执行安装就可以升级了：

//   go install     github.com/astaxie/beego
// 4. beego 的 git 分支
// beego 的 master 分支为相对稳定版本，dev 分支为开发者版本。大致流程如下：

//####。bee工具的使用

//1. bee工具的使用
// 2. bee 工具简介
// bee 工具是一个为了协助快速开发 beego 项目而创建的项目，通过 bee 您可以很容易的进行 beego 项目的创建、热编译、开发、测试、和部署。

// 2.1. bee 工具的安装
// 您可以通过如下的方式安装 bee 工具：

// go get github.com/beego/bee
// 安装完之后，bee 可执行文件默认存放在 $GOPATH/bin 里面，所以您需要把 $GOPATH/bin 添加到您的环境变量中，才可以进行下一步。

// 如何添加环境变量，请自行搜索 如果你本机设置了 GOBIN，那么上面的命令就会安装到 GOBIN 下，请添加 GOBIN 到你的环境变量中

// 2.2. bee 工具命令详解
// 我们在命令行输入 bee，可以看到如下的信息：

// Bee is a Fast and Flexible tool for managing your Beego Web Application.

// Usage:

//     bee command [arguments]

// The commands are:

//     version     show the bee & beego version
//     migrate     run database migrations
//     api         create an api application base on beego framework
//     bale        packs non-Go files to Go source files    
//     new         create an application base on beego framework
//     run         run the app which can hot compile
//     pack        compress an beego project
//     fix         Fixes your application by making it compatible with newer versions of Beego
//     dlv         Start a debugging session using Delve
//     dockerize   Generates a Dockerfile for your Beego application
//     generate    Source code generator
//     hprose      Creates an RPC application based on Hprose and Beego frameworks
//     new         Creates a Beego application
//     pack        Compresses a Beego application into a single file
//     rs          Run customized scripts
//     run         Run the application by starting a local development server
//     server      serving static content over HTTP on port

// Use bee help [command] for more information about a command.
// 2.2.1. new 命令
// new 命令是新建一个 Web 项目，我们在命令行下执行 bee new <项目名> 就可以创建一个新的项目。但是注意该命令必须在 $GOPATH/src 下执行。最后会在 $GOPATH/src 相应目录下生成如下目录结构的项目：

// bee new myproject
// [INFO] Creating application...
// /gopath/src/myproject/
// /gopath/src/myproject/conf/
// /gopath/src/myproject/controllers/
// /gopath/src/myproject/models/
// /gopath/src/myproject/static/
// /gopath/src/myproject/static/js/
// /gopath/src/myproject/static/css/
// /gopath/src/myproject/static/img/
// /gopath/src/myproject/views/
// /gopath/src/myproject/conf/app.conf
// /gopath/src/myproject/controllers/default.go
// /gopath/src/myproject/views/index.tpl
// /gopath/src/myproject/main.go
// 13-11-25 09:50:39 [SUCC] New application successfully created!
// myproject
// ├── conf
// │   └── app.conf
// ├── controllers
// │   └── default.go
// ├── main.go
// ├── models
// ├── routers
// │   └── router.go
// ├── static
// │   ├── css
// │   ├── img
// │   └── js
// ├── tests
// │   └── default_test.go
// └── views
//     └── index.tpl

// 8 directories, 4 files
// 2.2.2. api 命令
// 上面的 new 命令是用来新建 Web 项目，不过很多用户使用 beego 来开发 API 应用。所以这个 api 命令就是用来创建 API 应用的，执行命令之后如下所示：

// bee api apiproject
// create app folder: /gopath/src/apiproject
// create conf: /gopath/src/apiproject/conf
// create controllers: /gopath/src/apiproject/controllers
// create models: /gopath/src/apiproject/models
// create tests: /gopath/src/apiproject/tests
// create conf app.conf: /gopath/src/apiproject/conf/app.conf
// create controllers default.go: /gopath/src/apiproject/controllers/default.go
// create tests default.go: /gopath/src/apiproject/tests/default_test.go
// create models object.go: /gopath/src/apiproject/models/object.go
// create main.go: /gopath/src/apiproject/main.go
// 这个项目的目录结构如下：

// apiproject
// ├── conf
// │   └── app.conf
// ├── controllers
// │   └── object.go
// │   └── user.go
// ├── docs
// │   └── doc.go
// ├── main.go
// ├── models
// │   └── object.go
// │   └── user.go
// ├── routers
// │   └── router.go
// └── tests
//     └── default_test.go
// 从上面的目录我们可以看到和 Web 项目相比，少了 static 和 views 目录，多了一个 test 模块，用来做单元测试的。

// 同时，该命令还支持一些自定义参数自动连接数据库创建相关 model 和 controller: bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:<password>@tcp(127.0.0.1:3306)/test"] 如果 conn 参数为空则创建一个示例项目，否则将基于链接信息链接数据库创建项目。

// 2.2.3. run 命令
// 我们在开发 Go 项目的时候最大的问题是经常需要自己手动去编译再运行，bee run 命令是监控 beego 的项目，通过 fsnotify监控文件系统。但是注意该命令必须在 $GOPATH/src/appname 下执行。 这样我们在开发过程中就可以实时的看到项目修改之后的效果：

// bee run
// 13-11-25 09:53:04 [INFO] Uses 'myproject' as 'appname'
// 13-11-25 09:53:04 [INFO] Initializing watcher...
// 13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/controllers)
// 13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/models)
// 13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject)
// 13-11-25 09:53:04 [INFO] Start building...
// 13-11-25 09:53:16 [SUCC] Build was successful
// 13-11-25 09:53:16 [INFO] Restarting myproject ...
// 13-11-25 09:53:16 [INFO] ./myproject is running...
// 我们打开浏览器就可以看到效果 http://localhost:8080/:


















