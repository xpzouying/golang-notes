package biz

import (
	"context"
	"errors"

	"user_app/queryobject"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepo interface {
	// --- 查询条件 ---
	//
	// 查询条件写到接口里面是为了避免 /biz 层，依赖/data层的东西。

	WithUid(uid int) queryobject.Option
	WithName(name string) queryobject.Option
	OlderThan(age int) queryobject.Option

	// --- 方法 ---
	SaveUser(ctx context.Context, user *User) (*User, error)
	// GetUserByID(ctx context.Context, id int) (*User, error)

	GetUser(ctx context.Context, options ...queryobject.Option) (*User, error)

	FindUsers(ctx context.Context, options ...queryobject.Option) ([]*User, error)
}

type UserUsecase struct {
	userRepo UserRepo
}

func NewUserUsecase(userRepo UserRepo) (*UserUsecase, error) {
	if userRepo == nil {
		return nil, errors.New("user repository is nil")
	}

	return &UserUsecase{userRepo}, nil
}

func (uc *UserUsecase) SaveUser(ctx context.Context, name string, age int) (*User, error) {

	return uc.userRepo.SaveUser(ctx, &User{
		Name: name,
		Age:  age,
	})
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int) (*User, error) {

	return uc.userRepo.GetUser(ctx, uc.userRepo.WithUid(id))
}

func (uc *UserUsecase) FindAdult(ctx context.Context) ([]*User, error) {
	return uc.userRepo.FindUsers(ctx, uc.userRepo.OlderThan(17))

}
