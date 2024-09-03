#### gvm go多版本管理
>https://github.com/moovweb/gvm
```
安装gvm的时候需要设置终端代理

clashx
    export http_proxy=http://127.0.0.1:7890
    export https_proxy=http://127.0.0.1:7890
```
>使用
```
列出当前可以安装的 Go 版本
    gvm listall
安装
    gvm install go1.18.2 -B
    不加-B，将尝试安装预编译的 Go 二进制文件，而不是从源代码编译 Go
    -B 从源代码编译指定的 Go 版本
使用    
    gvm use go1.18.2
设置默认
    gvm use go1.18.2 --default

切换版本后统一设置GOPATH
    方法1：
        先在.zshrc中设置 GOPATH=/Users/xxx/go
        source ~/.zshrc
    方法2：
        gvm pkgenv命令修改（只能改当前版本）
        或者修改$GVM_ROOT|environments文件夹里，进去把GOPATH批量替换了就行了
    方法3：
        先在.zshrc中设置 GOPATH=/Users/xxx/go
        然后在~/.gvm/scripts/env/use gvm_use方法最后设置 source ~/.zshrc

    .zshrc 按照下面的顺序
        [[ -s "/Users/meng/.gvm/scripts/gvm" ]] && source "/Users/meng/.gvm/scripts/gvm"
        export GOPATH=/Users/meng/stu/go #项目路径
        export GOBIN=$GOPATH/bin   #GOLANG运行目录一般不可修改
        export PATH=$PATH:$GOBIN

```
#### go命令
```
修改配置
    修改.zshrc
        export GO111MODULE=on
        source .zshrc

    也可以在命令行执行
        export GO111MODULE=on

    修改配置 go env -w key=value
        go env -w GO111MODULE=on

go env 查看配置
    go env GO111MODULE 查看某个字段

运行
    go run main.go

打包
    go build
```
#### GO111MODULE
```
GO111MODULE=auto
    在 GO111MODULE=auto 模式下，你可以在包含 go.mod 文件的目录中使用模块模式，无论这个目录是否位于 GOPATH/src 下。

GO111MODULE=on
    强制使用Go modules，即使在`GOPATH/src`目录下也会使用`go.mod`文件，并忽略`GOPATH`中的包。

GO111MODULE=off
    不使用模块支持，Go命令会忽略 go.mod 文件并且会从`GOPATH`中查找依赖。


在 Go 1.13 及更高版本中
    如果你在一个包含 go.mod 文件的目录中运行 go 命令，模块模式会自动启动，无需设置 GO111MODULE 环境变量。
    然而，如果你在没有 go.mod 文件的目录中工作，或者你想显式地控制模块模式的行为，你可能需要设置 GO111MODULE 环境变量。

GO111MODULE为空默认auto
    Go1.16+ GO111MODULE 默认为 on

```
#### 传统模式（GOPATH）
>src目录下
```
GOPATH 
    是Golang 1.5版本之前一个重要的环境变量配置，是存放 Golang 项目代码的文件路径
    使用 GOPATH 模式下，我们需要将应用代码存放在固定的$GOPATH/src目录下，并且如果执行go get来拉取外部依赖会自动下载并安装到$GOPATH目录下。

Go1.18及以上 go get 不再执行编译和安装工作是指GO111MODULE=on在没有go.mod的目录下不能用，具体下面
    GO111MODULE=on 
        没有go.mod 不能用
        有go.mod，就是mod模式，下载到GOPATH/pkg/mod目录下

    GO111MODULE=auto
        没有go.mod，go get能下载到项目目录下(GOPATH/src) (Go1.22不支持)
        有go.mod，就是mod模式，下载到GOPATH/pkg/mod目录下

    GO111MODULE=off
        不管有没有go.mod，go get都下载到项目目录下(GOPATH/src) (Go1.22不支持)

Go1.18版本以下的：go get
    GO111MODULE=on 
        没有go.mod，下载到GOPATH/pkg/mod目录下 go get -u xxx
        有go.mod，就是mod模式，下载到GOPATH/pkg/mod目录下

    GO111MODULE=auto 
        没有go.mod，go get能下载到项目目录下(GOPATH/src)
        有go.mod，就是mod模式，下载到GOPATH/pkg/mod目录下

    GO111MODULE=off
        不管有没有go.mod，go get都下载到项目目录下(GOPATH/src)

go get 升级
    go list -m -u all #来检查可以升级的package
    运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
    运行 go get -u=patch 将会升级到最新的修订版本
    运行 go get package@version 将会升级到指定的版本号version
    运行go get如果有版本的更改，那么go.mod文件也会更改

go get = git clone + go install
```
#### mod模式（代码不用放GOPATH）
>mod目录下
```
初始化go.mod
    go mod init 模块名称

下载包（先定义go.mod）
    方式1：go mod download
        先定义go.mod的require，然后执行go mod download
        下载路径在：
            $GOPATH/pkg/mod/，同时缓存一份到：$GOPATH/pkg/mod/cache/download/

    方式2：go mod tidy
        根据源码import的第三方库，自动下载依赖包的命令：go mod tidy
        下载路径在：
            $GOPATH/pkg/mod/，同时缓存一份到：$GOPATH/pkg/mod/cache/download/

    方式3：go get（新版本不再支持非mod得项目）

清理moudle缓存
    go clean -modcache 

go mod 命令
    go mod graph #以文本模式打印模块需求图
    go mod verify #验证依赖是否正确
    go mod why #查找依赖

    go mod vendor 
        运行 go mod vendor 时：
            1、复制依赖项：将项目 go.mod 文件中列出的所有依赖项及其子依赖项复制到 vendor/ 目录下。
            2、更新 go.sum：更新 go.sum 文件以反映复制的依赖项的校验和。
            3、保持版本一致性：确保所有依赖项的版本与 go.mod 文件中指定的版本相匹配
        实现以下目的：
            离线构建：即使没有网络连接，你也可以构建项目，因为所有必需的依赖项都已存在于本地。
            版本控制：将依赖项版本控制在版本控制系统中，确保团队成员使用的依赖项版本相同。
            构建可靠性：确保每次构建使用相同的依赖项版本，从而避免因依赖项版本变化而导致的构建失败或不一致。

运行
    go run main.go

打包
    go build  也会优先使用 vendor/ 目录中的依赖项
    
    go build -mod vendor  优先使用 vendor/ 目录中的依赖项

mod依赖包失效：


```
#### 老项目使用mod
```
设置auto模式

进入目录
    1、运行 go mod init 模块名称生成go.mod文件
    2、下载包，参考上面
    3、go build 或者 go run
```