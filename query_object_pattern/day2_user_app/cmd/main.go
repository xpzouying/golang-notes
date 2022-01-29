package main

import (
	"context"
	"fmt"
	"log"

	"user_app/biz"
	"user_app/data"
)

func main() {
	var userUsecase *biz.UserUsecase

	// init
	{
		db, err := data.NewMemSqlite()
		if err != nil {
			log.Fatalf("new sqlite error: %v", err)
		}

		repo, err := data.NewUserRepo(db)
		if err != nil {
			log.Fatalf("new user repository error: %v", err)
		}

		userUsecase, err = biz.NewUserUsecase(repo)
		if err != nil {
			log.Fatalf("new user usecase error: %v", err)
		}
	}

	// create users
	{
		for i := 0; i < 100; i++ {
			var (
				name = fmt.Sprintf("user-%d", i)
				age  = (i % 5) + 15
			)

			user, err := userUsecase.SaveUser(context.Background(), name, age)
			if err != nil {
				log.Printf("save user error: %v", err)
				return
			}

			log.Printf("save user succ: %v", user)
		}
	}

	// take user
	{
		uid := 50

		user, err := userUsecase.GetUserByID(context.Background(), uid)
		if err != nil {
			log.Printf("get uid=%d error", err)
			return
		}

		log.Printf("get uid=%d - %v", uid, user)
	}

	// find users
	{

		users, err := userUsecase.FindAdult(context.Background())
		if err != nil {
			log.Printf("find adult error: %v", err)
			return
		}

		for _, u := range users {
			log.Printf("get adult: %v", u)
		}
	}
}
