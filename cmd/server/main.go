package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {

	server = gin.Default()

}

func main() {

	log.Fatal(server.Run(":8090"))
}
