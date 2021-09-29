```

go env -w会将配置写到 GOENV="/Users/youdi/Library/Application Support/go/env"
    go env -w #修改配置
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的

从指定源上面下载或者更新指定的代码和依赖，并对他们进行编译和安装
    go get = git clone + go install

项目中使用了go.mod时可以使用以下命令自动下载全部依赖
    方法一
        go get -d -v ./...
            -d标志只下载代码包，不执行安装命令；
            -v打印详细日志和调试日志。这里加上这个标志会把每个下载的包都打印出来；
    方法二
        go mod tidy

go mod命令
    go mod download #下载modules到本地cache，目前所有模块版本数据均缓存在 $GOPATH/pkg/mod和 ​$GOPATH/pkg/sum 下
    go mod edit #选项有-json、-require和-exclude

    go mod graph #以文本模式打印模块需求图
    go mod tidy  #删除错误或者不使用的modules
    go mod verify #验证依赖是否正确
    go mod why #查找依赖
    go mod vendor #从mod中拷贝到项目的vendor目录下

go clean -modcache #清理moudle 缓存

GO111MODULE 有三个值：off, on和auto（默认值）。
    auto 自动模式下，项目在$GOPATH/src里会使用$GOPATH/src的依赖包，在$GOPATH/src外，就使用go.mod 里 require的包
    on 开启模式，1.12后，无论在$GOPATH/src里还是在外面，都会使用go.mod 里 require的包
    off 关闭模式，就是老规矩。

mod依赖包失效：
    使用go.mod就简单了，在go.mod文件里用 replace 替换包，例如
        replace golang.org/x/text => github.com/golang/text latest
        replace (
            golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
        )

以前老项目如何用新的包管理:
    如果用auto模式，把项目移动到$GOPATH/src外
    进入目录，运行 go mod init 模块名称
    go build 或者 go run 一次

go get 升级
    go list -m -u all #来检查可以升级的package
    运行 go get -u 将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
    运行 go get -u=patch 将会升级到最新的修订版本
    运行 go get package@version 将会升级到指定的版本号version
    运行go get如果有版本的更改，那么go.mod文件也会更改


go build -mod vendor
```