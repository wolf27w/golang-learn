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

//如果我们修改了 Controller 下面的 default.go 文件，我们就可以看到命令行输出：
//
//13-11-25 10:11:20 [EVEN] "/gopath/src/myproject/controllers/default.go": DELETE|MODIFY
//13-11-25 10:11:20 [INFO] Start building...
//13-11-25 10:11:20 [SKIP] "/gopath/src/myproject/controllers/default.go": CREATE
//13-11-25 10:11:23 [SKIP] "/gopath/src/myproject/controllers/default.go": MODIFY
//13-11-25 10:11:23 [SUCC] Build was successful
//13-11-25 10:11:23 [INFO] Restarting myproject ...
//13-11-25 10:11:23 [INFO] ./myproject is running...
//刷新浏览器我们看到新的修改内容已经输出。
//2.2.4. pack 命令
//pack 目录用来发布应用的时候打包，会把项目打包成 zip 包，这样我们部署的时候直接把打包之后的项目上传，解压就可以部署了：
//
//bee pack
//app path: /gopath/src/apiproject
//GOOS darwin GOARCH amd64
//build apiproject
//build success
//exclude prefix:
//exclude suffix: .go:.DS_Store:.tmp
//file write to `/gopath/src/apiproject/apiproject.tar.gz`
//我们可以看到目录下有如下的压缩文件：
//
//rwxr-xr-x  1 astaxie  staff  8995376 11 25 22:46 apiproject
//-rw-r--r--  1 astaxie  staff  2240288 11 25 22:58 apiproject.tar.gz
//drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 conf
//drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 controllers
//-rw-r--r--  1 astaxie  staff      509 11 25 22:31 main.go
//drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 models
//drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 tests
//2.2.5. bale 命令
//这个命令目前仅限内部使用，具体实现方案未完善，主要用来压缩所有的静态文件变成一个变量申明文件，全部编译到二进制文件里面，用户发布的时候携带静态文件，包括 js、css、img 和 views。最后在启动运行时进行非覆盖式的自解压。
//
//2.2.6. version 命令
//这个命令是动态获取 bee、beego 和 Go 的版本，这样一旦用户出现错误，可以通过该命令来查看当前的版本
//
//$ bee version
//bee   :1.2.2
//beego :1.4.2
//Go    :go version go1.3.3 darwin/amd64
//2.2.7. generate 命令
//这个命令是用来自动化的生成代码的，包含了从数据库一键生成 model，还包含了 scaffold 的，通过这个命令，让大家开发代码不再慢
//
//bee generate scaffold [scaffoldname] [-fields=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
//    The generate scaffold command will do a number of things for you.
//    -fields: a list of table fields. Format: field:type, ...
//    -driver: [mysql | postgres | sqlite], the default is mysql
//    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
//    example: bee generate scaffold post -fields="title:string,body:text"
//
//bee generate model [modelname] [-fields=""]
//    generate RESTful model based on fields
//    -fields: a list of table fields. Format: field:type, ...
//
//bee generate controller [controllerfile]
//    generate RESTful controllers
//
//bee generate view [viewpath]
//    generate CRUD view in viewpath
//
//bee generate migration [migrationfile] [-fields=""]
//    generate migration file for making database schema update
//    -fields: a list of table fields. Format: field:type, ...
//
//bee generate docs
//    generate swagger doc file
//
//bee generate test [routerfile]
//    generate testcase
//
//bee generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
//    generate appcode based on an existing database
//    -tables: a list of table names separated by ',', default is empty, indicating all tables
//    -driver: [mysql | postgres | sqlite], the default is mysql
//    -conn:   the connection string used by the driver.
//             default for mysql:    root:@tcp(127.0.0.1:3306)/test
//             default for postgres: postgres://postgres:postgres@127.0.0.1:5432/postgres
//    -level:  [1 | 2 | 3], 1 = models; 2 = models,controllers; 3 = models,controllers,router
//2.2.8. migrate 命令
//这个命令是应用的数据库迁移命令，主要是用来每次应用升级，降级的SQL管理。
//
//bee migrate [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
//    run all outstanding migrations
//    -driver: [mysql | postgresql | sqlite], the default is mysql
//    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
//
//bee migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
//    rollback the last migration operation
//    -driver: [mysql | postgresql | sqlite], the default is mysql
//    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
//
//bee migrate reset [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
//    rollback all migrations
//    -driver: [mysql | postgresql | sqlite], the default is mysql
//    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
//
//bee migrate refresh [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
//    rollback all migrations and run them all again
//    -driver: [mysql | postgresql | sqlite], the default is mysql
//    -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
//2.2.9. dockerize 命令
//这个命令可以通过生成Dockerfile文件来实现docker化你的应用。
//
//例子:
//生成一个以1.6.4版本Go环境为基础镜像的Dockerfile,并暴露9000端口:
//
//$ bee dockerize -image="library/golang:1.6.4" -expose=9000
//______
//| ___ \
//| |_/ /  ___   ___
//| ___ \ / _ \ / _ \
//| |_/ /|  __/|  __/
//\____/  \___| \___| v1.6.2
//2016/12/26 22:34:54 INFO     ▶ 0001 Generating Dockerfile...
//2016/12/26 22:34:54 SUCCESS  ▶ 0002 Dockerfile generated.
//更多帮助信息可执行bee help dockerize.
//
//2.3. bee 工具配置文件
//您可能已经注意到，在 bee 工具的源码目录下有一个 bee.json 文件，这个文件是针对 bee 工具的一些行为进行配置。该功能还未完全开发完成，不过其中的一些选项已经可以使用：
//
//"version": 0：配置文件版本，用于对比是否发生不兼容的配置格式版本。
//"go_install": false：如果您的包均使用完整的导入路径（例如：github.com/user/repo/subpkg）,则可以启用该选项来进行 go install 操作，加快构建操作。
//"watch_ext": []：用于监控其它类型的文件（默认只监控后缀为 .go 的文件）。
//"dir_structure":{}：如果您的目录名与默认的 MVC 架构的不同，则可以使用该选项进行修改。
//"cmd_args": []：如果您需要在每次启动时加入启动参数，则可以使用该选项。
//"envs": []：如果您需要在每次启动时设置临时环境变量参数，则可以使用该选项。

//1. 新建项目
//2. 创建项目
//beego 的项目基本都是通过 bee 命令来创建的，所以在创建项目之前确保你已经安装了 bee 工具和 beego。如果你还没有安装，那么请查阅 beego 的安装 和 bee 工具的安装。
//
//现在一切就绪我们就可以开始创建项目了，打开终端，进入 $GOPATH/src 所在的目录：
//
//➜  src  bee new quickstart
//[INFO] Creating application...
///gopath/src/quickstart/
///gopath/src/quickstart/conf/
///gopath/src/quickstart/controllers/
///gopath/src/quickstart/models/
///gopath/src/quickstart/routers/
///gopath/src/quickstart/tests/
///gopath/src/quickstart/static/
///gopath/src/quickstart/static/js/
///gopath/src/quickstart/static/css/
///gopath/src/quickstart/static/img/
///gopath/src/quickstart/views/
///gopath/src/quickstart/conf/app.conf
///gopath/src/quickstart/controllers/default.go
///gopath/src/quickstart/views/index.tpl
///gopath/src/quickstart/routers/router.go
///gopath/src/quickstart/tests/default_test.go
///gopath/src/quickstart/main.go
//2014/11/06 18:17:09 [SUCC] New application successfully created!
//通过一个简单的命令就创建了一个 beego 项目。他的目录结构如下所示
//
//quickstart
//|-- conf
//|   `-- app.conf
//|-- controllers
//|   `-- default.go
//|-- main.go
//|-- models
//|-- routers
//|   `-- router.go
//|-- static
//|   |-- css
//|   |-- img
//|   `-- js
//|-- tests
//|   `-- default_test.go
//`-- views
//    `-- index.tpl
//从目录结构中我们也可以看出来这是一个典型的 MVC 架构的应用，main.go 是入口文件。
//
//2.1. 运行项目
//beego 项目创建之后，我们就开始运行项目，首先进入创建的项目，我们使用 bee run 来运行该项目，这样就可以做到热编译的效果：
//
//➜  src  cd quickstart
//➜  quickstart  bee run
//2014/11/06 18:18:34 [INFO] Uses 'quickstart' as 'appname'
//2014/11/06 18:18:34 [INFO] Initializing watcher...
//2014/11/06 18:18:34 [TRAC] Directory(/gopath/src/quickstart/controllers)
//2014/11/06 18:18:34 [TRAC] Directory(/gopath/src/quickstart)
//2014/11/06 18:18:34 [TRAC] Directory(/gopath/src/quickstart/routers)
//2014/11/06 18:18:34 [TRAC] Directory(/gopath/src/quickstart/tests)
//2014/11/06 18:18:34 [INFO] Start building...
//2014/11/06 18:18:35 [SUCC] Build was successful
//2014/11/06 18:18:35 [INFO] Restarting quickstart ...
//2014/11/06 18:18:35 [INFO] ./quickstart is running...
//2014/11/06 18:18:35 [app.go:96] [I] http server Running on :8080
//这样我们的应用已经在 8080 端口(beego 的默认端口)跑起来了.你是不是觉得很神奇，为什么没有 nginx 和 apache 居然可以自己干这个事情？是的，Go 其实已经做了网络层的东西，beego 只是封装了一下，所以可以做到不需要 nginx 和 apache。让我们打开浏览器看看效果吧：

//1. 路由设置
//2. 项目路由设置
//前面我们已经创建了 beego 项目，而且我们也看到它已经运行起来了，那么是如何运行起来的呢？让我们从入口文件先分析起来吧：
//
//package main
//
//import (
//    _ "quickstart/routers"
//    "github.com/astaxie/beego"
//)
//
//func main() {
//    beego.Run()
//}
//我们看到 main 函数是入口函数，但是我们知道 Go 的执行过程是如下图所示的方式：
//
//
//
//这里我们就看到了我们引入了一个包 _ "quickstart/routers",这个包只引入执行了里面的 init 函数，那么让我们看看这个里面做了什么事情：
//
//package routers
//
//import (
//    "quickstart/controllers"
//    "github.com/astaxie/beego"
//)
//
//func init() {
//    beego.Router("/", &controllers.MainController{})
//}
//路由包里面我们看到执行了路由注册 beego.Router, 这个函数的功能是映射 URL 到 controller，第一个参数是 URL (用户请求的地址)，这里我们注册的是 /，也就是我们访问的不带任何参数的 URL，第二个参数是对应的 Controller，也就是我们即将把请求分发到那个控制器来执行相应的逻辑，我们可以执行类似的方式注册如下路由：
//
//beego.Router("/user", &controllers.UserController{})
//这样用户就可以通过访问 /user 去执行 UserController 的逻辑。这就是我们所谓的路由，更多更复杂的路由规则请查询beego 的路由设置
//
//再回来看看 main 函数里面的 beego.Run， beego.Run 执行之后，我们看到的效果好像只是监听服务端口这个过程，但是它内部做了很多事情：
//
//解析配置文件
//
//beego 会自动解析在 conf 目录下面的配置文件 app.conf，通过修改配置文件相关的属性，我们可以定义：开启的端口，是否开启 session，应用名称等信息。
//
//执行用户的 hookfunc
//
//beego 会执行用户注册的 hookfunc，默认的已经存在了注册 mime，用户可以通过函数 AddAPPStartHook 注册自己的启动函数。
//
//是否开启 session
//
//会根据上面配置文件的分析之后判断是否开启 session，如果开启的话就初始化全局的 session。
//
//是否编译模板
//
//beego 会在启动的时候根据配置把 views 目录下的所有模板进行预编译，然后存在 map 里面，这样可以有效的提高模板运行的效率，无需进行多次编译。
//
//是否开启文档功能
//
//根据 EnableDocs 配置判断是否开启内置的文档路由功能
//
//是否启动管理模块
//
//beego 目前做了一个很酷的模块，应用内监控模块，会在 8088 端口做一个内部监听，我们可以通过这个端口查询到 QPS、CPU、内存、GC、goroutine、thread 等统计信息。
//
//监听服务端口
//
//这是最后一步也就是我们看到的访问 8080 看到的网页端口，内部其实调用了 ListenAndServe，充分利用了 goroutine 的优势
//
//一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。
//
//通过这个代码的分析我们了解了 beego 运行起来的过程，以及内部的一些机制。接下来让我们去剥离 Controller 如何来处理逻辑的。

//1. Controller运行机制
//2. controller 逻辑
//前面我们了解了如何把用户的请求分发到控制器，这小节我们就介绍大家如何来写控制器，首先我们还是从源码分析入手：
//
//package controllers
//
//import (
//        "github.com/astaxie/beego"
//)
//
//type MainController struct {
//        beego.Controller
//}
//
//func (this *MainController) Get() {
//        this.Data["Website"] = "beego.me"
//        this.Data["Email"] = "astaxie@gmail.com"
//        this.TplName = "index.tpl"
//}
//上面的代码显示首先我们声明了一个控制器 MainController，这个控制器里面内嵌了 beego.Controller，这就是 Go 的嵌入方式，也就是 MainController 自动拥有了所有 beego.Controller 的方法。
//
//而 beego.Controller 拥有很多方法，其中包括 Init、Prepare、Post、Get、Delete、Head 等方法。我们可以通过重写的方式来实现这些方法，而我们上面的代码就是重写了 Get 方法。
//
//我们先前介绍过 beego 是一个 RESTful 的框架，所以我们的请求默认是执行对应 req.Method 的方法。例如浏览器的是 GET 请求，那么默认就会执行 MainController 下的 Get 方法。这样我们上面的 Get 方法就会被执行到，这样就进入了我们的逻辑处理。（用户可以改变这个行为，通过注册自定义的函数名
//
//里面的代码是需要执行的逻辑，这里只是简单的输出数据，我们可以通过各种方式获取数据，然后赋值到 this.Data 中，这是一个用来存储输出数据的 map，可以赋值任意类型的值，这里我们只是简单举例输出两个字符串。
//
//最后一个就是需要去渲染的模板，this.TplName 就是需要渲染的模板，这里指定了 index.tpl，如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。
//
//用户设置了模板之后系统会自动的调用 Render 函数（这个函数是在 beego.Controller 中实现的），所以无需用户自己来调用渲染。
//
//当然也可以不使用模版，直接用 this.Ctx.WriteString 输出字符串，如：
//
//func (this *MainController) Get() {
//        this.Ctx.WriteString("hello")
//}

//1. Model逻辑
//2. model 分析
//我们知道 Web 应用中我们用的最多的就是数据库操作，而 model 层一般用来做这些操作，我们的 bee new 例子不存在 Model 的演示，但是 bee api 应用中存在 model 的应用。说的简单一点，如果您的应用足够简单，那么 Controller 可以处理一切的逻辑，如果您的逻辑里面存在着可以复用的东西，那么就抽取出来变成一个模块。因此 Model 就是逐步抽象的过程，一般我们会在 Model 里面处理一些数据读取，如下是一个日志分析应用中的代码片段：
//
//package models
//
//import (
//    "loggo/utils"
//    "path/filepath"
//    "strconv"
//    "strings"
//)
//
//var (
//    NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
//)
//
//const big = 0xFFFFFF
//
//func LogPV(urls string) bool {
//    ext := filepath.Ext(urls)
//    if ext == "" {
//        return true
//    }
//    for _, v := range NotPV {
//        if v == strings.ToLower(ext) {
//            return false
//        }
//    }
//    return true
//}
//所以如果您的应用足够简单，那么就不需要 Model 了；如果你的模块开始多了，需要复用，需要逻辑分离了，那么 Model 是必不可少的。接下来我们将分析如何编写 View 层的东西。


//1. View编写
//在前面编写 Controller 的时候，我们在 Get 里面写过这样的语句 this.TplName = "index.tpl"，设置显示的模板文件，默认支持 tpl 和 html 的后缀名，如果想设置其他后缀你可以调用 beego.AddTemplateExt 接口设置，那么模板如何来显示相应的数据呢？beego 采用了 Go 语言默认的模板引擎，所以和 Go 的模板语法一样，Go 模板的详细使用方法请参考《Go Web 编程》模板使用指南
//
//我们看看快速入门里面的代码（去掉了 css 样式）：
//
//<!DOCTYPE html>
//
//<html>
//    <head>
//        <title>Beego</title>
//        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
//    </head>
//    <body>
//        <header class="hero-unit" style="background-color:#A9F16C">
//            <div class="container">
//                <div class="row">
//                    <div class="hero-text">
//                        <h1>Welcome to Beego!</h1>
//                        <p class="description">
//                            Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
//                            <br />
//                            Official website: <a href="http://{{.Website}}">{{.Website}}</a>
//                            <br />
//                            Contact me: {{.Email}}
//                        </p>
//                    </div>
//                </div>
//            </div>
//        </header>
//    </body>
//</html>
//我们在 Controller 里面把数据赋值给了 data（map 类型），然后我们在模板中就直接通过 key 访问 .Website 和 .Email 。这样就做到了数据的输出。接下来我们讲解如何让静态文件输出。

//1. 静态文件处理
//前面我们介绍了如何输出静态页面，但是我们的网页往往包含了很多的静态文件，包括图片、JS、CSS 等，刚才创建的应用里面就创建了如下目录：
//
//├── static
//    │   ├── css
//    │   ├── img
//    │   └── js
//beego 默认注册了 static 目录为静态处理的目录，注册样式：URL 前缀和映射的目录（在/main.go文件中beego.Run()之前加入）：
//
//StaticDir["/static"] = "static"
//用户可以设置多个静态文件处理目录，例如你有多个文件下载目录 download1、download2，你可以这样映射（在 /main.go 文件中 beego.Run() 之前加入）：
//
//beego.SetStaticPath("/down1", "download1")
//beego.SetStaticPath("/down2", "download2")
//这样用户访问 URL http://localhost:8080/down1/123.txt 则会请求 download1 目录下的 123.txt 文件。

//1. 参数配置
//beego 目前支持 INI、XML、JSON、YAML 格式的配置文件解析，但是默认采用了 INI 格式解析，用户可以通过简单的配置就可以获得很大的灵活性。
//
//1.1. 默认配置解析
//beego 默认会解析当前应用下的 conf/app.conf 文件。
//
//通过这个文件你可以初始化很多 beego 的默认参数：
//
//appname = beepkg
//httpaddr = "127.0.0.1"
//httpport = 9090
//runmode ="dev"
//autorender = false
//recoverpanic = false
//viewspath = "myview"
//上面这些参数会替换 beego 默认的一些参数, beego 的参数主要有哪些呢？请参考https://godoc.org/github.com/astaxie/beego#pkg-constants 。 BConfig 就是 beego 里面的默认的配置，你也可以直接通过beego.BConfig.AppName="beepkg"这样来修改，和上面的配置效果一样，只是一个在代码里面写死了， 而配置文件就会显得更加灵活。
//
//你也可以在配置文件中配置应用需要用的一些配置信息，例如下面所示的数据库信息：
//
//mysqluser = "root"
//mysqlpass = "rootpass"
//mysqlurls = "127.0.0.1"
//mysqldb   = "beego"
//那么你就可以通过如下的方式获取设置的配置信息:
//
//beego.AppConfig.String("mysqluser")
//beego.AppConfig.String("mysqlpass")
//beego.AppConfig.String("mysqlurls")
//beego.AppConfig.String("mysqldb")
//AppConfig 的方法如下：
//
//Set(key, val string) error
//String(key string) string
//Strings(key string) []string
//Int(key string) (int, error)
//Int64(key string) (int64, error)
//Bool(key string) (bool, error)
//Float(key string) (float64, error)
//DefaultString(key string, defaultVal string) string
//DefaultStrings(key string, defaultVal []string)
//DefaultInt(key string, defaultVal int) int
//DefaultInt64(key string, defaultVal int64) int64
//DefaultBool(key string, defaultVal bool) bool
//DefaultFloat(key string, defaultVal float64) float64
//DIY(key string) (interface{}, error)
//GetSection(section string) (map[string]string, error)
//SaveConfigFile(filename string) error
//在使用 ini 类型的配置文件中, key 支持 section::key 模式.
//
//你可以用 Default* 方法返回默认值.
//
//1.1.1. 不同级别的配置
//在配置文件里面支持 section，可以有不同的 Runmode 的配置，默认优先读取 runmode 下的配置信息，例如下面的配置文件：
//
//appname = beepkg
//httpaddr = "127.0.0.1"
//httpport = 9090
//runmode ="dev"
//autorender = false
//recoverpanic = false
//viewspath = "myview"
//
//[dev]
//httpport = 8080
//[prod]
//httpport = 8088
//[test]
//httpport = 8888
//上面的配置文件就是在不同的 runmode 下解析不同的配置，例如在 dev 模式下，httpport 是 8080，在 prod 模式下是 8088，在 test 模式下是 8888。其他配置文件同理。解析的时候优先解析 runmode 下的配置，然后解析默认的配置。
//
//读取不同模式下配置参数的方法是“模式::配置参数名”，比如：beego.AppConfig.String("dev::mysqluser")。
//
//对于自定义的参数，需使用 beego.GetConfig(typ, key string, defaultVal interface{}) 来获取指定 runmode 下的配置（需 1.4.0 以上版本），typ 为参数类型，key 为参数名, defaultVal 为默认值。
//
//1.1.2. 多个配置文件
//INI 格式配置支持 include 方式，引用多个配置文件，例如下面的两个配置文件效果同上：
//
//app.conf
//
//appname = beepkg
//httpaddr = "127.0.0.1"
//httpport = 9090
//
//include "app2.conf"
//app2.conf
//
//runmode ="dev"
//autorender = false
//recoverpanic = false
//viewspath = "myview"
//
//[dev]
//httpport = 8080
//[prod]
//httpport = 8088
//[test]
//httpport = 8888
//1.1.3. 支持环境变量配置
//配置文件解析支持从环境变量中获取配置项，配置项格式：${环境变量}。例如下面的配置中优先使用环境变量中配置的 runmode 和 httpport，如果有配置环境变量 ProRunMode 则优先使用该环境变量值。如果不存在或者为空，则使用 "dev" 作为 runmode。
//
//app.conf
//
//runmode  = "${ProRunMode||dev}"
//httpport = "${ProPort||9090}"
//1.1.4. 系统默认参数
//beego 中带有很多可配置的参数，我们来一一认识一下它们，这样有利于我们在接下来的 beego 开发中可以充分的发挥他们的作用(你可以通过在 conf/app.conf 中设置对应的值，不区分大小写)：
//
//基础配置
//BConfig 保存了所有 beego 里面的系统默认参数，你可以通过 beego.BConfig 来访问和修改底下的所有配置信息.
//配置文件路径，默认是应用程序对应的目录下的 conf/app.conf，用户可以在程序代码中加载自己的配置文件 beego.LoadAppConfig("ini", "conf/app2.conf") 也可以加载多个文件，只要你调用多次就可以了，如果后面的文件和前面的 key 冲突，那么以最新加载的为最新值
//
//App 配置
//AppName
//
//应用名称，默认是 beego。通过 bee new 创建的是创建的项目名。
//
//beego.BConfig.AppName = "beego"
//
//RunMode
//
//应用的运行模式，可选值为 prod, dev 或者 test. 默认是 dev, 为开发模式，在开发模式下出错会提示友好的出错页面，如前面错误描述中所述。
//
//beego.BConfig.RunMode = "dev"
//
//RouterCaseSensitive
//
//是否路由忽略大小写匹配，默认是 true，区分大小写
//
//beego.BConfig.RouterCaseSensitive = true
//
//ServerName
//
//beego 服务器默认在请求的时候输出 server 为 beego。
//
//beego.BConfig.ServerName = "beego"
//
//RecoverPanic
//
//是否异常恢复，默认值为 true，即当应用出现异常的情况，通过 recover 恢复回来，而不会导致应用异常退出。
//
//beego.BConfig.RecoverPanic = true
//
//CopyRequestBody
//
//是否允许在 HTTP 请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件请求除外）。
//
//beego.BConfig.CopyRequestBody = false
//
//EnableGzip
//
//是否开启 gzip 支持，默认为 false 不支持 gzip，一旦开启了 gzip，那么在模板输出的内容会进行 gzip 或者 zlib 压缩，根据用户的 Accept-Encoding 来判断。
//
//beego.BConfig.EnableGzip = false
//
//Gzip允许用户自定义压缩级别、压缩长度阈值和针对请求类型压缩:
//
//压缩级别, gzipCompressLevel = 9,取值为 1~9,如果不设置为 1(最快压缩)
//
//压缩长度阈值, gzipMinLength = 256,当原始内容长度大于此阈值时才开启压缩,默认为 20B(ngnix默认长度)
//
//请求类型, includedMethods = get;post,针对哪些请求类型进行压缩,默认只针对 GET 请求压缩
//
//MaxMemory
//
//文件上传默认内存缓存大小，默认值是 1 << 26(64M)。
//
//beego.BConfig.MaxMemory = 1 << 26
//
//EnableErrorsShow
//
//是否显示系统错误信息，默认为 true。
//
//beego.BConfig.EnableErrorsShow = true
//
//EnableErrorsRender
//
//是否将错误信息进行渲染，默认值为 true，即出错会提示友好的出错页面，对于 API 类型的应用可能需要将该选项设置为 false 以阻止在 dev 模式下不必要的模板渲染信息返回。
//
//Web配置
//AutoRender
//
//是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。
//
//beego.BConfig.WebConfig.AutoRender = true
//
//EnableDocs
//
//是否开启文档内置功能，默认是 false
//
//beego.BConfig.WebConfig.EnableDocs = true
//
//FlashName
//
//Flash 数据设置时 Cookie 的名称，默认是 BEEGO_FLASH
//
//beego.BConfig.WebConfig.FlashName = "BEEGO_FLASH"
//
//FlashSeperator
//
//Flash 数据的分隔符，默认是 BEEGOFLASH
//
//beego.BConfig.WebConfig.FlashSeparator = "BEEGOFLASH"
//
//DirectoryIndex
//
//是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。
//
//beego.BConfig.WebConfig.DirectoryIndex = false
//
//StaticDir
//
//静态文件目录设置，默认是static
//
//可配置单个或多个目录:
//
//单个目录, StaticDir = download. 相当于 beego.SetStaticPath("/download","download")
//
//多个目录, StaticDir = download:down download2:down2. 相当于 beego.SetStaticPath("/download","down") 和 beego.SetStaticPath("/download2","down2")
//
//beego.BConfig.WebConfig.StaticDir
//
//StaticExtensionsToGzip
//
//允许哪些后缀名的静态文件进行 gzip 压缩，默认支持 .css 和 .js
//
//beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
//
//等价 config 文件中
//
//  StaticExtensionsToGzip = .css, .js
//TemplateLeft
//
//模板左标签，默认值是{{。
//
//beego.BConfig.WebConfig.TemplateLeft="{{"
//
//TemplateRight
//
//模板右标签，默认值是}}。
//
//beego.BConfig.WebConfig.TemplateRight="}}"
//
//ViewsPath
//
//模板路径，默认值是 views。
//
//beego.BConfig.WebConfig.ViewsPath="views"
//
//EnableXSRF
//
//是否开启 XSRF，默认为 false，不开启。
//
//beego.BConfig.WebConfig.EnableXSRF = false
//
//XSRFKEY
//
//XSRF 的 key 信息，默认值是 beegoxsrf。 EnableXSRF＝true 才有效
//
//beego.BConfig.WebConfig.XSRFKEY = "beegoxsrf"
//
//XSRFExpire
//
//XSRF 过期时间，默认值是 0，不过期。
//
//beego.BConfig.WebConfig.XSRFExpire = 0
//
//监听配置
//Graceful
//
//是否开启热升级，默认是 false，关闭热升级。
//
//beego.BConfig.Listen.Graceful=false
//
//ServerTimeOut
//
//设置 HTTP 的超时时间，默认是 0，不超时。
//
//beego.BConfig.Listen.ServerTimeOut=0
//
//ListenTCP4
//
//监听本地网络地址类型，默认是TCP6，可以通过设置为true设置为TCP4。
//
//beego.BConfig.Listen.ListenTCP4 = true
//
//EnableHTTP
//
//是否启用 HTTP 监听，默认是 true。
//
//beego.BConfig.Listen.EnableHTTP = true
//
//HTTPAddr
//
//应用监听地址，默认为空，监听所有的网卡 IP。
//
//beego.BConfig.Listen.HTTPAddr = ""
//
//HTTPPort
//
//应用监听端口，默认为 8080。
//
//beego.BConfig.Listen.HTTPPort = 8080
//
//EnableHTTPS
//
//是否启用 HTTPS，默认是 false 关闭。当需要启用时，先设置 EnableHTTPS = true，并设置 HTTPSCertFile 和 HTTPSKeyFile
//
//beego.BConfig.Listen.EnableHTTPS = false
//
//HTTPSAddr
//
//应用监听地址，默认为空，监听所有的网卡 IP。
//
//beego.BConfig.Listen.HTTPSAddr = ""
//
//HTTPSPort
//
//应用监听端口，默认为 10443
//
//beego.BConfig.Listen.HTTPSPort = 10443
//
//HTTPSCertFile
//
//开启 HTTPS 后，ssl 证书路径，默认为空。
//
//beego.BConfig.Listen.HTTPSCertFile = "conf/ssl.crt"
//
//HTTPSKeyFile
//
//开启 HTTPS 之后，SSL 证书 keyfile 的路径。
//
//beego.BConfig.Listen.HTTPSKeyFile = "conf/ssl.key"
//
//EnableAdmin
//
//是否开启进程内监控模块，默认 false 关闭。
//
//beego.BConfig.Listen.EnableAdmin = false
//
//AdminAddr
//
//监控程序监听的地址，默认值是 localhost 。
//
//beego.BConfig.Listen.AdminAddr = "localhost"
//
//AdminPort
//
//监控程序监听的地址，默认值是 8088 。
//
//beego.BConfig.Listen.AdminPort = 8088
//
//EnableFcgi
//
//是否启用 fastcgi ， 默认是 false。
//
//beego.BConfig.Listen.EnableFcgi = false
//
//EnableStdIo
//
//通过fastcgi 标准I/O，启用 fastcgi 后才生效，默认 false。
//
//beego.BConfig.Listen.EnableStdIo = false
//
//Session配置
//SessionOn
//
//session 是否开启，默认是 false。
//
//beego.BConfig.WebConfig.Session.SessionOn = false
//
//SessionProvider
//
//session 的引擎，默认是 memory，详细参见 session 模块。
//
//beego.BConfig.WebConfig.Session.SessionProvider = ""
//
//SessionName
//
//存在客户端的 cookie 名称，默认值是 beegosessionID。
//
//beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"
//
//SessionGCMaxLifetime
//
//session 过期时间，默认值是 3600 秒。
//
//beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
//
//SessionProviderConfig
//
//配置信息，根据不同的引擎设置不同的配置信息，详细的配置请看下面的引擎设置，详细参见 session 模块
//
//SessionCookieLifeTime
//
//session 默认存在客户端的 cookie 的时间，默认值是 3600 秒。
//
//beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
//
//SessionAutoSetCookie
//
//是否开启SetCookie, 默认值 true 开启。
//
//beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
//
//SessionDomain
//
//session cookie 存储域名, 默认空。
//
//beego.BConfig.WebConfig.Session.SessionDomain = ""
//
//Log配置
//log详细配置，请参见 `logs 模块`。
//AccessLogs
//
//是否输出日志到 Log，默认在 prod 模式下不会输出日志，默认为 false 不输出日志。此参数不支持配置文件配置。
//
//beego.BConfig.Log.AccessLogs = false
//
//FileLineNum
//
//是否在日志里面显示文件名和输出日志行号，默认 true。此参数不支持配置文件配置。
//
//beego.BConfig.Log.FileLineNum = true
//
//Outputs
//
//日志输出配置，参考 logs 模块，console file 等配置，此参数不支持配置文件配置。
//
//beego.BConfig.Log.Outputs = map[string]string{"console": ""}
//
//or
//
//beego.BConfig.Log.Outputs["console"] = ""

//1. 路由设置
//什么是路由设置呢？前面介绍的 MVC 结构执行时，介绍过 beego 存在三种方式的路由:固定路由、正则路由、自动路由，接下来详细的讲解如何使用这三种路由。
//
//1.1. 基础路由
//从 beego 1.2 版本开始支持了基本的 RESTful 函数式路由,应用中的大多数路由都会定义在 routers/router.go 文件中。最简单的 beego 路由由 URI 和闭包函数组成。
//
//1.1.1. 基本 GET 路由
//beego.Get("/",func(ctx *context.Context){
//     ctx.Output.Body([]byte("hello world"))
//})
//1.1.2. 基本 POST 路由
//beego.Post("/alice",func(ctx *context.Context){
//     ctx.Output.Body([]byte("bob"))
//})
//1.1.3. 注册一个可以响应任何 HTTP 的路由
//beego.Any("/foo",func(ctx *context.Context){
//     ctx.Output.Body([]byte("bar"))
//})
//1.1.4. 所有的支持的基础函数如下所示
//beego.Get(router, beego.FilterFunc)
//beego.Post(router, beego.FilterFunc)
//beego.Put(router, beego.FilterFunc)
//beego.Patch(router, beego.FilterFunc)
//beego.Head(router, beego.FilterFunc)
//beego.Options(router, beego.FilterFunc)
//beego.Delete(router, beego.FilterFunc)
//beego.Any(router, beego.FilterFunc)
//1.1.5. 支持自定义的 handler 实现
//有些时候我们已经实现了一些 rpc 的应用,但是想要集成到 beego 中,或者其他的 httpserver 应用,集成到 beego 中来.现在可以很方便的集成:
//
//s := rpc.NewServer()
//s.RegisterCodec(json.NewCodec(), "application/json")
//s.RegisterService(new(HelloService), "")
//beego.Handler("/rpc", s)
//beego.Handler(router, http.Handler) 这个函数是关键,第一个参数表示路由 URI, 第二个就是你自己实现的 http.Handler, 注册之后就会把所有 rpc 作为前缀的请求分发到 http.Handler 中进行处理.
//
//这个函数其实还有第三个参数就是是否是前缀匹配,默认是 false, 如果设置了 true, 那么就会在路由匹配的时候前缀匹配,即 /rpc/user 这样的也会匹配去运行
//
//1.1.6. 路由参数
//后面会讲到固定路由,正则路由,这些参数一样适用于上面的这些函数
//
//1.2. RESTful Controller 路由
//在介绍这三种 beego 的路由实现之前先介绍 RESTful，我们知道 RESTful 是一种目前 API 开发中广泛采用的形式，beego 默认就是支持这样的请求方法，也就是用户 Get 请求就执行 Get 方法，Post 请求就执行 Post 方法。因此默认的路由是这样 RESTful 的请求方式。
//
//1.3. 固定路由
//固定路由也就是全匹配的路由，如下所示：
//
//beego.Router("/", &controllers.MainController{})
//beego.Router("/admin", &admin.UserController{})
//beego.Router("/admin/index", &admin.ArticleController{})
//beego.Router("/admin/addpkg", &admin.AddController{})
//如上所示的路由就是我们最常用的路由方式，一个固定的路由，一个控制器，然后根据用户请求方法不同请求控制器中对应的方法，典型的 RESTful 方式。
//
//1.4. 正则路由
//为了用户更加方便的路由设置，beego 参考了 sinatra 的路由实现，支持多种方式的路由：
//
//beego.Router("/api/?:id", &controllers.RController{})
//
//默认匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"
//
//beego.Router("/api/:id", &controllers.RController{})
//
//默认匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"，但URL"/api/"匹配失败
//
//beego.Router("/api/:id([0-9]+)", &controllers.RController{})
//
//自定义正则匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"
//
//beego.Router("/user/:username([\\w]+)", &controllers.RController{})
//
//正则字符串匹配 //例如对于URL"/user/astaxie"可以匹配成功，此时变量":username"值为"astaxie"
//
//beego.Router("/download/*.*", &controllers.RController{})
//
//*匹配方式 //例如对于URL"/download/file/api.xml"可以匹配成功，此时变量":path"值为"file/api"， ":ext"值为"xml"
//
//beego.Router("/download/ceshi/*", &controllers.RController{})
//
//*全匹配方式 //例如对于URL"/download/ceshi/file/api.json"可以匹配成功，此时变量":splat"值为"file/api.json"
//
//beego.Router("/:id:int", &controllers.RController{})
//
//int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)
//
//beego.Router("/:hi:string", &controllers.RController{})
//
//string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)
//
//beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})
//
//带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123
//
//可以在 Controller 中通过如下方式获取上面的变量：
//
//this.Ctx.Input.Param(":id")
//this.Ctx.Input.Param(":username")
//this.Ctx.Input.Param(":splat")
//this.Ctx.Input.Param(":path")
//this.Ctx.Input.Param(":ext")
//1.5. 自定义方法及 RESTful 规则
//上面列举的是默认的请求方法名（请求的 method 和函数名一致，例如 GET 请求执行 Get 函数，POST 请求执行 Post 函数），如果用户期望自定义函数名，那么可以使用如下方式：
//
//beego.Router("/",&IndexController{},"*:Index")
//使用第三个参数，第三个参数就是用来设置对应 method 到函数名，定义如下
//
//*表示任意的 method 都执行该函数
//使用 httpmethod:funcname 格式来展示
//多个不同的格式使用 ; 分割
//多个 method 对应同一个 funcname，method 之间通过 , 来分割
//以下是一个 RESTful 的设计示例：
//
//beego.Router("/api/list",&RestController{},"*:ListFood")
//beego.Router("/api/create",&RestController{},"post:CreateFood")
//beego.Router("/api/update",&RestController{},"put:UpdateFood")
//beego.Router("/api/delete",&RestController{},"delete:DeleteFood")
//以下是多个 HTTP Method 指向同一个函数的示例：
//
//beego.Router("/api",&RestController{},"get,post:ApiFunc")
//以下是不同的 method 对应不同的函数，通过 ; 进行分割的示例：
//
//beego.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")
//可用的 HTTP Method：
//
//包含以下所有的函数
//
//* get: GET 请求
//* post: POST 请求
//* put: PUT 请求
//* delete: DELETE 请求
//* patch: PATCH 请求
//* options: OPTIONS 请求
//* head: HEAD 请求
//如果同时存在 * 和对应的 HTTP Method，那么优先执行 HTTP Method 的方法，例如同时注册了如下所示的路由：
//
//beego.Router("/simple",&SimpleController{},"`*:AllFunc;post:PostFunc`")
//那么执行 POST 请求的时候，执行 PostFunc 而不执行 AllFunc。
//
//自定义函数的路由默认不支持 RESTful 的方法，也就是如果你设置了 beego.Router("/api",&RestController{},"post:ApiFunc") 这样的路由，如果请求的方法是 POST，那么不会默认去执行 Post 函数。
//
//1.6. 自动匹配
//用户首先需要把需要路由的控制器注册到自动路由中：
//
//beego.AutoRouter(&controllers.ObjectController{})
//那么 beego 就会通过反射获取该结构体中所有的实现方法，你就可以通过如下的方式访问到对应的方法中：
//
///object/login   调用 ObjectController 中的 Login 方法
///object/logout  调用 ObjectController 中的 Logout 方法
//除了前缀两个 /:controller/:method 的匹配之外，剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中：
//
///object/blog/2013/09/12  调用 ObjectController 中的 Blog 方法，参数如下：map[0:2013 1:09 2:12]
//方法名在内部是保存了用户设置的，例如 Login，url 匹配的时候都会转化为小写，所以，/object/LOGIN 这样的 url 也一样可以路由到用户定义的 Login 方法中。
//
//现在已经可以通过自动识别出来下面类似的所有 url，都会把请求分发到 controller 的 simple 方法：
//
///controller/simple
///controller/simple.html
///controller/simple.json
///controller/simple.xml
//可以通过 this.Ctx.Input.Param(":ext") 获取后缀名。
//
//1.7. 注解路由
//从 beego 1.3 版本开始支持了注解路由，用户无需在 router 中注册路由，只需要 Include 相应地 controller，然后在 controller 的 method 方法上面写上 router 注释（// @router）就可以了，详细的使用请看下面的例子：
//
//// CMS API
//type CMSController struct {
//    beego.Controller
//}
//
//func (c *CMSController) URLMapping() {
//    c.Mapping("StaticBlock", c.StaticBlock)
//    c.Mapping("AllBlock", c.AllBlock)
//}
//
//
//// @router /staticblock/:key [get]
//func (this *CMSController) StaticBlock() {
//
//}
//
//// @router /all/:key [get]
//func (this *CMSController) AllBlock() {
//
//}
//可以在 router.go 中通过如下方式注册路由：
//
//beego.Include(&CMSController{})
//beego 自动会进行源码分析，注意只会在 dev 模式下进行生成，生成的路由放在 "/routers/commentsRouter.go" 文件中。
//
//这样上面的路由就支持了如下的路由：
//
//GET /staticblock/:key
//GET /all/:key
//其实效果和自己通过 Router 函数注册是一样的：
//
//beego.Router("/staticblock/:key", &CMSController{}, "get:StaticBlock")
//beego.Router("/all/:key", &CMSController{}, "get:AllBlock")
//同时大家注意到新版本里面增加了 URLMapping 这个函数，这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过 interface 来进行执行函数，性能上面会提升很多。
//
//1.8. namespace
////初始化 namespace
//ns :=
//beego.NewNamespace("/v1",
//    beego.NSCond(func(ctx *context.Context) bool {
//        if ctx.Input.Domain() == "api.beego.me" {
//            return true
//        }
//        return false
//    }),
//    beego.NSBefore(auth),
//    beego.NSGet("/notallowed", func(ctx *context.Context) {
//        ctx.Output.Body([]byte("notAllowed"))
//    }),
//    beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
//    beego.NSRouter("/changepassword", &UserController{}),
//    beego.NSNamespace("/shop",
//        beego.NSBefore(sentry),
//        beego.NSGet("/:id", func(ctx *context.Context) {
//            ctx.Output.Body([]byte("notAllowed"))
//        }),
//    ),
//    beego.NSNamespace("/cms",
//        beego.NSInclude(
//            &controllers.MainController{},
//            &controllers.CMSController{},
//            &controllers.BlockController{},
//        ),
//    ),
//)
////注册 namespace
//beego.AddNamespace(ns)
//上面这个代码支持了如下这样的请求 URL
//
//GET /v1/notallowed
//GET /v1/version
//GET /v1/changepassword
//POST /v1/changepassword
//GET /v1/shop/123
//GET /v1/cms/ 对应 MainController、CMSController、BlockController 中得注解路由
//而且还支持前置过滤,条件判断,无限嵌套 namespace
//
//namespace 的接口如下:
//
//NewNamespace(prefix string, funcs ...interface{})
//
//初始化 namespace 对象,下面这些函数都是 namespace 对象的方法,但是强烈推荐使用 NS 开头的相应函数注册，因为这样更容易通过 gofmt 工具看的更清楚路由的级别关系
//
//NSCond(cond namespaceCond)
//
//支持满足条件的就执行该 namespace, 不满足就不执行
//
//NSBefore(filiterList ...FilterFunc)
//
//NSAfter(filiterList ...FilterFunc)
//
//上面分别对应 beforeRouter 和 FinishRouter 两个过滤器，可以同时注册多个过滤器
//
//NSInclude(cList ...ControllerInterface)
//
//NSRouter(rootpath string, c ControllerInterface, mappingMethods ...string)
//NSGet(rootpath string, f FilterFunc)
//NSPost(rootpath string, f FilterFunc)
//NSDelete(rootpath string, f FilterFunc)
//NSPut(rootpath string, f FilterFunc)
//NSHead(rootpath string, f FilterFunc)
//NSOptions(rootpath string, f FilterFunc)
//NSPatch(rootpath string, f FilterFunc)
//NSAny(rootpath string, f FilterFunc)
//NSHandler(rootpath string, h http.Handler)
//NSAutoRouter(c ControllerInterface)
//NSAutoPrefix(prefix string, c ControllerInterface)
//
//上面这些都是设置路由的函数,详细的使用和上面 beego 的对应函数是一样的
//
//NSNamespace(prefix string, params ...innnerNamespace)
//
//嵌套其他 namespace
//
//  ns :=
//    beego.NewNamespace("/v1",
//      beego.NSNamespace("/shop",
//          beego.NSGet("/:id", func(ctx *context.Context) {
//              ctx.Output.Body([]byte("shopinfo"))
//          }),
//      ),
//      beego.NSNamespace("/order",
//          beego.NSGet("/:id", func(ctx *context.Context) {
//              ctx.Output.Body([]byte("orderinfo"))
//          }),
//      ),
//      beego.NSNamespace("/crm",
//          beego.NSGet("/:id", func(ctx *context.Context) {
//              ctx.Output.Body([]byte("crminfo"))
//          }),
//      ),
//  )
//下面这些函数都是属于 *Namespace 对象的方法：不建议直接使用，当然效果和上面的 NS 开头的函数是一样的，只是上面的方式更优雅，写出来的代码更容易看得懂
//
//Cond(cond namespaceCond)
//
//支持满足条件的就执行该 namespace, 不满足就不执行,例如你可以根据域名来控制 namespace
//
//Filter(action string, filter FilterFunc)
//
//action 表示你需要执行的位置, before 和 after 分别表示执行逻辑之前和执行逻辑之后的 filter
//
//Router(rootpath string, c ControllerInterface, mappingMethods ...string)
//
//AutoRouter(c ControllerInterface)
//AutoPrefix(prefix string, c ControllerInterface)
//Get(rootpath string, f FilterFunc)
//Post(rootpath string, f FilterFunc)
//Delete(rootpath string, f FilterFunc)
//Put(rootpath string, f FilterFunc)
//Head(rootpath string, f FilterFunc)
//Options(rootpath string, f FilterFunc)
//Patch(rootpath string, f FilterFunc)
//Any(rootpath string, f FilterFunc)
//Handler(rootpath string, h http.Handler)
//
//上面这些都是设置路由的函数,详细的使用和上面 beego 的对应函数是一样的
//
//Namespace(ns ...*Namespace)
//
//更多的例子代码：
//
////APIS
//ns :=
//    beego.NewNamespace("/api",
//        //此处正式版时改为验证加密请求
//        beego.NSCond(func(ctx *context.Context) bool {
//            if ua := ctx.Input.Request.UserAgent(); ua != "" {
//                return true
//            }
//            return false
//        }),
//        beego.NSNamespace("/ios",
//            //CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
//            beego.NSNamespace("/create",
//                // /api/ios/create/node/
//                beego.NSRouter("/node", &apis.CreateNodeHandler{}),
//                // /api/ios/create/topic/
//                beego.NSRouter("/topic", &apis.CreateTopicHandler{}),
//            ),
//            beego.NSNamespace("/read",
//                beego.NSRouter("/node", &apis.ReadNodeHandler{}),
//                beego.NSRouter("/topic", &apis.ReadTopicHandler{}),
//            ),
//            beego.NSNamespace("/update",
//                beego.NSRouter("/node", &apis.UpdateNodeHandler{}),
//                beego.NSRouter("/topic", &apis.UpdateTopicHandler{}),
//            ),
//            beego.NSNamespace("/delete",
//                beego.NSRouter("/node", &apis.DeleteNodeHandler{}),
//                beego.NSRouter("/topic", &apis.DeleteTopicHandler{}),
//            )
//        ),
//    )
//
//beego.AddNamespace(ns)

//1. 控制器函数
//提示：在 v1.6 中，此文档所涉及的 API 有重大变更，this.ServeJson() 更改为 this.ServeJSON()，this.TplNames 更改为 this.TplName。
//
//基于 beego 的 Controller 设计，只需要匿名组合 beego.Controller 就可以了，如下所示：
//
//type xxxController struct {
//    beego.Controller
//}
//beego.Controller 实现了接口 beego.ControllerInterface，beego.ControllerInterface 定义了如下函数：
//
//Init(ct *context.Context, childName string, app interface{})
//
//这个函数主要初始化了 Context、相应的 Controller 名称，模板名，初始化模板参数的容器 Data，app 即为当前执行的 Controller 的 reflecttype，这个 app 可以用来执行子类的方法。
//
//Prepare()
//
//这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
//
//Get()
//
//如果用户请求的 HTTP Method 是 GET，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Get 请求。
//
//Post()
//
//如果用户请求的 HTTP Method 是 POST，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Post 请求。
//
//Delete()
//
//如果用户请求的 HTTP Method 是 DELETE，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Delete 请求。
//
//Put()
//
//如果用户请求的 HTTP Method 是 PUT，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Put 请求.
//
//Head()
//
//如果用户请求的 HTTP Method 是 HEAD，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Head 请求。
//
//Patch()
//
//如果用户请求的 HTTP Method 是 PATCH，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Patch 请求.
//
//Options()
//
//如果用户请求的HTTP Method是OPTIONS，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Options 请求。
//
//Finish()
//
//这个函数是在执行完相应的 HTTP Method 方法之后执行的，默认是空，用户可以在子 struct 中重写这个函数，执行例如数据库关闭，清理数据之类的工作。
//
//Render() error
//
//这个函数主要用来实现渲染模板，如果 beego.AutoRender 为 true 的情况下才会执行。
//
//所以通过子 struct 的方法重写，用户就可以实现自己的逻辑，接下来我们看一个实际的例子：
//
//type AddController struct {
//    beego.Controller
//}
//
//func (this *AddController) Prepare() {
//
//}
//
//func (this *AddController) Get() {
//    this.Data["content"] = "value"
//    this.Layout = "admin/layout.html"
//    this.TplName = "admin/add.tpl"
//}
//
//func (this *AddController) Post() {
//    pkgname := this.GetString("pkgname")
//    content := this.GetString("content")
//    pk := models.GetCruPkg(pkgname)
//    if pk.Id == 0 {
//        var pp models.PkgEntity
//        pp.Pid = 0
//        pp.Pathname = pkgname
//        pp.Intro = pkgname
//        models.InsertPkg(pp)
//        pk = models.GetCruPkg(pkgname)
//    }
//    var at models.Article
//    at.Pkgid = pk.Id
//    at.Content = content
//    models.InsertArticle(at)
//    this.Ctx.Redirect(302, "/admin/index")
//}
//从上面的例子可以看出来，通过重写方法可以实现对应 method 的逻辑，实现 RESTful 结构的逻辑处理。
//
//下面我们再来看一种比较流行的架构，首先实现一个自己的基类 baseController，实现一些初始化的方法，然后其他所有的逻辑继承自该基类：
//
//type NestPreparer interface {
//        NestPrepare()
//}
//
//// baseRouter implemented global settings for all other routers.
//type baseController struct {
//        beego.Controller
//        i18n.Locale
//        user    models.User
//        isLogin bool
//}
//// Prepare implemented Prepare method for baseRouter.
//func (this *baseController) Prepare() {
//
//        // page start time
//        this.Data["PageStartTime"] = time.Now()
//
//        // Setting properties.
//        this.Data["AppDescription"] = utils.AppDescription
//        this.Data["AppKeywords"] = utils.AppKeywords
//        this.Data["AppName"] = utils.AppName
//        this.Data["AppVer"] = utils.AppVer
//        this.Data["AppUrl"] = utils.AppUrl
//        this.Data["AppLogo"] = utils.AppLogo
//        this.Data["AvatarURL"] = utils.AvatarURL
//        this.Data["IsProMode"] = utils.IsProMode
//
//        if app, ok := this.AppController.(NestPreparer); ok {
//                app.NestPrepare()
//        }
//}
//上面定义了基类，大概是初始化了一些变量，最后有一个 Init 函数中那个 app 的应用，判断当前运行的 Controller 是否是 NestPreparer 实现，如果是的话调用子类的方法，下面我们来看一下 NestPreparer 的实现：
//
//type BaseAdminRouter struct {
//    baseController
//}
//
//func (this *BaseAdminRouter) NestPrepare() {
//    if this.CheckActiveRedirect() {
//            return
//    }
//
//    // if user isn't admin, then logout user
//    if !this.user.IsAdmin {
//            models.LogoutUser(&this.Controller)
//
//            // write flash message
//            this.FlashWrite("NotPermit", "true")
//
//            this.Redirect("/login", 302)
//            return
//    }
//
//    // current in admin page
//    this.Data["IsAdmin"] = true
//
//    if app, ok := this.AppController.(ModelPreparer); ok {
//            app.ModelPrepare()
//            return
//    }
//}
//
//func (this *BaseAdminRouter) Get(){
//    this.TplName = "Get.tpl"
//}
//
//func (this *BaseAdminRouter) Post(){
//    this.TplName = "Post.tpl"
//}
//这样我们的执行器执行的逻辑是这样的，首先执行 Prepare，这个就是 Go 语言中 struct 中寻找方法的顺序，依次往父类寻找。执行 BaseAdminRouter 时，查找他是否有 Prepare 方法，没有就寻找 baseController，找到了，那么就执行逻辑，然后在 baseController 里面的 this.AppController 即为当前执行的控制器 BaseAdminRouter，因为会执行 BaseAdminRouter.NestPrepare 方法。然后开始执行相应的 Get 方法或者 Post 方法。
//
//1.1. 提前终止运行
//我们应用中经常会遇到这样的情况，在 Prepare 阶段进行判断，如果用户认证不通过，就输出一段信息，然后直接中止进程，之后的 Post、Get 之类的不再执行，那么如何终止呢？可以使用 StopRun 来终止执行逻辑，可以在任意的地方执行。
//
//type RController struct {
//    beego.Controller
//}
//
//func (this *RController) Prepare() {
//    this.Data["json"] = map[string]interface{}{"name": "astaxie"}
//    this.ServeJSON()
//    this.StopRun()
//}
//调用 StopRun 之后，如果你还定义了 Finish 函数就不会再执行，如果需要释放资源，那么请自己在调用 StopRun 之前手工调用 Finish 函数。
//
//1.2. 在表单中使用 PUT 方法
//首先要说明, 在 XHTML 1.x 标准中, 表单只支持 GET 或者 POST 方法. 虽然说根据标准, 你不应该将表单提交到 PUT 方法, 但是如果你真想的话, 也很容易, 通常可以这么做:
//
//首先表单本身还是使用 POST 方法提交, 但是可以在表单中添加一个隐藏字段:
//
//<form method="post" ...>
//  <input type="hidden" name="_method" value="put" />

//接着在 Beego 中添加一个过滤器来判断是否将请求当做 PUT 来解析:
//
//var FilterMethod = func(ctx *context.Context) {
//    if ctx.BeegoInput.Query("_method")!="" && ctx.BeegoInput.IsPost(){
//          ctx.Request.Method = ctx.BeegoInput.Query("_method")
//    }
//}
//
//beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

//1. XSRF过滤
//2. 跨站请求伪造
//跨站请求伪造(Cross-site request forgery)， 简称为 XSRF，是 Web 应用中常见的一个安全问题。前面的链接也详细讲述了 XSRF 攻击的实现方式。
//
//当前防范 XSRF 的一种通用的方法，是对每一个用户都记录一个无法预知的 cookie 数据，然后要求所有提交的请求（POST/PUT/DELETE）中都必须带有这个 cookie 数据。如果此数据不匹配 ，那么这个请求就可能是被伪造的。
//
//beego 有内建的 XSRF 的防范机制，要使用此机制，你需要在应用配置文件中加上 enablexsrf 设定：
//
//enablexsrf = true
//xsrfkey = 61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o
//xsrfexpire = 3600
//或者直接在 main 入口处这样设置：
//
//beego.EnableXSRF = true
//beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
//beego.XSRFExpire = 3600  //过期时间，默认1小时
//如果开启了 XSRF，那么 beego 的 Web 应用将对所有用户设置一个 _xsrf 的 cookie 值（默认过期 1 小时），如果 POST PUT DELET 请求中没有这个 cookie 值，那么这个请求会被直接拒绝。如果你开启了这个机制，那么在所有被提交的表单中，你都需要加上一个域来提供这个值。你可以通过在模板中使用 专门的函数 XSRFFormHTML() 来做到这一点：
//
//过期时间上面我们设置了全局的过期时间 beego.XSRFExpire，但是有些时候我们也可以在控制器中修改这个过期时间，专门针对某一类处理逻辑：
//
//func (this *HomeController) Get(){
//    this.XSRFExpire = 7200
//    this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
//}
//2.1.1. 在表单中使用
//在 Controller 中这样设置数据：
//
//func (this *HomeController) Get(){
//    this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
//}
//然后在模板中这样设置：
//
//<form action="/new_message" method="post">
//  {{ .xsrfdata }}
//  <input type="text" name="message"/>
//  <input type="submit" value="Post"/>
//</form>
//2.1.2. 在 JavaScript 中使用
//如果你提交的是 AJAX 的 POST 请求，你还是需要在每一个请求中通过脚本添加上 _xsrf 这个值。下面是在 AJAX 的 POST 请求，使用了 jQuery 函数来为所有请求都添加 _xsrf 值：
//
//jQuery cookie插件：https://github.com/carhartl/jquery-cookie base64 插件：http://phpjs.org/functions/base64_decode/
//
//jQuery.postJSON = function(url, args, callback) {
//   var xsrf, xsrflist;
//   xsrf = $.cookie("_xsrf");
//   xsrflist = xsrf.split("|");
//   args._xsrf = base64_decode(xsrflist[0]);
//    $.ajax({url: url, data: $.param(args), dataType: "text", type: "POST",
//        success: function(response) {
//        callback(eval("(" + response + ")"));
//    }});
//};
//扩展 jQuery
//通过扩展 ajax 给每个请求加入 xsrf 的 header
//
//需要你在 html 里保存一个 _xsrf 值
//
//func (this *HomeController) Get(){
//    this.Data["xsrf_token"] = this.XSRFToken()
//}
//放在你的 head 中
//
//<head>
//    <meta name="_xsrf" content="{{.xsrf_token}}" />
//</head>
//扩展 ajax 方法，将 _xsrf 值加入 header，扩展后支持 jquery post/get 等内部使用了 ajax 的方法
//
//var ajax = $.ajax;
//$.extend({
//    ajax: function(url, options) {
//        if (typeof url === 'object') {
//            options = url;
//            url = undefined;
//        }
//        options = options || {};
//        url = options.url;
//        var xsrftoken = $('meta[name=_xsrf]').attr('content');
//        var headers = options.headers || {};
//        var domain = document.domain.replace(/\./ig, '\\.');
//        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
//            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
//        }
//        options.headers = headers;
//        return ajax(url, options);
//    }
//});
//对于 PUT 和 DELETE 请求（以及不使用将 form 内容作为参数的 POST 请求）来说，你也可以在 HTTP 头中以 X-XSRFToken 这个参数传递 XSRF token。
//
//如果你需要针对每一个请求处理器定制 XSRF 行为，你可以重写 Controller 的 CheckXSRFCookie 方法。例如你需要使用一个不支持 cookie 的 API， 你可以通过将 CheckXSRFCookie() 函数设空来禁用 XSRF 保护机制。然而如果 你需要同时支持 cookie 和非 cookie 认证方式，那么只要当前请求是通过 cookie 进行认证的，你就应该对其使用 XSRF 保护机制，这一点至关重要。

//2.2. 支持controller 级别的屏蔽
//XSRF 之前是全局设置的一个参数,如果设置了那么所有的 API 请求都会进行验证,但是有些时候API 逻辑是不需要进行验证的,因此现在支持在controller 级别设置屏蔽:
//
//type AdminController struct{
//    beego.Controller
//}
//
//func (a *AdminController) Prepare() {
//    a.EnableXSRF = false
//}

//1. 请求数据处理
//2. 获取参数
//我们经常需要获取用户传递的数据，包括 Get、POST 等方式的请求，beego 里面会自动解析这些数据，你可以通过如下方式获取数据：
//
//GetString(key string) string
//GetStrings(key string) []string
//GetInt(key string) (int64, error)
//GetBool(key string) (bool, error)
//GetFloat(key string) (float64, error)
//使用例子如下：
//
//func (this *MainController) Post() {
//    jsoninfo := this.GetString("jsoninfo")
//    if jsoninfo == "" {
//        this.Ctx.WriteString("jsoninfo is empty")
//        return
//    }
//}
//如果你需要的数据可能是其他类型的，例如是 int 类型而不是 int64，那么你需要这样处理：
//
//func (this *MainController) Post() {
//    id := this.Input().Get("id")
//    intid, err := strconv.Atoi(id)
//}
//更多其他的 request 的信息，用户可以通过 this.Ctx.Request 获取信息，关于该对象的属性和方法参考手册 Request。
//
//2.1. 直接解析到 struct
//如果要把表单里的内容赋值到一个 struct 里，除了用上面的方法一个一个获取再赋值外，beego 提供了通过另外一个更便捷的方式，就是通过 struct 的字段名或 tag 与表单字段对应直接解析到 struct。
//
//定义 struct：
//
//type user struct {
//    Id    int         `form:"-"`
//    Name  interface{} `form:"username"`
//    Age   int         `form:"age"`
//    Email string
//}
//表单：
//
//<form id="user">
//    名字：<input name="username" type="text" />
//    年龄：<input name="age" type="text" />
//    邮箱：<input name="Email" type="text" />
//    <input type="submit" value="提交" />
//</form>
//Controller 里解析：
//
//func (this *MainController) Post() {
//    u := user{}
//    if err := this.ParseForm(&u); err != nil {
//        //handle error
//    }
//}
//注意：
//
//StructTag form 的定义和renderform方法共用一个标签
//定义 struct 时，字段名后如果有 form 这个 tag，则会以把 form 表单里的 name 和 tag 的名称一样的字段赋值给这个字段，否则就会把 form 表单里与字段名一样的表单内容赋值给这个字段。如上面例子中，会把表单中的 username 和 age 分别赋值给 user 里的 Name 和 Age 字段，而 Email 里的内容则会赋给 Email 这个字段。
//调用 Controller ParseForm 这个方法的时候，传入的参数必须为一个 struct 的指针，否则对 struct 的赋值不会成功并返回 xx must be a struct pointer 的错误。
//如果要忽略一个字段，有两种办法，一是：字段名小写开头，二是：form 标签的值设置为 -
//2.2. 获取 Request Body 里的内容
//在 API 的开发中，我们经常会用到 JSON 或 XML 来作为数据交互的格式，如何在 beego 中获取 Request Body 里的 JSON 或 XML 的数据呢？
//
//在配置文件里设置 copyrequestbody = true
//在 Controller 中
//func (this *ObjectController) Post() {
//    var ob models.Object
//    var err error
//    if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
//        objectid := models.AddOne(ob)
//        this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
//    } else {
//        this.Data["json"] = err.Error()
//    }
//    this.ServeJSON()
//}
//2.3. 文件上传
//在 beego 中你可以很容易的处理文件上传，就是别忘记在你的 form 表单中增加这个属性 enctype="multipart/form-data"，否则你的浏览器不会传输你的上传文件。
//
//文件上传之后一般是放在系统的内存里面，如果文件的 size 大于设置的缓存内存大小，那么就放在临时文件中，默认的缓存内存是 64M，你可以通过如下来调整这个缓存内存大小:
//
//beego.MaxMemory = 1<<22
//或者在配置文件中通过如下设置：
//
//maxmemory = 1<<22
//Beego 提供了两个很方便的方法来处理文件上传：
//
//GetFile(key string) (multipart.File, *multipart.FileHeader, error)
//
//该方法主要用于用户读取表单中的文件名 the_file，然后返回相应的信息，用户根据这些变量来处理文件上传：过滤、保存文件等。
//
//SaveToFile(fromfile, tofile string) error
//
//该方法是在 GetFile 的基础上实现了快速保存的功能 fromfile 是提交时候的 html 表单中的 name
//
//<form enctype="multipart/form-data" method="post">
//    <input type="file" name="uploadname" />
//    <input type="submit">
//</form>
//保存的代码例子如下：
//
//func (c *FormController) Post() {
//    f, h, err := c.GetFile("uploadname")
//    if err != nil {
//        log.Fatal("getfile err ", err)
//    }
//    defer f.Close()
//    c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
//
//}
//2.4. 数据绑定
//支持从用户请求中直接数据 bind 到指定的对象,例如请求地址如下
//
//?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
//var id int
//this.Ctx.Input.Bind(&id, "id")  //id ==123
//
//var isok bool
//this.Ctx.Input.Bind(&isok, "isok")  //isok ==true
//
//var ft float64
//this.Ctx.Input.Bind(&ft, "ft")  //ft ==1.2
//
//ol := make([]int, 0, 2)
//this.Ctx.Input.Bind(&ol, "ol")  //ol ==[1 2]
//
//ul := make([]string, 0, 2)
//this.Ctx.Input.Bind(&ul, "ul")  //ul ==[str array]
//
//user struct{Name}
//this.Ctx.Input.Bind(&user, "user")  //user =={Name:"astaxie"}

//1. Session控制
//beego 内置了 session 模块，目前 session 模块支持的后端引擎包括 memory、cookie、file、mysql、redis、couchbase、memcache、postgres，用户也可以根据相应的 interface 实现自己的引擎。
//
//beego 中使用 session 相当方便，只要在 main 入口函数中设置如下：
//
//beego.BConfig.WebConfig.Session.SessionOn = true
//或者通过配置文件配置如下：
//
//sessionon = true
//通过这种方式就可以开启 session，如何使用 session，请看下面的例子：
//
//func (this *MainController) Get() {
//    v := this.GetSession("asta")
//    if v == nil {
//        this.SetSession("asta", int(1))
//        this.Data["num"] = 0
//    } else {
//        this.SetSession("asta", v.(int)+1)
//        this.Data["num"] = v.(int)
//    }
//    this.TplName = "index.tpl"
//}
//session 有几个方便的方法：
//
//SetSession(name string, value interface{})
//GetSession(name string) interface{}
//DelSession(name string)
//SessionRegenerateID()
//DestroySession()
//session 操作主要有设置 session、获取 session、删除 session。
//
//当然你可以通过下面的方式自己控制这些逻辑：
//
//sess:=this.StartSession()
//defer sess.SessionRelease()
//sess 对象具有如下方法：
//
//sess.Set()
//sess.Get()
//sess.Delete()
//sess.SessionID()
//sess.Flush()
//但是我还是建议大家采用 SetSession、GetSession、DelSession 三个方法来操作，避免自己在操作的过程中资源没释放的问题。
//
//关于 Session 模块使用中的一些参数设置：
//
//beego.BConfig.WebConfig.Session.SessionOn
//
//设置是否开启 Session，默认是 false，配置文件对应的参数名：sessionon。
//
//beego.BConfig.WebConfig.Session.SessionProvider
//
//设置 Session 的引擎，默认是 memory，目前支持还有 file、mysql、redis 等，配置文件对应的参数名：sessionprovider。
//
//beego.BConfig.WebConfig.Session.SessionName
//
//设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID，配置文件对应的参数名是：sessionname。
//
//beego.BConfig.WebConfig.Session.SessionGCMaxLifetime
//
//设置 Session 过期的时间，默认值是 3600 秒，配置文件对应的参数：sessiongcmaxlifetime。
//
//beego.BConfig.WebConfig.Session.SessionProviderConfig
//
//设置对应 file、mysql、redis 引擎的保存路径或者链接地址，默认值是空，配置文件对应的参数：sessionproviderconfig。
//
//beego.BConfig.WebConfig.Session.SessionHashFunc
//
//默认值为 sha1，采用 sha1 加密算法生产 sessionid
//
//beego.BConfig.WebConfig.Session.SessionHashKey
//
//默认的 key 是 beegoserversessionkey，建议用户使用的时候修改该参数
//
//beego.BConfig.WebConfig.Session.SessionCookieLifeTime
//
//设置 cookie 的过期时间，cookie 是用来存储保存在客户端的数据。
//
//从 beego1.1.3 版本开始移除了第三方依赖库,也就是如果你想使用 mysql、redis、couchbase、memcache、postgres 这些引擎,那么你首先需要安装
//
//go get -u github.com/astaxie/beego/session/mysql
//然后在你的 main 函数中引入该库, 和数据库的驱动引入是一样的:
//
//import _ "github.com/astaxie/beego/session/mysql"
//当 SessionProvider 为 file SessionProviderConfig 是指保存文件的目录，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider="file"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
//当 SessionProvider 为 mysql 时，SessionProviderConfig 是链接地址，采用 go-sql-driver，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "username:password@protocol(address)/dbname?param=value"
//
//需要特别注意的是，在使用 mysql 存储 session 信息的时候，需要事先在 mysql 创建表，建表语句如下
//    CREATE TABLE `session` (
//        `session_key` char(64) NOT NULL,
//        `session_data` blob,
//        `session_expiry` int(11) unsigned NOT NULL,
//        PRIMARY KEY (`session_key`)
//    ) ENGINE=MyISAM DEFAULT CHARSET=utf8;
//当 SessionProvider 为 redis 时，SessionProviderConfig 是 redis 的链接地址，采用了 redigo，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider = "redis"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
//当 SessionProvider 为 memcache 时，SessionProviderConfig 是 memcache 的链接地址，采用了 memcache，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider = "memcache"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:7080"
//当 SessionProvider 为 postgres 时，SessionProviderConfig 是 postgres 的链接地址，采用了 postgres，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider = "postgresql"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
//当 SessionProvider 为 couchbase 时，SessionProviderConfig 是 couchbase 的链接地址，采用了 couchbase，如下所示：
//
//beego.BConfig.WebConfig.Session.SessionProvider = "couchbase"
//beego.BConfig.WebConfig.Session.SessionProviderConfig = "http://bucketname:bucketpass@myserver:8091"
//1.1. 特别注意点
//因为 session 内部采用了 gob 来注册存储的对象，例如 struct，所以如果你采用了非 memory 的引擎，请自己在 main.go 的 init 里面注册需要保存的这些结构体，不然会引起应用重启之后出现无法解析的错误

//1. 过滤器
//beego 支持自定义过滤中间件，例如安全验证，强制跳转等。
//
//过滤器函数如下所示：
//
//beego.InsertFilter(pattern string, position int, filter FilterFunc, params ...bool)
//InsertFilter 函数的三个必填参数，一个可选参数
//
//pattern 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用 *
//position 执行 Filter 的地方，五个固定参数如下，分别表示不同的执行过程
//BeforeStatic 静态地址之前
//BeforeRouter 寻找路由之前
//BeforeExec 找到路由之后，开始执行相应的 Controller 之前
//AfterExec 执行完 Controller 逻辑之后执行的过滤器
//FinishRouter 执行完逻辑之后执行的过滤器
//filter filter 函数 type FilterFunc func(*context.Context)
//params
//设置 returnOnOutput 的值(默认 true), 如果在进行到此过滤之前已经有输出，是否不再继续执行此过滤器,默认设置为如果前面已有输出(参数为true)，则不再执行此过滤器
//是否重置 filters 的参数，默认是 false，因为在 filters 的 pattern 和本身的路由的 pattern 冲突的时候，可以把 filters 的参数重置，这样可以保证在后续的逻辑中获取到正确的参数，例如设置了 /api/* 的 filter，同时又设置了 /api/docs/* 的 router，那么在访问 /api/docs/swagger/abc.js 的时候，在执行 filters 的时候设置 :splat 参数为 docs/swagger/abc.js，但是如果不清楚 filter 的这个路由参数，就会在执行路由逻辑的时候保持 docs/swagger/abc.js，如果设置了 true，就会重置 :splat 参数.
//AddFilter 从beego1.3 版本开始已经废除
//
//如下例子所示，验证用户是否已经登录，应用于全部的请求：
//
//var FilterUser = func(ctx *context.Context) {
//    _, ok := ctx.Input.Session("uid").(int)
//    if !ok && ctx.Request.RequestURI != "/login" {
//        ctx.Redirect(302, "/login")
//    }
//}
//
//beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
//这里需要特别注意使用 session 的 Filter 必须在 BeforeStatic 之后才能获取，因为 session 没有在这之前初始化。
//
//还可以通过正则路由进行过滤，如果匹配参数就执行：
//
//var FilterUser = func(ctx *context.Context) {
//    _, ok := ctx.Input.Session("uid").(int)
//    if !ok {
//        ctx.Redirect(302, "/login")
//    }
//}
//beego.InsertFilter("/user/:id([0-9]+)",beego.BeforeRouter,FilterUser)

//1.1. 过滤器实现路由
//beego1.1.2 开始 Context.Input 中增加了 RunController 和 RunMethod, 这样我们就可以在执行路由查找之前,在 filter 中实现自己的路由规则.
//
//如下示例实现了如何实现自己的路由规则:
//
//var UrlManager = func(ctx *context.Context) {
//    // 数据库读取全部的 url mapping 数据
//    urlMapping := model.GetUrlMapping()
//    for baseurl,rule:=range urlMapping {
//        if baseurl == ctx.Request.RequestURI {
//            ctx.Input.RunController = rule.controller
//            ctx.Input.RunMethod = rule.method
//            break
//        }
//    }
//}
//
//beego.InsertFilter("/*",beego.BeforeRouter,UrlManager)

//1. Flash数据
//这个 flash 与 Adobe/Macromedia Flash 没有任何关系。它主要用于在两个逻辑间传递临时数据，flash 中存放的所有数据会在紧接着的下一个逻辑中调用后清除。一般用于传递提示和错误消息。它适合 Post/Redirect/Get 模式。下面看使用的例子：
//
//// 显示设置信息
//func (c *MainController) Get() {
//    flash:=beego.ReadFromRequest(&c.Controller)
//    if n,ok:=flash.Data["notice"];ok{
//        // 显示设置成功
//        c.TplName = "set_success.html"
//    }else if n,ok=flash.Data["error"];ok{
//        // 显示错误
//        c.TplName = "set_error.html"
//    }else{
//        // 不然默认显示设置页面
//        c.Data["list"]=GetInfo()
//        c.TplName = "setting_list.html"
//    }
//}
//
//// 处理设置信息
//func (c *MainController) Post() {
//    flash:=beego.NewFlash()
//    setting:=Settings{}
//    valid := Validation{}
//    c.ParseForm(&setting)
//    if b, err := valid.Valid(setting);err!=nil {
//        flash.Error("Settings invalid!")
//        flash.Store(&c.Controller)
//        c.Redirect("/setting",302)
//        return
//    }else if b!=nil{
//        flash.Error("validation err!")
//        flash.Store(&c.Controller)
//        c.Redirect("/setting",302)
//        return
//    }
//    saveSetting(setting)
//    flash.Notice("Settings saved!")
//    flash.Store(&c.Controller)
//    c.Redirect("/setting",302)
//}
//上面的代码执行的大概逻辑是这样的：
//
//Get 方法执行，因为没有 flash 数据，所以显示设置页面。
//用户设置信息之后点击递交，执行 Post，然后初始化一个 flash，通过验证，验证出错或者验证不通过设置 flash 的错误，如果通过了就保存设置，然后设置 flash 成功设置的信息。
//设置完成后跳转到 Get 请求。
//Get 请求获取到了 Flash 信息，然后执行相应的逻辑，如果出错显示出错的页面，如果成功显示成功的页面。
//默认情况下 ReadFromRequest 函数已经实现了读取的数据赋值给 flash，所以在你的模板里面你可以这样读取数据：

//{{.flash.error}}
//{{.flash.warning}}
//{{.flash.notice}}
//flash 对象有三个级别的设置：
//
//Notice 提示信息
//Warning 警告信息
//Error 错误信息

//1. URL构建
//如果可以匹配 URL ，那么 beego 也可以生成 URL 吗？当然可以。 UrlFor() 函数就是用于构建指定函数的 URL 的。它把对应控制器和函数名结合的字符串作为第一个参数，其余参数对应 URL 中的变量。未知变量将添加到 URL 中作为查询参数。 例如：
//
//下面定义了一个相应的控制器
//
//type TestController struct {
//    beego.Controller
//}
//
//func (this *TestController) Get() {
//    this.Data["Username"] = "astaxie"
//    this.Ctx.Output.Body([]byte("ok"))
//}
//
//func (this *TestController) List() {
//    this.Ctx.Output.Body([]byte("i am list"))
//}
//
//func (this *TestController) Params() {
//    this.Ctx.Output.Body([]byte(this.Ctx.Input.Params()["0"] + this.Ctx.Input.Params()["1"] + this.Ctx.Input.Params()["2"]))
//}
//
//func (this *TestController) Myext() {
//    this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
//}
//
//func (this *TestController) GetUrl() {
//    this.Ctx.Output.Body([]byte(this.UrlFor(".Myext")))
//}
//下面是我们注册的路由：
//
//beego.Router("/api/list", &TestController{}, "*:List")
//beego.Router("/person/:last/:first", &TestController{})
//beego.AutoRouter(&TestController{})
//那么通过方式可以获取相应的URL地址：
//
//beego.URLFor("TestController.List")
//// 输出 /api/list
//
//beego.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
//// 输出 /person/xie/asta
//
//beego.URLFor("TestController.Myext")
//// 输出 /Test/Myext
//
//beego.URLFor("TestController.GetUrl")
//// 输出 /Test/GetUrl
//1.1. 模板中如何使用
//默认情况下，beego 已经注册了 urlfor 函数，用户可以通过如下的代码进行调用
//
//{{urlfor "TestController.List"}}
//为什么不在把 URL 写死在模板中，反而要动态构建？有两个很好的理由：
//
//反向解析通常比硬编码 URL 更直观。同时，更重要的是你可以只在一个地方改变 URL ，而不用到处乱找。
//URL 创建会为你处理特殊字符的转义和 Unicode 数据，不用你操心。

//1. 多种格式数据输出
//2. JSON、XML、JSONP
//beego 当初设计的时候就考虑了 API 功能的设计，而我们在设计 API 的时候经常是输出 JSON 或者 XML 数据，那么 beego 提供了这样的方式直接输出：
//
//注意 struct 属性应该 为 exported Identifier 首字母应该大写
//
//JSON 数据直接输出：
//
//  func (this *AddController) Get() {
//      mystruct := { ... }
//      this.Data["json"] = &mystruct
//      this.ServeJSON()
//  }
//调用 ServeJSON 之后，会设置 content-type 为 application/json，然后同时把数据进行 JSON 序列化输出。
//
//XML 数据直接输出：
//
//  func (this *AddController) Get() {
//      mystruct := { ... }
//      this.Data["xml"]=&mystruct
//      this.ServeXML()
//  }
//调用 ServeXML 之后，会设置 content-type 为 application/xml，同时数据会进行 XML 序列化输出。
//
//jsonp 调用
//
//  func (this *AddController) Get() {
//      mystruct := { ... }
//      this.Data["jsonp"] = &mystruct
//      this.ServeJSONP()
//  }
//调用 ServeJSONP 之后，会设置 content-type 为 application/javascript，然后同时把数据进行 JSON 序列化，然后根据请求的 callback 参数设置 jsonp 输出。
//
//开发模式下序列化后输出的是格式化易阅读的 JSON 或 XML 字符串；在生产模式下序列化后输出的是压缩的字符串。

//1. 表单数据验证
//2. 表单验证
//表单验证是用于数据验证和错误收集的模块。
//
//2.1. 安装及测试
//安装：
//
//go get github.com/astaxie/beego/validation
//测试：
//
//go test github.com/astaxie/beego/validation

//2.2. 示例
//直接使用示例：
//
//import (
//    "github.com/astaxie/beego/validation"
//    "log"
//)
//
//type User struct {
//    Name string
//    Age int
//}
//
//func main() {
//    u := User{"man", 40}
//    valid := validation.Validation{}
//    valid.Required(u.Name, "name")
//    valid.MaxSize(u.Name, 15, "nameMax")
//    valid.Range(u.Age, 0, 18, "age")
//
//    if valid.HasErrors() {
//        // 如果有错误信息，证明验证没通过
//        // 打印错误信息
//        for _, err := range valid.Errors {
//            log.Println(err.Key, err.Message)
//        }
//    }
//    // or use like this
//    if v := valid.Max(u.Age, 140, "age"); !v.Ok {
//        log.Println(v.Error.Key, v.Error.Message)
//    }
//    // 定制错误信息
//    minAge := 18
//    valid.Min(u.Age, minAge, "age").Message("少儿不宜！")
//    // 错误信息格式化
//    valid.Min(u.Age, minAge, "age").Message("%d不禁", minAge)
//}
//通过 StructTag 使用示例：
//
//import (
//    "log"
//    "strings"
//
//    "github.com/astaxie/beego/validation"
//)
//
//// 验证函数写在 "valid" tag 的标签里
//// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
//// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
//// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
//// 各个函数的结果的 key 值为字段名.验证函数名
//type user struct {
//    Id     int
//    Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
//    Age    int    `valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
//    Email  string `valid:"Email; MaxSize(100)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
//    Mobile string `valid:"Mobile"` // Mobile 必须为正确的手机号
//    IP     string `valid:"IP"` // IP 必须为一个正确的 IPv4 地址
//}
//
//// 如果你的 struct 实现了接口 validation.ValidFormer
//// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
//func (u *user) Valid(v *validation.Validation) {
//    if strings.Index(u.Name, "admin") != -1 {
//        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
//        v.SetError("Name", "名称里不能含有 admin")
//    }
//}
//
//func main() {
//    valid := validation.Validation{}
//    u := user{Name: "Beego", Age: 2, Email: "dev@beego.me"}
//    b, err := valid.Valid(&u)
//    if err != nil {
//        // handle error
//    }
//    if !b {
//        // validation does not pass
//        // blabla...
//        for _, err := range valid.Errors {
//            log.Println(err.Key, err.Message)
//        }
//    }
//}
//StructTag 可用的验证函数：
//
//Required 不为空，即各个类型要求不为其零值
//Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
//Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
//Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
//MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
//MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
//Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
//Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
//Numeric 数字，有效类型：string，其他类型都将不能通过验证
//AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
//Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf("%v", obj).Match)
//AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
//Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
//IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
//Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
//Mobile 手机号，有效类型：string，其他类型都将不能通过验证
//Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
//Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
//ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证
//2.2.1. API 文档
//请移步 Go Walker。

//1. 错误处理
//我们在做 Web 开发的时候，经常需要页面跳转和错误处理，beego 这方面也进行了考虑，通过 Redirect 方法来进行跳转：
//
//func (this *AddController) Get() {
//    this.Redirect("/", 302)
//}
//如何中止此次请求并抛出异常，beego 可以在控制器中这样操作：
//
//func (this *MainController) Get() {
//    this.Abort("401")
//    v := this.GetSession("asta")
//    if v == nil {
//        this.SetSession("asta", int(1))
//        this.Data["Email"] = 0
//    } else {
//        this.SetSession("asta", v.(int)+1)
//        this.Data["Email"] = v.(int)
//    }
//    this.TplName = "index.tpl"
//}
//这样 this.Abort("401") 之后的代码不会再执行，而且会默认显示给用户如下页面：

//beego 框架默认支持 401、403、404、500、503 这几种错误的处理。用户可以自定义相应的错误处理，例如下面重新定义 404 页面：
//
//func page_not_found(rw http.ResponseWriter, r *http.Request){
//    t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
//    data :=make(map[string]interface{})
//    data["content"] = "page not found"
//    t.Execute(rw, data)
//}
//
//func main() {
//    beego.ErrorHandler("404",page_not_found)
//    beego.Router("/", &controllers.MainController{})
//    beego.Run()
//}
//我们可以通过自定义错误页面 404.html 来处理 404 错误。
//
//beego 更加人性化的还有一个设计就是支持用户自定义字符串错误类型处理函数，例如下面的代码，用户注册了一个数据库出错的处理页面：
//
//func dbError(rw http.ResponseWriter, r *http.Request){
//    t,_:= template.New("dberror.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/dberror.html")
//    data :=make(map[string]interface{})
//    data["content"] = "database is now down"
//    t.Execute(rw, data)
//}
//
//func main() {
//    beego.ErrorHandler("dbError",dbError)
//    beego.Router("/", &controllers.MainController{})
//    beego.Run()
//}
//一旦在入口注册该错误处理代码，那么你可以在任何你的逻辑中遇到数据库错误调用 this.Abort("dbError") 来进行异常页面处理。

//2. Controller 定义 Error
//从 1.4.3 版本开始，支持 Controller 方式定义 Error 错误处理函数，这样就可以充分利用系统自带的模板处理，以及 context 等方法。
//
//package controllers
//
//import (
//    "github.com/astaxie/beego"
//)
//
//type ErrorController struct {
//    beego.Controller
//}
//
//func (c *ErrorController) Error404() {
//    c.Data["content"] = "page not found"
//    c.TplName = "404.tpl"
//}
//
//func (c *ErrorController) Error501() {
//    c.Data["content"] = "server error"
//    c.TplName = "501.tpl"
//}
//
//
//func (c *ErrorController) ErrorDb() {
//    c.Data["content"] = "database is now down"
//    c.TplName = "dberror.tpl"
//}
//通过上面的例子我们可以看到，所有的函数都是有一定规律的，都是 Error 开头，后面的名字就是我们调用 Abort 的名字，例如 Error404 函数其实调用对应的就是 Abort("404")
//
//我们就只要在 beego.Run 之前采用 beego.ErrorController 注册这个错误处理函数就可以了
//
//package main
//
//import (
//    _ "btest/routers"
//    "btest/controllers"
//
//    "github.com/astaxie/beego"
//)
//
//func main() {
//    beego.ErrorController(&controllers.ErrorController{})
//    beego.Run()
//}

//1. 日志处理
//beego 之前介绍的时候说过是基于几个模块搭建的，beego 的日志处理是基于 logs 模块搭建的，内置了一个变量 BeeLogger，默认已经是 logs.BeeLogger 类型，初始化了 console，也就是默认输出到 console。
//
//1.1. 使用入门
//一般在程序中我们使用如下的方式进行输出：
//
//beego.Emergency("this is emergency")
//beego.Alert("this is alert")
//beego.Critical("this is critical")
//beego.Error("this is error")
//beego.Warning("this is warning")
//beego.Notice("this is notice")
//beego.Informational("this is informational")
//beego.Debug("this is debug")

//1.2. 设置输出
//我们的程序往往期望把信息输出到 log 中，现在设置输出到文件很方便，如下所示：
//
//beego.SetLogger("file", `{"filename":"logs/test.log"}`)
//更多详细的日志配置请查看日志配置
//
//这个默认情况就会同时输出到两个地方，一个 console，一个 file，如果只想输出到文件，就需要调用删除操作：
//
//beego.BeeLogger.DelLogger("console")
//1.3. 设置级别
//日志的级别如上所示的代码这样分为八个级别：
//
//LevelEmergency
//LevelAlert
//LevelCritical
//LevelError
//LevelWarning
//LevelNotice
//LevelInformational
//LevelDebug
//级别依次降低，默认全部打印，但是一般我们在部署环境，可以通过设置级别设置日志级别：
//
//beego.SetLevel(beego.LevelInformational)
//1.4. 输出文件名和行号
//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
//
//beego.SetLogFuncCall(true)
//开启传入参数 true, 关闭传入参数 false, 默认是关闭的.
//
//1.5. 完整示例
//func internalCalculationFunc(x, y int) (result int, err error) {
//    beego.Debug("calculating z. x:", x, " y:", y)
//    z := y
//    switch {
//    case x == 3:
//        beego.Debug("x == 3")
//        panic("Failure.")
//    case y == 1:
//        beego.Debug("y == 1")
//        return 0, errors.New("Error!")
//    case y == 2:
//        beego.Debug("y == 2")
//        z = x
//    default:
//        beego.Debug("default")
//        z += x
//    }
//    retVal := z - 3
//    beego.Debug("Returning ", retVal)
//
//    return retVal, nil
//}
//
//func processInput(input inputData) {
//    defer func() {
//        if r := recover(); r != nil {
//            beego.Error("Unexpected error occurred: ", r)
//            outputs <- outputData{result: 0, error: true}
//        }
//    }()
//    beego.Informational("Received input signal. x:", input.x, " y:", input.y)
//
//    res, err := internalCalculationFunc(input.x, input.y)
//    if err != nil {
//        beego.Warning("Error in calculation:", err.Error())
//    }
//
//    beego.Informational("Returning result: ", res, " error: ", err)
//    outputs <- outputData{result: res, error: err != nil}
//}
//
//func main() {
//    inputs = make(chan inputData)
//    outputs = make(chan outputData)
//    criticalChan = make(chan int)
//    beego.Informational("App started.")
//
//    go consumeResults(outputs)
//    beego.Informational("Started receiving results.")
//
//    go generateInputs(inputs)
//    beego.Informational("Started sending signals.")
//
//    for {
//        select {
//        case input := <-inputs:
//            processInput(input)
//        case <-criticalChan:
//            beego.Critical("Caught value from criticalChan: Go shut down.")
//            panic("Shut down due to critical fault.")
//        }
//    }
//}


//1. 概述
//2. 模型（Models）－ beego ORM
//beego ORM 是一个强大的 Go 语言 ORM 框架。她的灵感主要来自 Django ORM 和 SQLAlchemy。
//
//目前该框架仍处于开发阶段，可能发生任何导致不兼容的改动。
//
//已支持数据库驱动：
//
//MySQL：github.com/go-sql-driver/mysql
//PostgreSQL：github.com/lib/pq
//Sqlite3：github.com/mattn/go-sqlite3
//以上数据库驱动均通过基本测试，但我们仍需要您的反馈。
//
//ORM 特性：
//
//支持 Go 的所有类型存储
//轻松上手，采用简单的 CRUD 风格
//自动 Join 关联表
//跨数据库兼容查询
//允许直接使用 SQL 查询／映射
//严格完整的测试保证 ORM 的稳定与健壮
//更多特性请在文档中自行品读。
//
//安装 ORM：
//
//go get github.com/astaxie/beego/orm

//2.1. 修改日志
//2016-01-18: 规范了数据库驱动的命名
//2014-03-10: GetDB 从注册的数据库中返回 *sql.DB. ResetModelCache 重置已注册的模型struct
//2014-02-10: 随着 beego1.1.0 的发布提交的改进
//
//关于 时区设置
//
//新增的 api: Ormer.InsertMulti Ormer.ReadOrCreate RawSeter.RowsToMap RawSeter.RowsToStruct orm.NewOrmWithDB
//
//改进的 api: RawSeter.Values 支持设置 columns RawSeter.ValuesList 支持设置 columns RawSeter.ValuesFlat 支持设置 column RawSeter.QueryRow/QueryRows 从对应每个strcut field位置的赋值，改为对应名称取值（不需要对应好字段数量与位置）
//
//2013-10-14: 自动载入关系字段，多对多关系操作，完善关系查询
//
//2013-10-09: 原子操作更新值
//2013-09-22: RegisterDataBase maxIdle / maxConn 设置为可选参数, MySQL 自定义引擎
//2013-09-16: 支持设置 空闲链接数 和 最大链接数 SetMaxIdleConns / SetMaxOpenConns
//2013-09-12: Read 支持设定条件字段 Update / All / One 支持设定返回字段
//2013-09-09: Raw SQL QueryRow/QueryRows 功能完成
//2013-08-27: 自动建表继续改进
//2013-08-19: 自动建表功能完成
//2013-08-13: 更新数据库类型测试
//2013-08-13: 增加 Go 类型支持，包括 int8、uint8、byte、rune 等
//2013-08-13: 增强 date／datetime 的时区支持

//2.2. 快速入门
//2.2.1. 简单示例
//package main
//
//import (
//    "fmt"
//    "github.com/astaxie/beego/orm"
//    _ "github.com/go-sql-driver/mysql" // import your used driver
//)
//
//// Model Struct
//type User struct {
//    Id   int
//    Name string `orm:"size(100)"`
//}
//
//func init() {
//    // set default database
//    orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
//
//    // register model
//    orm.RegisterModel(new(User))
//
//    // create table
//    orm.RunSyncdb("default", false, true)
//}
//
//func main() {
//    o := orm.NewOrm()
//
//    user := User{Name: "slene"}
//
//    // insert
//    id, err := o.Insert(&user)
//    fmt.Printf("ID: %d, ERR: %v\n", id, err)
//
//    // update
//    user.Name = "astaxie"
//    num, err := o.Update(&user)
//    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
//
//    // read one
//    u := User{Id: user.Id}
//    err = o.Read(&u)
//    fmt.Printf("ERR: %v\n", err)
//
//    // delete
//    num, err = o.Delete(&u)
//    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
//}

//2.2.2. 关联查询
//type Post struct {
//    Id    int    `orm:"auto"`
//    Title string `orm:"size(100)"`
//    User  *User  `orm:"rel(fk)"`
//}
//
//var posts []*Post
//qs := o.QueryTable("post")
//num, err := qs.Filter("User__Name", "slene").All(&posts)

//2.2.3. SQL 查询
//当您无法使用 ORM 来达到您的需求时，也可以直接使用 SQL 来完成查询／映射操作。
//
//var maps []orm.Params
//num, err := o.Raw("SELECT * FROM user").Values(&maps)
//for _,term := range maps{
//    fmt.Println(term["id"],":",term["name"])
//}

//2.2.4. 事务处理
//o.Begin()
//...
//user := User{Name: "slene"}
//id, err := o.Insert(&user)
//if err == nil {
//    o.Commit()
//} else {
//    o.Rollback()
//}

//2.2.5. 调试查询日志
//在开发环境下，您可以使用以下指令来开启查询调试模式：
//
//func main() {
//    orm.Debug = true
//...

//开启后将会输出所有查询语句，包括执行、准备、事务等。
//
//例如：
//
//[ORM] - 2013-08-09 13:18:16 - [Queries/default] - [    db.Exec /     0.4ms] -     [INSERT INTO `user` (`name`) VALUES (?)] - `slene`
//...
//注意：我们不建议您在部署产品后这样做。
















