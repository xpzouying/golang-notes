# Query Object Pattern

## 前言

在开发过程中，我们常常需要从数据库中查询实际的数据，因此我们常常会讲我们的数据库依赖放到我们的业务对象中，

比如：

```go
type UserUsecase struct {
	userRepo UserRepo
}

func (uc *UserUsecase) SaveUser(ctx context.Context, name string, age int) (*User, error) {

	return uc.userRepo.SaveUser(ctx, &User{
		Name: name,
		Age:  age,
	})
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int) (*User, error) {

	return uc.userRepo.GetUserByID(ctx, id)
}
```

在上面代码中，我们定义了个`UserUsecase`，封装我们的User相关的业务服务，目前我们提供了2个业务功能：

- SaveUser - 创建一个User对象。
- GetUserByID - 查询对应uid的记录。


在`UserUsecase`中，我们依赖的是接口，而不是具体的实现。其中`UserRepo`的接口约定了2个方法：

```go
type UserRepo interface {
	SaveUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
}
```

在`main()`启动的过程中，我们将`Sqlite`对应的Repo实现注入到`UserUsecase`中。

实际的`biz.UserRepo`实现是`data.UserRepo`对象，所以`data.UserRepo`中，我们也需要实现接口约定的所有方法。


## DAY 1 - 基本实现

参见代码：`day1_user_app`中，相关代码。

**运行**

```bash
go run cmd/main.go
```


**问题**

随着业务的发展，之前仅仅通过用户uid进行查询已经满足不了业务的需求了，
我们需要增加更多的查询条件，比如：

- GetUserByName - 查询名字为XXX的用户

- FindUserByAge - 查询年龄18岁的用户

- FindUserOlder - 查询大于18岁的用户

- XXX等等


此时，如果我们对每一种查询条件都增加一个方法，肯定会使得`biz.UserRepo interface`和`data.UserRepo struct`同时爆炸，并且极为不灵活。

那么，如何写出更好的SQL查询？


## 参考资料

- [Go 工程化(十一) 如何优雅的写出 repo 层代码](https://lailin.xyz/post/graceful-repo-code.html)
