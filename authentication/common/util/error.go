package util

import (
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
)

// FailOnErrorMQ Write error message on log for MQ Broker
func FailOnErrorMQ(err error, msg string) {
	if err != nil {
		log.Fatalf(env.FailedMQError, msg, err.Error())
	}
}
