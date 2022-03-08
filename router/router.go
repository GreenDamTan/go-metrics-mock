package router

import (
	"github.com/gin-gonic/gin"
	"go-metrics-mock/api/metrics"
	"net/http"
)

// RouterInit 路由
func RouterInit(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "metrics 演示工具")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/metrics", metrics.MetricsFunc)
	BaseCount := r.Group("/BaseCount")
	{
		BaseCount.GET("/AddCount", metrics.AddCount)
	}
}
