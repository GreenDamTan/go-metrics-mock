package main

import (
	"github.com/gin-gonic/gin"
	"go-metrics-mock/router"
)

func main() {
	r := gin.Default()
	router.RouterInit(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
