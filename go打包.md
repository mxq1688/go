#### go build
```
go build [-o 输出名] [-i] [编译标记] [包名]
    go build -o /go/bin/app

    这里的包名
        当使用 go build 命令时，如果没有指定任何包名，则默认编译当前目录下的主包（即含有 main 函数的包），并且这个包会被编译成一个可执行文件。

            当你在含有 main 函数的包的目录中时，可以直接运行，下面3条命令等价
                go build 
                go build .
                go build main.go
            
            指定包 比如main.go不在go.mod目录下
                go build xxx/main.go


参数
    -o output：指定输出文件名或路径。
    -v：在构建过程中显示每个被构建的包。
    -x：打印构建命令。
    -i：忽略缓存中的包信息，总是重新下载依赖包。
    -mod=[mod|vendor|readonly]：控制模块模式的行为。
    -tags list：指定要构建的包需要的构建标签列表。

mod模式下
    如果没有使用 -o 选项来指定输出文件的位置和名称，go build 将会在当前工作目录下生成一个可执行文件，文件名默认为 go mod 初始化的模块名称或者当前目录下的包名（即 main 包所在的目录名）。
path模式
    在 bin目录下

```
>构建特定平台上的二进制文件
```
GOOS=linux GOARCH=amd64 go build -o myapp-linux-amd64

编译跨平台的只需要修改GOOS、GOARCH、CGO_ENABLED三个环境变量即可

    GOOS：目标平台的操作系统(darwin、freebsd、linux、windows)

    GOARCH：目标平台的体系架构32位还是64位(386、amd64、arm)
```