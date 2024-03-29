# Best Practices - Testing

在这里总结了一些单元测试的最佳实践。

## 规范

- 【强制】好的单元测试必须满足`AIR原则`。
    - A - Automatic（自动化）
    - I - Independent（独立性）
    - R - Repeatable（可重复）

- 【强制】单元测试必须是自动化的，不能依赖人工校验。

    禁止通过类似于`log.Printf`的日志输出人肉校验，必须通过`assert`之类的进行校验。
    
- 【强制】单元测试是可以被重复执行的，不受外界环境的影响，当前运行以及未来运行、自己环境运行和他人环境都能得到相同的结果。
    
    说明：1.单元测试如果强依赖于数据库、中间件、网络，往往会导致单元测试不稳定；2.由于测试数据库中的数据有可能被他人修改，也往往会导致单元测试依赖的数据发生变化，从而影响单元测试。
    
- 【强制】必须保证测试的独立性。为了保证单元测试稳定可靠且便于维护，单元测试用例之间决不能互相调用，也不能依赖执行的先后次序。一个测试不能依赖于另外一个测试的结果。

  反例：method2 需要依赖 method1 的执行，将执行结果作为 method2 的输入。

- 【强制】测试用例没有副作用。执行前后，环境一致。

- 【强制】单元测试的粒度尽可能简单。每个单元测试仅仅对一个函数进行校验。

- 【强制】业务修改后，导致对应的单元测试发生错误，必须修复。

- 【强制】Golang测试要求：
    - 测试文件命名：xxx_test.go。
    - 测试方法：TestXxxx开头，以`t *testing.T`为函数参数。
    - 测试文件和被测试的文件必须在一个package里面。

- 【推荐】优先编写核心组件和逻辑模块的测试用例。

## 总结

**写一个测试时，应该问问自己以下几个问题，如果你回答不是，拿你写的就不是单元测试**

1. 它是可以重复执行的
2. 它是不依赖第三方包括数据库依赖的
3. 任何人都可以一键执行的
4. 它是完全隔离的，不受其他测试影响的
5. 他的结果是稳定的，无论时间空间结果都是一样的
6. 他的运行速度是非常快的

**写一个单元测试后，你应该问问这个测试以下几个问题**

1. 我两周前写的一个测试，今天或几年后还能运行并获得相同的结果吗？
2. 我两个月前写的测试，团队里其他成员还能运行他并能获取结果吗？
3. 我写了一个测试，其他成员运行他需要了解基础设施如何设置吗？例如数据库/缓存/配置/dbconfig/config.....
4. 我几分钟内能跑完所有测试吗？
5. 我能几分钟写出一个基本测试吗？

如果你说不能，一样也不是单元测试，你可能需要重构你的代码和你的测试代码。你写的这个其实是个集成测试，集成测试不是不重要，而是应该尽量少。