# 工程目录规范

最近在做团队项目的目录结构的规范，发现这块问题要想做好也挺困难的，学习一部分资料后，总结出该文章。

## 目标

这边文章围绕下列几点聊一下相关想法：

1.  如何进行分层。

1.  每一层的职责是什么。

1.  每一层的数据定义是什么。

1.  每一层数据的流转是怎么样的。

1.  每一层的依赖是什么？

1.  如何解决每一层的依赖问题。

1.  如何保障代码的可测试性。

定义上述规范，我们希望最终能够达成的目标：

1.  服务的健壮性。**可测试性**， Bob 大叔在 [《代码整洁之道》](https://book.douban.com/subject/4199741/) 中也提到了，只有代码随时的、快速的能进行测试，才能保证代码变得越来越好。

2.  解决各种各样依赖的问题，尽可能让单元模块具备：独立性。

## 分层规范

从根目录，开始分为：

-   `/api`

    **职责：**

    -   定义接口协议 DTO 结构体的定义
    -   路由注册
    -   讲请求/响应序列化和反序列化

    **测试方法：**

    -   使用 ` [httptest](https://pkg.go.dev/net/http/httptest) ` 模拟 http server ， 然后请求接口

    **测试关注点：**

    -   请求是否被转换成预定义的 struct。
    -   响应是否被转换成预期的 struct。

-   `/cmd`

    **职责：**

    -   main.go 存放的目录。
    -   各种依赖的初始化。如果强依赖的初始化出现错误，那么可以直接 panic ，让服务在启动时就失败。
    -   `func main` 函数中，做服务的初始化、依赖注入。

    **测试：** 不做测试。

    由于依赖比较多，所以会导致函数会很长。依赖注入的过程也比较繁琐，后续可以借助 ` [wire](https://pkg.go.dev/github.com/google/wire) ` 工具直接生成相应的代码。

-   `/internal`

    强制增加 `/internal` package，防止随意引用。1、可以避免循环引用的问题；2、规范调用关系，如果不是我们自己服务的调用的话，那么就应该使用 rpc 的调用方式。

    -   `/server` - http server, grpc server 的定义。

        里面依赖多个 service，每一个 service 算是未来的一个「微服务」能力。比如：UserService、FeedService 等等。

        这里的难点就是要如何定义好各种各样的业务 service 。

        **职责：**

        -   创建 http server，**管理 http server 的生命周期**。 ( 重点 )
        -   ( 类似于 grpc ) 使用 Register 的方式将 server 注入到 api 中，绑定 server 与 router 的关系。

        **测试：** 暂时不需要测试。

    -   `/service`

        **调用关系：** service —> biz 中的 Usercase。

        **职责：**

        1.  **参数转换**，并做简单的参数校验；
        2.  这里面只做编排，不做任何业务逻辑。
        3.  做业务数据的渲染； ( 由于没有 BFF，所以将 BFF 的功能放到这一层做，但是会导致这一层的代码膨胀 ) （ *这个能力暂时存疑。* ）

        -   代码示例

            注意：函数的入参和响应，都依赖于 `/api` 层的 struct。

            ```go
            // ListUserCommunityInfos 这里需要换成 http.ReqListUserCommunityInfo + http.RespListUserCommunityInfo
            func ( us *UserService ) ListUserCommunityInfos ( ctx context.Context, req v1.ListUserCommunityInfosReq ) ( *v1.ListUserCommunityInfosResp, error ) {

                rv, err := us.userUC.ListUserCommunityInfos ( ctx, int ( req.UID ))
                if err != nil {
                    return nil, err
                }

                return &v1.ListUserCommunityInfosResp{
                    CommunityInfos: rv,
                }, nil
            }
            ```

    -   `/biz` - Use Cases

        **职责：**

        -   **包含：**
            -   具体的业务逻辑。
            -   这里面设计对象 ( Domain Object ) 的，可以将业务逻辑放到 Domain Object 中。参考文章： [链接 1](https://blog.csdn.net/abchywabc/article/details/79362975) ， [阿里技术专家详解 DDD 系列 第五讲：聊聊如何避免写流水账代码](https://zhuanlan.zhihu.com/p/366395817)

        -   **不包含**：UI 渲染；数据库或 RPC 框架的具体实现。

        -   这一层按照现在的分层模式，非常独立，不会向上依赖，也不会向下依赖。

        **测试：**

        -   repo 的依赖，由于是 interface 注入，所以直接 mock 的方式。 ( 后续会引入 Go 官方的 [gomock](https://pkg.go.dev/github.com/golang/mock/gomock) )
        -   测试重点为业务逻辑是否符合预期。

        **难点：**

        这一层的难点是，如何定义各种各样的业务用例 ( Usecase ) 。

        **问题：**

        -   这一层下面，是否需要再拆分子目录。

            *个人想法先打平，因为业务太复杂，如果创建子 package 的话，担心导致循环引用的问题。*

        -   Usecase 是否还需要抽象 interface ？还是不在这一层抽象，只抽象 repo 层的 interface。

            个人感觉这一层不需要抽象 interface，暂时只抽象 repo 层的 interface。

        -   类似于 DDD 中的聚合根问题：聚合根非常难以定义。

        -   「聚合根」是否会产生相互依赖？标准来说，应该是按照事件 ( 过于复杂！ ) 。

    -   `/repo`

        **职责：**数据访问层，包括 DB、RPC、缓存等。这里面存放 PO 数据，这些数据就是**简单的表映射**。

    -   `/domain` - ( 可选 ) 所有的 Domain 的定义、Domain 中的接口定义都在这里。这层如果没有的话，那么就可以将所有的 domain object 放在 /biz 层。

-   `/pkg`

    里面定义可以共享出去的工具。由于是可以直接让别人用，这里面的 package 当作基础依赖库使用。既然又是基础依赖库，它里面尽可能的不包含第三方依赖。

## 参考资料

这篇文章的想法借鉴了下列一系列的资料。

1.  基本思想。

  -   [控制反转 IoC](https://zh.wikipedia.org/wiki/%E6%8E%A7%E5%88%B6%E5%8F%8D%E8%BD%AC)

  -   [阿里技术专家详解 DDD 系列 第五讲：聊聊如何避免写流水账代码](https://zhuanlan.zhihu.com/p/366395817) - 这个系列的文章都比较推荐。DDD 没有直接使用，但是在项目中可以参考里面提到的一些理念想法。

  -   [Clean Architecture, 2 years later](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)

  -   [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Bob 大叔的整洁架构的文章。这篇文章也包含在《架构整洁之道》中。

1.  Bob 大叔的 - [《代码整洁之道》](https://book.douban.com/subject/4199741/) ， [《架构整洁之道》](https://book.douban.com/subject/30333919/) 。

1.  Golang 的开源项目，其中包含，
  -   [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - 注意：这里并不是 Go 官方目录结构。
  -   [go-clean-arch](https://github.com/bxcodec/go-clean-arch) - 整洁代码规范
  -   [go-kratos/kratos](https://github.com/go-kratos/kratos) - B 站微服务框架
  -   [go-kratos/beer-shop](https://github.com/go-kratos/beer-shop) - kratos 的示例
  -   [go-kit/kit](https://github.com/go-kit/kit) - go-kit 微服务框架
  -   [go-kit/examples](https://github.com/go-kit/examples) - go-kit 示例代码
