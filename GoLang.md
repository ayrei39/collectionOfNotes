# GoLang

## Environment Setup

[Download](https://go.dev/)

> 会自动配置环境变量：`Path`:`Your Go Path\bin`

Visual Studio Code

Plugins: Go

> 第一次新建`.go`文件时会提示缺少`gopls`，vsCode中网络问题可能会下载失败
>
> 解决办法：
>
> 配置环境变量：`GOPATH`:`Your Go Path`
>
> git bash 网络代理，`go install golang.org/x/tools/gopls@latest`

> `gopls`需要识别项目为Go模块，模块初始化：`go mod init`
>
> 对于简单代码文件，可以不初始化，报错可以忽略

## Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

### How do we run the code in our project?

- `go build`: 编译项目成一个可执行文件（需要`package main`）
- `go run`: 编译并执行go文件

> 其他：
>
> -  `go fmt`: 格式化 Go 代码文件。它会根据 Go 语言的编码规范，对代码文件进行格式化，确保它们符合一致的风格。
> - `go install`: 用于编译和安装 Go 程序。该命令会编译指定的 Go 源码文件或包，并将生成的可执行文件或库安装到 Go 工作区的 `bin` 或 `pkg` 目录中。
> - `go get`: 用于获取并安装远程仓库中的 Go 包或命令行工具。
> - `go test`: 用于运行测试。该命令会自动查找项目中符合命名规范的测试文件，并执行这些测试。

### What does 'package main' mean?

p6

### What does 'import "fmt"' mean?

### What's that 'func' thing?

#### How is the main.go file organized?