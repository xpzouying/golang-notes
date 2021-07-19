package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	log.Printf("error: %v", err)

	log.Printf("errors.Is: %v", errors.Is(err, sql.ErrNoRows))
	log.Printf("errors.Cause: %v", errors.Cause(err))
}
