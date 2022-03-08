package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-metrics-mock/service/count"
	"go-metrics-mock/service/metrics"
)

var BaseCount = count.NewCount()

func MetricsFunc(c *gin.Context) {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		infoVec,
		histogramVec,
	)
	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	h.ServeHTTP(c.Writer, c.Request)
}
func prepareData() {
	//infoVec
	infoVec.WithLabelValues("valueA", "valueB")
	infoVec.WithLabelValues("valueC", "valueD")
	//histogramVec
	histogramVec.WithLabelValues("valueA").Observe(1)
	histogramVec.WithLabelValues("valueA").Observe(6)
	histogramVec.WithLabelValues("valueA").Observe(26)
}
func staticData() {}

var (
	infoVec      = metrics.CreateExampleInfoVec()
	histogramVec = metrics.CreateExampleHistogramVec()
)

func AddCount(c *gin.Context) {
	nowCount := BaseCount.AddCount(1)
	prepareData()
	c.JSON(200, gin.H{
		"nowCount": nowCount,
	})
}
