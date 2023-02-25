package custom_http_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"k8s.io/utils/env"
	"net/http"
	"net/http/httputil"
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
	request, err := httputil.DumpRequest(r, false)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write(request)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	v := env.GetString("study.golang.version", "not found")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("%s:%s", "version", v)))
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	sIp := r.Header.Get("X-Forwarded-For")
	fmt.Println(fmt.Sprintf("sourceIp: %s, response code: %d", sIp, 200))
	w.WriteHeader(200)
}
