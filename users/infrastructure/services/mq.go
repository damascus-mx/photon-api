package infrastructure

import (
	"github.com/streadway/amqp"

	core "github.com/damascus-mx/photon-api/users/core/util"
)

// InitMQ Connect to MQ Server and get a channel
func InitMQ(connectionString string) *amqp.Connection {
	conn, err := amqp.Dial(connectionString)
	core.FailOnError("Failed to connect to RabbitMQ", err)
	defer conn.Close()
	return conn
}
