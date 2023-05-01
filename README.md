# go-bootloader

##  Quick Start
### go get
```bash
go get -u github.com/xiaolongliu98/go-bootloader
```
### apply in your project

```go
package main

import (
	"context"
	"github.com/xiaolongliu98/go-bootloader"
	"github.com/xiaolongliu98/go-bootloader/ctx"
	"github.com/xiaolongliu98/go-bootloader/model"
)

type FirstLoader struct {
}

func (l FirstLoader) Load(ctx ctx.LoaderContext) error {
    // your project startup code
	return nil
}

func (l FirstLoader) Require() []string {
	return []string{}
}

type SecondLoader struct {
}

func (l SecondLoader) Load(ctx ctx.LoaderContext) error {
	// your project startup code
	return nil
}

func (l SecondLoader) Require() []string {
	return []string{"FirstLoader"}
}

func main() {
	list := model.LoaderList{
		&FirstLoader{},
		&SecondLoader{},
	}
	bootloader.Load(context.TODO(), list)
	bootloader.WaitShutdownGracefully()
}
```
更多的例子请参考example目录下的代码

For more examples, please refer to the code in the example directory.


# Introduction
## 一个简单的启动依赖加载器
## A simple startup dependency loader.

在完成一个项目时，我们常会遇到这种情况：

程序启动时，有许多模块需要加载，例如：
1. 获得项目路径
2. 解析并绑定配置文件
3. 加载数据库
4. 启动Server
...

如果你细心，会发现它们实际上是有一个加载依赖顺序存在的：
1. 我们只有获得项目路径后才能获取配置文件的路径
2. 有了配置文件我们才能加载数据库
3. 数据库连接成功后我们才启动Server
...


另外，一个专业的程序员是绝不会忽略单元测试的，在你的项目中，你是否遇到过编写单元测试时，会操心如何加载上述的依赖呢？非常麻烦！

使用bootloader还有一个好处就是：支持最小范围的模块启动的单元测试


## 本仓库设计了一个bootloader模块加载机制，使用依赖拓扑序进行依次加载
也就是说把你需要加载的模块做出一个列表填进去即可！



---
When completing a project, we often encounter this situation:

When the program starts, many modules need to be loaded, such as:
1. Obtain the project path
2. Parse and bind the configuration file
3. Load the database
4. Start the server
...

If you are careful, you will find that they actually have a loading dependency order:
1. We can only get the path of the configuration file after obtaining the project path
2. With the configuration file, we can load the database
3. After the database connection is successful, we start the server
...

In addition, a professional programmer would never overlook unit testing. Have you ever encountered the trouble of loading dependencies when writing unit tests in your project? It's very troublesome!

Using a bootloader also has another advantage: supporting unit testing of modules that start in the smallest range.

## This repository has designed a bootloader module loading mechanism, using dependency topological order for sequential loading.
That is to say, just make a list of the modules you need to load and fill it in!
