package metrics

import (
	"time"

	"github.com/odnobit/exporter/storage"

	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "mexample"

var startTime time.Time

func init() {
	startTime = time.Now()
}

func getAppUptimeSec() float64 {
	currentTime := time.Now()
	uptime := currentTime.Sub(startTime)
	return uptime.Seconds()
}

type Metrics struct {
	AppUptimeSec           *prometheus.Desc
	MessagesTotalCounter   *prometheus.Desc
	MessagesFailedCounter  *prometheus.Desc
	MessagesSuccessCounter *prometheus.Desc
}

func NewMetrics() *Metrics {
	return &Metrics{
		AppUptimeSec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "app", "uptime"),
			"application uptime in seconds",
			nil,
			nil,
		),
		MessagesTotalCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "messages", "total"),
			"total messages counter",
			nil,
			nil,
		),
		MessagesFailedCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "messages", "failed"),
			"failed messages counter",
			nil,
			nil,
		),
		MessagesSuccessCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "messages", "success"),
			"success messages counter",
			nil,
			nil,
		),
	}
}

func (collector *Metrics) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(collector, ch)
}

func (collector *Metrics) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(
		collector.AppUptimeSec,
		prometheus.GaugeValue,
		getAppUptimeSec(),
	)
	ch <- prometheus.MustNewConstMetric(
		collector.MessagesTotalCounter,
		prometheus.CounterValue,
		storage.Memory.GetTotalMessages(),
	)
	ch <- prometheus.MustNewConstMetric(
		collector.MessagesFailedCounter,
		prometheus.CounterValue,
		storage.Memory.GetFailedMessages(),
	)
	ch <- prometheus.MustNewConstMetric(
		collector.MessagesSuccessCounter,
		prometheus.CounterValue,
		storage.Memory.GetSuccessMessages(),
	)
}
