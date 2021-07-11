package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	f, err := os.OpenFile("/tmp/go_logs/example.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	panicError(err)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
}

func main() {

	log.WithFields(log.Fields{
		"event":   "create_profile",
		"user_id": 10,
	}).Info("This is an info message.")

	log.WithFields(log.Fields{
		"event":   "delete_profile",
		"user_id": 11,
	}).Warn("This is a warning message.")

	log.WithFields(log.Fields{
		"event":   "edit_profile",
		"user_id": 13,
		"package": "main",
	}).Fatal("This is a critical message.")
}
