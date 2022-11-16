package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const healthPath = "/health"
const keyServerAddr = "serverAddr"
const hostEnvVar = "HOST"
const defaultHost = "localhost"
const portEnvVar = "PORT"
const defaultPort = "8080"
const cfgEnvVar = "CONFIG_FILE"
const defaultCfg = "config.json"

type CityTZ struct {
	City string `json:"city"`
	TZ   string `json:"tz"`
}

var config []CityTZ

func loadConfig() {
	cfg := os.Getenv(cfgEnvVar)
	if len(cfg) == 0 {
		cfg = defaultCfg
	}

	file, err := os.ReadFile(cfg)

	if err != nil {
		log.Fatal("error opening config file: ", err)
	}

	err = json.Unmarshal([]byte(file), &config)

	if err != nil {
		log.Fatal("error unmarshalling config file: ", err)
	}
}

func timeRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Debug("got root path / request for address ", ctx.Value(keyServerAddr))

	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)

	content := "<html><head><title>Local times</title></head><body>"
	content += "<table><tr><th>City</th><th>Time</th></tr>"

	for i := 0; i < len(config); i++ {
		loc, err := time.LoadLocation(config[i].TZ)

		if err != nil {
			log.Warning(fmt.Sprintf("Error getting time for %s: ", config[i].City), err)
			continue
		}

		time := time.Now().In(loc)

		content += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", config[i].City, time)
	}

	content += "</body></html>"

	io.WriteString(w, content)
}

func healthRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Debug("got /health request for endpoint ", ctx.Value(keyServerAddr))

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	content := `{"status_code": 200}`

	io.WriteString(w, content)
}

func TimeServer(quit <-chan bool) {
	loadConfig()

	host := os.Getenv(hostEnvVar)
	if len(host) == 0 {
		host = defaultHost
	}

	port := os.Getenv(portEnvVar)
	if len(port) == 0 {
		port = defaultPort
	}

	listenAddr := fmt.Sprintf("%s:%s", host, port)

	mux := http.NewServeMux()
	server := http.Server{Addr: listenAddr, Handler: mux}
	mux.HandleFunc("/", timeRequest)
	mux.HandleFunc(healthPath, healthRequest)

	go func() {
		log.Info("starting http server listening on ", listenAddr)

		err := server.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			log.Info("server gracefully shut down\n")
		} else if err != nil {
			log.Fatal("error starting server: ", err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-signalChannel:
			switch sig {
			case os.Interrupt:
				log.Info("SIGINT received, shutting down")
				server.Shutdown(context.Background())
				return
			case syscall.SIGTERM:
				log.Info("SIGTERM received, shutting down")
				server.Shutdown(context.Background())
				return
			}

		case <-quit:
			log.Info("quit received, shutting down")
			server.Shutdown(context.Background())
			return
		}
	}
}

func main() {
	quit := make(chan bool)

	TimeServer(quit)
}
