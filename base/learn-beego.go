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


















