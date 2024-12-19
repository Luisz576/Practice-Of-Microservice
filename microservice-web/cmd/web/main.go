package main

import (
	"fmt"
	"microservice-web/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFS("/", http.Dir("website"))

	router.Run(fmt.Sprintf(":%d", config.API_PORT))
}
