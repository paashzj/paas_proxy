package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	pulsarEnvHost                     = "PULSAR_HOST"
	pulsarEnvPort                     = "PULSAR_PORT"
	pulsarEnvConnectionTimeoutSeconds = "PULSAR_CONNECTION_TIMEOUT_SECONDS"
	pulsarEnvOperationTimeoutSeconds  = "PULSAR_OPERATION_TIMEOUT_SECONDS"
)

var (
	err                            error
	PulsarHost                     string
	PulsarPort                     int
	PulsarConnectionTimeoutSeconds int
	PulsarOperationTimeoutSeconds  int
)

func init() {
	var aux string
	PulsarHost = os.Getenv(pulsarEnvHost)
	if PulsarHost == "" {
		PulsarHost = "localhost"
	}
	aux = os.Getenv(pulsarEnvPort)
	if aux == "" {
		PulsarPort = 6650
	} else {
		PulsarPort, err = strconv.Atoi(aux)
		if err != nil {
			logrus.Info("get env port failed ", err)
			os.Exit(1)
		}
	}
	if aux == "" {
		PulsarConnectionTimeoutSeconds = 15
	} else {
		PulsarConnectionTimeoutSeconds, err = strconv.Atoi(os.Getenv(pulsarEnvConnectionTimeoutSeconds))
		if err != nil {
			logrus.Info("get connection timeout failed ", err)
			os.Exit(1)
		}
	}
	if aux == "" {
		PulsarOperationTimeoutSeconds = 15
	} else {
		PulsarOperationTimeoutSeconds, err = strconv.Atoi(os.Getenv(pulsarEnvOperationTimeoutSeconds))
		if err != nil {
			logrus.Info("get operation timeout failed ", err)
			os.Exit(1)
		}
	}
}
