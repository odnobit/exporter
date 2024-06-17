#### Prometheus simple app exporter

This project provides a simple example of a Prometheus exporter written in Golang. The main application runs as a daemon, writing metric values to memory and exposing them for Prometheus to scrape. The metrics structure implements the prometheus.Collector interface to collect metrics from memory.

metrics output
```
curl -s http://127.0.0.1:9123/metrics | grep ^mexample

mexample_app_uptime 102.201255334
mexample_messages_failed 7
mexample_messages_success 7
mexample_messages_total 14
```
