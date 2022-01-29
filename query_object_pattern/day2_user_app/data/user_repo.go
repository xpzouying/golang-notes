package data

import (
	"context"
	"errors"
	"fmt"

	"user_app/biz"
	"user_app/queryobject"

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

func NewUserRepo(db *gorm.DB) (*UserRepo, error) {

	if db == nil {
		return nil, errors.New("connection is nil")
	}

	// Auto Migrate
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, errors.New("migrate table error")
	}

	return &UserRepo{db}, nil
}

func (us *UserRepo) WithUid(uid int) queryobject.Option {
	return us.where(queryobject.OpEqual, "id", uid)
}

func (us *UserRepo) WithName(name string) queryobject.Option {

	return us.where(queryobject.OpEqual, "name", name)
}

func (us *UserRepo) OlderThan(age int) queryobject.Option {

	return us.where(queryobject.OpGreater, "age", age)
}

func (us *UserRepo) where(op queryobject.Operator, field queryobject.Field, value interface{}) queryobject.Option {

	return func(src interface{}) interface{} {

		impl, ok := src.(*gorm.DB)
		if !ok {
			return src
		}

		return impl.Where(
			fmt.Sprintf("%s %v ?", field, op),
			value,
		)
	}
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

func (us *UserRepo) GetUser(ctx context.Context, options ...queryobject.Option) (*biz.User, error) {

	db, err := us.beginQuery(ctx, options...)
	if err != nil {
		return nil, err
	}

	var rv User
	if err := db.Take(&rv).Error; err != nil {
		return nil, err
	}

	return &biz.User{
		ID:   rv.ID,
		Age:  int(rv.Age),
		Name: rv.Name,
	}, nil
}

func (us *UserRepo) FindUsers(ctx context.Context, options ...queryobject.Option) ([]*biz.User, error) {

	db, err := us.beginQuery(ctx, options...)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return us.assembleBizUsers(users), nil
}

func (us *UserRepo) beginQuery(ctx context.Context, options ...queryobject.Option) (*gorm.DB, error) {

	var db interface{} = us.db.WithContext(ctx)

	for _, o := range options {
		db = o(db)
	}

	// return db
	rv, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("cannt begin query")
	}

	return rv, nil
}

func (us *UserRepo) assembleBizUsers(users []*User) []*biz.User {
	rv := make([]*biz.User, 0, len(users))
	for _, u := range users {
		rv = append(rv, &biz.User{
			ID:   u.ID,
			Name: u.Name,
			Age:  int(u.Age),
		})
	}

	return rv
}
