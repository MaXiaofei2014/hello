package hello

import "fmt"

/**
- 使用go module发布包

在上面的小节中我们学习了如何在项目中引入别人提供的依赖包，那么当我们想要在社区发布一个自己编写的代码包或者在公司内部编写一个供内部使用的公用组件时，我们该怎么做呢？接下来，我们就一起编写一个代码包并将它发布到github.com仓库，让它能够被全球的Go语言开发者使用。

我们首先在自己的 github 账号下新建一个项目，并把它下载到本地。我这里就以创建和发布一个名为hello的项目为例进行演示。这个hello包将对外提供一个名为SayHi的函数，它的作用非常简单就是向调用者发去问候。

$ git clone https://github.com/q1mi/hello
$ cd hello
我们当前位于hello项目目录下，执行下面的命令初始化项目，创建go.mod文件。需要注意的是这里定义项目的引入路径为github.com/q1mi/hello，读者在自行测试时需要将这部分替换为自己的仓库路径。

hello $ go mod init github.com/q1mi/hello
go: creating new go.mod: module github.com/q1mi/hello
接下来我们在该项目根目录下创建 hello.go 文件，添加下面的内容：

package hello

import "fmt"

func SayHi() {
	fmt.Println("你好，我是七米。很高兴认识你。")
}
然后将该项目的代码 push 到仓库的远端分支，这样就对外发布了一个Go包。其他的开发者可以通过github.com/q1mi/hello这个引入路径下载并使用这个包了。

一个设计完善的包应该包含开源许可证及文档等内容，并且我们还应该尽心维护并适时发布适当的版本。github 上发布版本号使用git tag为代码包打上标签即可。

hello $ git tag -a v0.1.0 -m "release version v0.1.0"
hello $ git push origin v0.1.0
经过上面的操作我们就发布了一个版本号为v0.1.0的版本。

Go modules中建议使用语义化版本控制，其建议的版本号格式如下：
其中：

主版本号：发布了不兼容的版本迭代时递增（breaking changes）。
次版本号：发布了功能性更新时递增。
修订号：发布了bug修复类更新时递增。
发布新的主版本

现在我们的hello项目要进行与之前版本不兼容的更新，我们计划让SayHi函数支持向指定人发出问候。更新后的SayHi函数内容如下：

package hello

import "fmt"

// SayHi 向指定人打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s，我是七米。很高兴认识你。\n", name)
}
由于这次改动巨大（修改了函数之前的调用规则），对之前使用该包作为依赖的用户影响巨大。因此我们需要发布一个主版本号递增的v2版本。在这种情况下，我们通常会修改当前包的引入路径，像下面的示例一样为引入路径添加版本后缀。

// hello/go.mod

module github.com/q1mi/hello/v2

go 1.16
把修改后的代码提交：

hello $ git add .
hello $ git commit -m "feat: SayHi现在支持给指定人打招呼啦"
hello $ git push
打好 tag 推送到远程仓库。

hello $ git tag -a v2.0.0 -m "release version v2.0.0"
hello $ git push origin v2.0.0
这样在不影响使用旧版本的用户的前提下，我们新的版本也发布出去了。想要使用v2版本的代码包的用户只需按修改后的引入路径下载即可。

go get github.com/q1mi/hello/v2@v2.0.0
在代码中使用的过程与之前类似，只是需要注意引入路径要添加 v2 版本后缀。

package main

import (
	"fmt"

	"github.com/q1mi/hello/v2" // 引入v2版本
)

func main() {
	fmt.Println("现在是假期时间...")

	hello.SayHi("张三") // v2版本的SayHi函数需要传入字符串参数
}
废弃已发布版本

如果某个发布的版本存在致命缺陷不再想让用户使用时，我们可以使用retract声明废弃的版本。例如我们在hello/go.mod文件中按如下方式声明即可对外废弃v0.1.2版本。

module github.com/q1mi/hello

go 1.16


retract v0.1.2
用户使用go get下载v0.1.2版本时就会收到提示，催促其升级到其他版本。


*/

// SayHi 向指定人打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s，我是大飞。很高兴认识你。\n", name)
}
