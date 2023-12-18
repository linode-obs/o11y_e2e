package server

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"log/slog"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupServer() http.Handler {

	const (
		defaultHTML = `<html>
				<head><title>o11y_e2e</title></head>
				<body>
				<h1>o11y_e2e</h1>
				<p><a href='%s'>Metrics</a></p>
				</body>
				</html>`
		defaultMetricsPath = "/metrics"
	)

	mux := http.NewServeMux()

	mux.Handle(defaultMetricsPath, promhttp.Handler())

	// TODO: start internal metrics interval

	// for non-standard web servers, need to register handlers
	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf(defaultHTML, defaultMetricsPath)
		_, err := w.Write([]byte(response))
		if err != nil {
			slog.Error("Failed to write main page response", "error", err)
		}
	})

	return mux
}
