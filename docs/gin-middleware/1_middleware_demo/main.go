package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type M map[string]interface{}

// HandleHiInternal is internal url.
// Panic if request without token or with invalid token
func HandleHiInternal(c *gin.Context) {
	var req M
	if err := c.Bind(&req); err != nil {
		panic(err)
	}

	log.Printf("request in HandleHiInternal: %+v", req)

	time.Sleep(100 * time.Millisecond)

	c.JSON(http.StatusOK, M{"result": true})
}

func HandleHi(c *gin.Context) {

	var req M
	if err := c.Bind(&req); err != nil {
		panic(err)
	}

	log.Printf("request in HandleHi: %+v", req)

	time.Sleep(100 * time.Millisecond)

	c.JSON(http.StatusOK, M{"result": true})
}

// Protector is middleware for protect us
func Protector() gin.HandlerFunc {

	return func(c *gin.Context) {

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		data, err := ioutil.ReadAll(tee)
		if err != nil {
			panic(err)
		}
		c.Request.Body = ioutil.NopCloser(&buf)

		var req M
		if err := json.Unmarshal(data, &req); err != nil {
			panic(err)
		}

		token, ok := req["token"]
		if !ok {
			panic("request without token")
		}

		if token != "123456" {
			panic("request with invalid token")
		}

		log.Printf("request in Middleware: %+v", req)

		c.Next()
	}
}

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	r.Use(gin.Logger())

	internalURL := r.Group("/internal")
	internalURL.Use(Protector())

	r.POST("/hi", HandleHi)
	internalURL.POST("/hi", HandleHiInternal)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
