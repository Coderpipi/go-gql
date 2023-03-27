# go-gql搭建过程与学习心得![](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/owner-pip1-blue.svg)![](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/type-doc-brightgreen.svg)

## 介绍

1. `GraphQL`:
   **简介**: `A query language for your API`
   **描述**: `GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.(GraphQL是一种用于API的查询语言，也是用你的现有数据完成这些查询的运行时间。GraphQL为你的API中的数据提供了一个完整的、可理解的描述，使客户有能力准确地要求他们所需要的东西，而不是更多，使API更容易随着时间的推移而发展，并使强大的开发人员工具成为可能。)`
   **教程**: [GraphQL](https://graphql.org/)
2. `graphql-go`:
   **简介**: `一个支持GraphQL Server的go第三方库`
   **github地址**: [graphql-go](https://github.com/graph-gophers/graphql-go)

## 搭建教程

1. 准备工作
   + 操作系统: MacOS BigSur 11.6.3
   + 安装Go二进制包, 这里使用的是Go(go1.19)
   + 安装集成开发环境(Goland), 并安装插件: GraphQL
2. 初始化Go module项目
   + 打开**Goland**新建一个``Go Project`:
     ![CleanShot 2023-03-27 at 18.41.48](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/CleanShot%202023-03-27%20at%2018.41.48.png)

+ 创建`main.go`, 并运行项目, 如果控制台输出了`Hello Go`, 则表示项目创建成功且正常运行
  ```go
  package main
  
  import "fmt"
  
  func main() {
  	fmt.Println("Hello Go")
  }
  ```

+ ![CleanShot 2023-03-27 at 18.47.47](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/CleanShot%202023-03-27%20at%2018.47.47.png)

+ 使用Gin框架创建一个简单的server

  1. `terminal`执行以下命令, 安装`Gin`

     ```shell
     go get github.com/gin-gonic/gin@v1.7.7 
     ```

  2. 在项目根目录下创建一个`cmd`文件夹,`cmd`下创建一个`router.go`文件

     ```go
     package cmd
     
     import (
     	"github.com/gin-gonic/gin"
     	"net/http"
     )
     
     func SetupRouter() *gin.Engine {
     	r := gin.Default()
     
     	r.GET("ping", func(c *gin.Context) {
     		c.String(http.StatusOK, "pong")
     	})
     	return r
     }
     ```

  3. 在`main.go`中调用`SetupRouter`函数, 并启动项目, 启动成功页面如下图所示
     ```go
     package main
     
     import "empty-project/cmd"
     
     func main() {
     	r := cmd.SetupRouter()
     
     	r.Run(":8080")
     }
     ```

     ![CleanShot 2023-03-27 at 18.58.16](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/CleanShot%202023-03-27%20at%2018.58.16.png)

  4. 测试`server`是否正常响应请求

     ```shell
     curl http://localhost:8080/ping
     ```

     如下图则表示请求正常被响应

     ![CleanShot 2023-03-27 at 19.00.28](https://lbw-picgo.oss-cn-shenzhen.aliyuncs.com/CleanShot%202023-03-27%20at%2019.00.28.png)