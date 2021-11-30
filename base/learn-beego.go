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




