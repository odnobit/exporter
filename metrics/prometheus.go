package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
)

var (
	metricsPath = "/metrics"
	bindAddr    = "127.0.0.1:9123"
)

func RunPrometheusServer(c prometheus.Collector) error {
	var links []web.LandingLinks
	var metricsLink = web.LandingLinks{
		Address: metricsPath,
		Text:    "metrics",
	}
	links = append(links, metricsLink)
	landingConfig := web.LandingConfig{
		Name:        "mexample_exporter",
		Description: "mexample description",
		Version:     "v1.0.0-alpha",
		Links:       links,
	}
	landingPageHandler, err := web.NewLandingPage(landingConfig)
	if err != nil {
		return err
	}
	prometheus.MustRegister(c)
	http.Handle("/", landingPageHandler)
	http.Handle(metricsPath, promhttp.Handler())
	log.Printf("starting prometheus server listening on %s", bindAddr)
	log.Println("ERROR", http.ListenAndServe(bindAddr, nil))
	return nil
}
