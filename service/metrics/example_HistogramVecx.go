package metrics

import "github.com/prometheus/client_golang/prometheus"

// CreateExampleHistogramVec 创建一个*prometheus.HistogramVec
func CreateExampleHistogramVec() *prometheus.HistogramVec {
	HistogramVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "z_HistogramVec",
			Help:    "统计低于X值的数值个数",
			Buckets: prometheus.ExponentialBuckets(1, 5, 3),
		},
		[]string{"LabelsA"},
	)
	return HistogramVec
}
