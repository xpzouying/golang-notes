package main

import (
	"context"
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

	{
		user, err := userUsecase.SaveUser(context.Background(), "zouying", 18)
		if err != nil {
			log.Printf("save user error: %v", err)
			return
		}

		log.Printf("save user succ: %v", user)
	}
}
