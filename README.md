# working-with-go-zh-cn

通过一系列Go代码实例学习Go语言。 https://github.com/mkaz/working-with-go.git

## Install Go

Go [官方二进制发行版](https://golang.org/dl/)可支持FreeBSD、Linux、Mac OS X（Snow Leopard、Lion和Mountain Lion）、 NetBSD和Windows操作系统以及32位（386）和64位（amd64）的 x86 处理器架构。

如果上不去，可以去[golangtc.com](http://golangtc.com/download) 下载。

## Clone and Go

Go环境安装好后，克隆本仓库
	
	$ git clone https://github.com/jackerloong/working-with-go-zh-ch

然后你可以运行里面的程序，他们之间都编好了号，相互之间独立。如:

运行第一个实例程序:

	$ go run 01-hello.go

如果你想构建一个可执行文件，运行:

	$ go build -o hello 01-hello.go
	$ ./hello

## Development Env

Sublime 是个不错的选择， 安装了Package Control后，按Ctrl+Shift+P输入Install Package，在弹出的输入框中输入 GoSublime， 安装后，编辑.go文件将自带提示和自动格式化。

LiteIDE是一个专业的跨平台Go语言集成开发工具，[下载](http://golangtc.com/download/liteide)安装后，设置好GOROOT环境变量即可。

## License

Working with Go is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>. 
