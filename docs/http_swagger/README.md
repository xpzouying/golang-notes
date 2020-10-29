# swagger入门教程

对Golang http server增加swagger文档。

使用最常见的http router之一gin来启动http server的demo，

- [gin-swagger](https://github.com/swaggo/gin-swagger)提供了gin swagger框架
- [wagger注释标准](https://swaggo.github.io/swaggo.io/declarative_comments_format/)

1. 生成HTTP文档：
   ```bash
   # 编写对应的http处理函数，增加对应的注释后，运行下列命令生成文档
   swag init
   ```
2. 运行程序：
   ```bash
   go run .
   ```
3. 打开swagger API接口文档：[swagger index](http://zydev:8080/swagger/index.html)
