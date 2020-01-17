package service

import (
	"fmt"
	"log"

	env "github.com/damascus-mx/photon-api/authentication/common/config"
	"github.com/streadway/amqp"
)

// InitMQBroker Start MQ Broker
func InitMQBroker() *amqp.Connection {
	conn, err := amqp.Dial(env.MQ_CONNECTION)
	if err != nil {
		panic(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf(env.FailedService, "MQ Broker"))
	}
	defer ch.Close()

	log.Printf(env.ServiceConnected, "MQ Broker")
	return conn
}
