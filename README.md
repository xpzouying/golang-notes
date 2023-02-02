# golang-notes

这里记录着我的 Golang 笔记和思考。

另外， [Go 常见的错误汇总](./CommonError.md) 汇总常见的错误问题。


# List Table

- [Go 打包前端代码的示例](./go_embed_front_source/) - 使用 Go embed 特性直接在二进制中打包发布前端服务的示例。


# Reading List

-   [Go 语言之父谈 Go 编程语言与环境](https://mp.weixin.qq.com/s?__biz=MzIyNzM0MDk0Mg==&mid=2247490227&idx=1&sn=620d9ab5f06c64852a141e43abf44fef&cur_album_id=1509674724665163776#wechat_redirect) - Rob Pike 介绍 Go 的特性。这篇文章可以看到 Go 语言与其他语言的区别所在。

-   [2021 年 Go 语言盘点：厉兵秣马强技能，蓄势待发新征程](https://tonybai.com/2022/01/16/the-2021-review-of-go-programming-language/) - 汇总 Go 2021 年事件盘点。在今年发布了 2 个版本：Go1.16、Go1.17，另外还发布了 Go1.18 beta 版本，在这里尝试加入泛型的技术方案。在 2022 年，泛型将被正式加入到 Go 中，从 [Go 语言之父谈 Go 编程语言与环境](https://mp.weixin.qq.com/s?__biz=MzIyNzM0MDk0Mg==&mid=2247490227&idx=1&sn=620d9ab5f06c64852a141e43abf44fef&cur_album_id=1509674724665163776#wechat_redirect) 的文章中可以看到，对于泛型加入后的社区反映，`Rob Pike` 也不确定反响如何，让我们 2022 年拭目以待。

-   [MySQL 游标分页与传统分页](https://github.com/x1ah/Blog/issues/15) - 业务开发中，经常会遇到的分页问题。推荐用游标分页的方式，而不是 `offset+limit` 的方式。

-   [DDD 的学习书单](https://zhuanlan.zhihu.com/p/138884686) - 推荐了 2 本 Martin Flowler 的另外 2 本书：《企业应用架构模式》 & 《分析模式》。

-   [你也能写个 Shadowsocks](https://github.com/gwuhaolin/blog/issues/12) - 介绍制作Proxy的原理。

## 记录

-   [HTTP 服务如何优雅退出和重启](./graceful_shutdown_and_restart/README.md)

-   读书：[《代码整洁之道》](https://book.douban.com/subject/34986245/)

-   读书：[《代码整洁之道 - 程序员的职业素养》](https://book.douban.com/subject/26919457/) - 2020年2月13日完成。
    Bob大叔在本书中给出的建议覆盖了程序员工作方式、工作态度、学习方式、到编码技巧的全方位最佳实践。推荐各位程序员，特别是团队新人阅读。

-   读书：[《架构整洁之道》](https://book.douban.com/subject/30333919/)


## Todos

-   Raft 的实现

  -   https://github.com/eliben/raft

  -   https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/

-   WAL 的实现

-   《Clean Architechture》读书笔记

  -   [《clean architecture》第一部分](https://mp.weixin.qq.com/s?__biz=MzI4NDM0MzIyMg==&mid=2247489322&idx=1&sn=84f956b3c50ea95a544dbd1355e3c266&scene=21#wechat_redirect)

  -   [Go 工程化 ( 一 ) 架构整洁之道阅读笔记](https://lailin.xyz/post/go-training-week4-clean-arch.html)

  -   [clean architecture ( 上 )](https://xargin.com/clean-architecture-1/)

    这个博客有不少可以参考的文章。

-   [Go 语言设计模式](https://github.com/senghoo/golang-design-pattern)

  -   [深入设计模式](https://refactoringguru.cn/design-patterns/builder)

-   规范相关

  -   [Microsoft](https://github.com/microsoft/code-with-engineering-playbook/)

## 如何使用

1.  打开 [Issues](https://github.com/xpzouying/golang-notes/issues) 可以查看所有的笔记。可以搜索相关问题，或使用 `Labels` 查看某类问题。

2.  若涉及示例代码，则保存在 [docs](https://github.com/xpzouying/golang-notes/tree/master/docs) 目录中。

3.  讨论。所有的 Issue 暂时不会关闭，可以在相应的 Issue 下面进行讨论。

## 更新提醒

-   不要使用 `Fork`

-   点击 `Star`，保持 `Watching` 状态

## 进展 / 状态

建立该仓库的初衷是记录我在学习、研究 Golang 自身语言的过程中，记录其中的分析笔记。

原本，准备完全围绕 `Golang` 语言自身展开，但是最近想提升自身的架构能力，所以想研究更高、更广层面的内容，比如现在很火的 docker、k8s，还有一些中间件，比如对象存储、消息队列，或者一些协议，比如 raft 等等。

因此，接下来的笔记不光是记录 Golang 语言自身，还会涉及到 Go 生态的其他内容。

## 感谢

-   [7days-golang](https://github.com/geektutu/7days-golang) - 该仓库的学习榜样 [geektutu](https://geektutu.com/post/gee.html) 。另外，也感谢他分享的 ` 微习惯 ` 养成思维。
