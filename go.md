#### gvm go多版本管理
>https://github.com/moovweb/gvm
```
安装
    gvm install go1.18.2 -B
使用    
    gvm use go1.18.2
设置默认
    gvm use go1.18.2 --default

```
#### GO111MODULE
```
GO111MODULE为空默认auto
GO111MODULE=auto
    在 GO111MODULE=auto 模式下，你可以在包含 go.mod 文件的目录中使用模块模式，无论这个目录是否位于 GOPATH/src 下。

GO111MODULE=on
    强制使用Go modules，即使在`GOPATH/src`目录下也会使用`go.mod`文件，并忽略`GOPATH`中的包。

GO111MODULE=off
    不使用模块支持，Go命令会忽略 go.mod 文件并且会从`GOPATH`中查找依赖。

在 Go 1.13 及更高版本中，如果你在一个包含 go.mod 文件的目录中运行 go 命令，模块模式会自动启动，无需设置 GO111MODULE 环境变量。然而，如果你在没有 go.mod 文件的目录中工作，或者你想显式地控制模块模式的行为，你可能需要设置 GO111MODULE 环境变量。
```
#### 修改配置 go env -w
```
修改.zshrc
    export GO111MODULE=on
    source .zshrc

也可以在命令行执行
    export GO111MODULE=on
```
#### 传统模式（GOPATH）
>src目录下
```
go clean -modcache #清理moudle 缓存


go get 升级
    go list -m -u all #来检查可以升级的package
    运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
    运行 go get -u=patch 将会升级到最新的修订版本
    运行 go get package@version 将会升级到指定的版本号version
    运行go get如果有版本的更改，那么go.mod文件也会更改

go get = git clone + go install
```
#### mod模式
>mod目录下
```
初始化go.mod
    go mod init 模块名称

下载包
    方式1：
        先定义go.mod，然后执行go mod download
        下载路径在：
            $GOPATH/pkg/mod/，同时缓存一份到：$GOPATH/pkg/mod/cache/download/

    方式2：
        根据源码import的第三方库，自动下载依赖包的命令：go mod tidy。
        下载路径在：
            $GOPATH/pkg/mod/，同时缓存一份到：$GOPATH/pkg/mod/cache/download/

    方式3：go get（不再支持非mod得项目）
        go get


mod依赖包失效：
    使用go.mod就简单了，在go.mod文件里用 replace 替换包，例如
        replace golang.org/x/text => github.com/golang/text latest
        replace (
            golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
        )


go mod命令
    go mod download #下载modules到本地cache，目前所有模块版本数据均缓存在 $GOPATH/pkg/mod和 ​$GOPATH/pkg/sum 下
    go mod edit #选项有-json、-require和-exclude

    go mod graph #以文本模式打印模块需求图
    go mod tidy  #删除错误或者不使用的modules
    go mod verify #验证依赖是否正确
    go mod why #查找依赖

    go mod vendor #从mod中拷贝到项目的vendor目录下
        所谓 vendor 机制，就是每个项目的根目录下可以有一个 vendor 目录，里面存放了该项目的依赖的 package。go build 的时候会先去 vendor 目录查找依赖，如果没有找到会再去 GOPATH 目录下查找。

go build -mod vendor


```
#### 老项目使用mod
```
以前老项目如何用新的包管理:
    进入目录，运行 go mod init 模块名称
        go build 或者 go run 一次

```