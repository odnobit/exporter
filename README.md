#### Prometheus simple app exporter

This project provides a simple example of a Prometheus exporter written in Golang. The main application runs as a daemon, writing metric values to memory and exposing them for Prometheus to scrape. The metrics structure implements the prometheus.Collector interface to collect metrics from memory.
