package main

import (
	_ "github.com/paashzj/paas_proxy/pkg/log"
	"github.com/sirupsen/logrus"
)
import (
	"github.com/paashzj/paas_proxy/pkg/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exitChan := make(chan os.Signal, 1)
	// catch SIGTERM or SIGINTERRUPT
	signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)
	go http.Start()
	logrus.Info("start over")
	sig := <-exitChan
	logrus.Printf("Caught SIGTERM %v", sig)
}
