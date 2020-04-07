package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/multicharts/go-zookeeper/zk"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/tink-ab/{{ cookiecutter.name }}/web"
	"github.com/tink-ab/tink-go-libraries/discovery"
	"github.com/tink-ab/tink-go-libraries/discovery/announce"
	"github.com/tink-ab/tink-go-libraries/net"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server.",
	Long:  `Start the microservice and announce that it's available to Zookeeper for discovery.`,
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
	initFlags()
}

func initFlags() {
	registerStringFlag(serveCmd, "listen", ":8080", "The address and port to listen on.")
	registerStringFlag(serveCmd, "prometheus-listen", ":9130", "The address and port to listen on.")
	registerStringFlag(serveCmd, "zookeeper-hosts", "localhost", "Zookeeper host. Can be comma separated for multiple hosts.")
	registerDurationFlag(serveCmd, "shutdown-timeout", 5*time.Minute, "Time to wait for requests to complete before shutting down")
}

func serve(cmd *cobra.Command, args []string) {
	listen := getString("listen", true)
	shutdownTimeout := getDuration("shutdown-timeout", true)
	zookeeperHosts := getString("zookeeper-hosts", true)

	// Setup web stuff.
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(&webLogWriter{Logger: log.New(os.Stdout, "", 0)}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", (&web.RouteBuilder{}).Routes())

	zkConn, zkEvents, node, err := startZookeeperConnection(listen, zookeeperHosts)
	if err != nil {
		log.Fatalln(err)
	}

	serviceAnnouncer, err := announce.Start(zkConn, zkEvents, *node)
	if err != nil {
		log.Fatalln(err)
	}

	// Spin up Prometheus metric exporter in separate thread
	go startPrometheusExporter()

	server := http.Server{Addr: listen, Handler: r}
	go startWebServer(&server, listen)

	// notifies when the process received termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	if err = serviceAnnouncer.Stop(); err != nil {
		log.Printf("Graceful shutdown of zookeeper announcer failed: %v", err)
	}

	log.Printf("Interrupt received, waiting %v to shut down", shutdownTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown of HTTP server failed: %v", err)
	}
	log.Printf("HTTP server shutdown successfully")
}

func startZookeeperConnection(listen, zookeeperHost string) (*zk.Conn, <-chan zk.Event, *discovery.ServiceEntity, error) {
	addressAndPort := strings.Split(listen, ":")
	port64, err := strconv.Atoi(addressAndPort[1])
	if err != nil {
		return nil, nil, nil, err
	}
	port := uint(port64)

	eIP, err := net.ExternalIP()
	if err != nil {
		return nil, nil, nil, err
	}

	zkSessionTimeout := time.Minute
	zkConn, zkEvents, err := zk.Connect([]string{zookeeperHost}, zkSessionTimeout)

	if err != nil {
		return nil, nil, nil, err
	}

	announcedNode := discovery.ServiceEntity{
		Address: eIP.String(),
		Name:    "categorization",
		Port:    &port,
	}

	return zkConn, zkEvents, &announcedNode, nil
}

func startPrometheusExporter() {
	prometheusListen := getString("prometheus-listen", true)

	promHandler := http.NewServeMux()
	promHandler.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting Prometheus metric exporter on %s", prometheusListen)
	log.Fatalln(http.ListenAndServe(prometheusListen, promHandler))

}

func startWebServer(server *http.Server, listen string) {
	log.Printf("Starting HTTP server on %s", listen)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server error: %v", err)
	}
}
