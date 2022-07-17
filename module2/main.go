package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	os.Setenv("version", "0.1")
	go trapShutdownSignal()
	http.HandleFunc("/response-header/request", setResHeaderWithReq)
	http.HandleFunc("/response-header/sysenv", setResHeaderWithSysVersion)
	http.HandleFunc("/info-logging", logInfo)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func trapShutdownSignal() {
	log.Println("watching for termination signals")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	// when we get a signal, flip the global ShuttingDown flag
	sig := <-sigChan
	log.Println("got signal:", sig)

	// wait for the liveness checks to fail and kubernetes to reconfigure
	log.Println("graceful shutdown has begun")
	time.Sleep(time.Second * 20)
	// sleep while the cluster removes this instance from incoming service traffic
	log.Println("exiting clean due to shutdown signal")
	os.Exit(0)
}

func setResHeaderWithReq(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("key: %s values: %v\n", k, v)
		w.Header().Set(k, strings.Join(v[:], ","))
	}
}

func setResHeaderWithSysVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("version", os.Getenv("version"))
}

func logInfo(w http.ResponseWriter, r *http.Request) {
	ip := r.Header.Get("X-REAL-IP")
	log.Printf("client IP: %s\n", ip)
	log.Printf("httpCode %d\n", http.StatusOK)
	w.Write([]byte(ip))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
