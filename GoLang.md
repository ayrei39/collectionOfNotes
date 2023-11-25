# GoLang

## Environment Setup

Download：https://go.dev/

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
> - `go get`: 用于获取并安装远程仓库中的 Go 包或命令行工具。(需要模块化)
> - `go test`: 用于运行测试。该命令会自动查找项目中符合命名规范的测试文件，并执行这些测试。(需要模块化)

### What does 'package main' mean?

package = project = workspace 

公共源代码文件的集合

一个包里可以有很多相关文件

每个文件第一行都必须声明所属的包

同一个文件夹下只能有一个包名，否则编译报错

#### Why "main"?

go中有2中不同类型的包

- 可执行类型 Executable,编译时生成一个可执行文件
- 可重用类型 Reusable,代码依赖项,存放可重用的逻辑或辅助函数或有助于我们重用的东西

package main - 可执行类型，必须包含一个名为`func`的函数

package others - 可重用类型

### What does 'import "fmt"' mean?

让该文件可以访问名为“fmt”的包的所有代码&函数

fmt：go语言默认包含的标准库包，format缩写

标准库：

- fmt
- math
- debug
- io
- encoding
- crypto

参考：https://pkg.go.dev/std

也可以导入自己的包

### What's that 'func' thing?

函数

```go
func funcName() {
    // func code
}
```



#### How is the main.go file organized?

- 包声明
- 导包 we need
- 声明函数 tell Go to do things

## Real first project

> Cards 模拟扑克
>
> - newDeck 新牌
> - print 打印所有牌
> - shuttle 洗牌
> - deal 手牌
> - saveToFile 保存到本地
> - newDeckFromFile 从本地加载

## Variable Declarations

```go
var varName varType = varValue
var identifier1, identifier2 type = value1, value2
```

**Basic Go Types**

- bool
- string
- int
- float64/float32

new version can do

```go
var a = 1 	// 自动识别
var b := 2	// 简短声明
```

## Function and Return Types

```go
func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
```

## Array & Slice

Array: Fixed length list of things

Slice: An array that can grow or shrink. Every element in a slice must be of same type

```go
// 声明一个未指定大小的数组来定义切片
var identifier []type
// 切片不需要说明长度。
// 或使用 make() 函数来创建切片:
var slice1 []type = make([]type, len)
// 也可以简写为
slice1 := make([]type, len)
```

```go
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Five of Spades")
	fmt.Println(cards)
}

// [Ace of Diamonds Five of Diamonds]
// [Ace of Diamonds Five of Diamonds Five of Spades]
```

### For loop

```go
for i, card := range cards {
		fmt.Println(i, card)
	}

// 0 Ace of Diamonds
// 1 Five of Diamonds
// 2 Five of Spades
```

## OO Approach vs Go Approach

Go不是一种面向对象的编程语言，没有类的概念

**Base Go Types**

- string
- integer
- float
- array
- map

“Extent ” a base type and add some extra functionality to it

→ type deck []sring (create an array of strings and add a bunch of functions specifically made to work with it)

→ Functions with ‘deck’ as a receiver

### ‘cards’ folder

- main.go
- deck.go
- deck_test.go
