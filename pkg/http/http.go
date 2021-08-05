package http

import (
	"github.com/gorilla/mux"
	"github.com/paashzj/paas_proxy/pkg/http/pulsar"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/pulsar/tenants/{tenant}/namespaces/{namespace}/topics/{topic}/produce", pulsar.ProduceHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:20001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	logrus.Error("http start failed", err)
	os.Exit(1)
}
