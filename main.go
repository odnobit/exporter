package main

import (
	"log"

	"github.com/odnobit/exporter/app"
	"github.com/odnobit/exporter/metrics"
)

func main() {
	m := metrics.NewMetrics()
	agent := app.NewAppDaemon()
	go func() {
		err := metrics.RunPrometheusServer(m)
		if err != nil {
			log.Panicf("prometheus server has failed: %s", err)
		}
	}()
	go agent.Run()
	select {}
}
