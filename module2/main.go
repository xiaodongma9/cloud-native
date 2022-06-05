package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.Setenv("version", "0.1")
	http.HandleFunc("/response-header/request", setResHeaderWithReq)
	http.HandleFunc("/response-header/sysenv", setResHeaderWithSysVersion)
	http.HandleFunc("/info-logging", logInfo)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
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
	ip := r.RemoteAddr
	log.Printf("client IP: %s\n", ip)
	log.Printf("httpCode %d\n", http.StatusOK)
	w.Write([]byte(ip))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
