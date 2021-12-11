package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

type ExecutionTime struct {
	histogram *prometheus.HistogramVec
	start     time.Time
	stop      time.Time
}

const Namespace = "httpserver"

var (
	functionLatency = CreateExecutionTimeMetric(Namespace,
		"Time spent.")
)

func NewTimer() *ExecutionTime {
	return NewExecutionTimer(functionLatency)
}

func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		fmt.Println(err)
	}
}

func NewExecutionTimer(histogram *prometheus.HistogramVec) *ExecutionTime {
	now := time.Now()
	return &ExecutionTime{
		histogram: histogram,
		start:     now,
		stop:      now,
	}
}

func (t *ExecutionTime) ObserveTotal() {
	t.histogram.WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_used_second",
			Help:      help,
			// prometheus.ExponentialBuckets 创建一个 bucket
			// 以 start(0.001) 开始，共 count(15) 个 桶，后一个桶 是前一个 桶 * factor(2)
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
		},
		[]string{"step"},
	)
}
