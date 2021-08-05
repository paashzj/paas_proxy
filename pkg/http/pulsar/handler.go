package pulsar

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gorilla/mux"
	"github.com/paashzj/paas_proxy/pkg/config"
	"github.com/paashzj/paas_proxy/pkg/http/module"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var client pulsar.Client

func init() {
	var err error
	client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:               fmt.Sprintf("pulsar://%s:%d", config.PulsarHost, config.PulsarPort),
		ConnectionTimeout: time.Duration(config.PulsarConnectionTimeoutSeconds) * time.Second,
		OperationTimeout:  time.Duration(config.PulsarOperationTimeoutSeconds) * time.Second,
	})
	if err != nil {
		logrus.Error("create client err ", err)
		os.Exit(1)
	}
}

func ProduceHandler(w http.ResponseWriter, r *http.Request) {
	// vars contain the path variable
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	namespace := vars["namespace"]
	topic := vars["topic"]
	var req module.ProduceMsgReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := Produce(tenant, namespace, topic, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := module.ProduceMsgResp{Cost: cost}
	payload, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	logrus.Error("write payload failed ", err)
}

func Produce(tenant, namespace, topic string, req module.ProduceMsgReq) (int64, error) {
	startTime := time.Now().Unix()
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: fmt.Sprintf("persistent://%s/%s/%s", tenant, namespace, topic),
	})
	if err != nil {
		logrus.Error("Failed to create producer ", err)
		return 0, err
	}
	defer producer.Close()
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(req.Msg),
	})
	if err != nil {
		logrus.Error("Failed to publish message ", err)
		return 0, err
	}
	return time.Now().Unix() - startTime, nil
}
