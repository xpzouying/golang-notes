package data

import (
	"context"
	"errors"

	"user_app/biz"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   int `gorm:"primaryKey"`
	Name string
	Age  uint8

	gorm.Model
}

func (User) Tablename() string {
	return "user"
}

type UserRepo struct {
	db *gorm.DB
}

func NewMemSqlite() (*gorm.DB, error) {
	return gorm.Open(
		sqlite.Open("file::memory:?cache=shared"),
		&gorm.Config{},
	)
}

func NewUserRepo(db *gorm.DB) (*User, error) {

	if db == nil {
		return nil, errors.New("connection is nil")
	}

	// Auto Migrate
	db.AutoMigrate(&User{})

	return &UserRepo{db}, nil
}

func (us *UserRepo) SaveUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{
		Name: user.Name,
		Age:  uint8(user.Age),
	}
	if err := us.db.Create(&u).Error; err != nil {
		return nil, err
	}

	return &biz.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  int(u.Age),
	}, nil
}

func (us *UserRepo) GetUserByID(ctx context.Context, uid int) (*biz.User, error) {

	var user User
	if err := us.db.Where("id = ?", uid).Take(&user).Error; err != nil {
		return nil, err
	}

	return &biz.User{
		ID:   user.ID,
		Age:  int(user.Age),
		Name: user.Name,
	}, nil
}
