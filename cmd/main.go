package main

import (
	"flag"
	"fmt"
	"net/http"
	"o11y_e2e/internal/server"
	"os"

	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultLogLevel      = "info"
	defaultListenAddress = "0.0.0.0:9927"
	defaultMetricsPath   = "/metrics"
)

var (
	listenAddress = flag.String("web.listen-address", defaultListenAddress, "Address to listen on for telemetry")
	showVersion   = flag.Bool("version", false, "show version information")
	logLevel      = flag.String("log.level", defaultLogLevel,
		"Minimum Log level [info, debug, warn, error]")

	// Build info for o11y_e2e itself, will be populated by linker during build
	Version   string
	BuildDate string
	Commit    string

	versionInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "o11y_e2e_version_info",
			Help: "o11y_e2e build information",
		},
		[]string{"version", "commit", "builddate"},
	)
)

func printVersion() {
	fmt.Printf("o11y_e2e\n")
	fmt.Printf("Version:   %s\n", Version)
	fmt.Printf("BuildDate: %s\n", BuildDate)
	fmt.Printf("Commit:    %s\n", Commit)
	fmt.Printf("o11y_e2e pipeline measurement tool\n")
}

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	versionInfo.WithLabelValues(Version, Commit, BuildDate).Set(1)
	prometheus.MustRegister(versionInfo)

	slogLogLevel := new(slog.LevelVar)
	switch *logLevel {
	case "debug":
		slogLogLevel.Set(-4)
	case "warn":
		slogLogLevel.Set(4)
	case "error":
		slogLogLevel.Set(8)
	default:
		slogLogLevel.Set(0)
	}

	opts := &slog.HandlerOptions{
		Level: slogLogLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)

	slog.SetDefault(slog.New(handler))

	http.Handle(defaultMetricsPath, promhttp.Handler())
	http.Handle("/", server.SetupServer())

	slog.Info("Starting server", "address", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		slog.Error("Failed to start the server", "error", err)
	}
}
