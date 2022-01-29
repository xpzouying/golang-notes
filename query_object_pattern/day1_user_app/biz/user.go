package biz

import (
	"context"
	"errors"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type UserRepo interface {
	SaveUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
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

	return uc.userRepo.GetUserByID(ctx, id)
}
