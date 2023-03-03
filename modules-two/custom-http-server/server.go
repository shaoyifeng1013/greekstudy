package custom_http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

type CustomHttpServer struct {
	s *http.Server
}

func NewCustomHttpServer() *CustomHttpServer {
	router := mux.NewRouter()

	router.HandleFunc("/header", headerHandler)
	router.HandleFunc("/version", versionHandler)
	router.HandleFunc("/print", printHandler)
	router.HandleFunc("/localhost/healthz", healthCheckHandler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &CustomHttpServer{s}
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%s,%v", k, v)
		w.Header().Set(k, fmt.Sprintf("%v", v))
	}
	w.WriteHeader(200)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	v := os.Getenv("study.golang.version")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%s:%s", "version", v)))
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	sIp := r.Header.Get("X-Forwarded-For")
	fmt.Println(fmt.Sprintf("sourceIp: %s, response code: %d", sIp, 200))
	w.WriteHeader(200)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
