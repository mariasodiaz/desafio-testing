package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mariasodiaz/desafio-testing/cmd/router"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}
