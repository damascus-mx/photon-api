package handler

import (
	"log"

	"github.com/damascus-mx/photon-api/authentication/common/util"
	"github.com/streadway/amqp"
)

// MQUser MQ Broker for user
type MQUser struct {
	MQ *amqp.Connection
}

// NewMQUser Get MQ Broker for user
func NewMQUser(mqConn *amqp.Connection) *MQUser {
	return &MQUser{MQ: mqConn}
}

func (m *MQUser) init() {
	// Start event-listeners
	ch, err := m.MQ.Channel()
	util.FailOnErrorMQ(err, "onUserExchange")
	defer ch.Close()

	// Create an exchange to bind queues
	err = ch.ExchangeDeclare(
		"user",  // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	util.FailOnErrorMQ(err, "onUserExchange")

	// Start and bind queues
	m.onUserCreate()
	m.onUserChange()
}

// onUserCreate Listen to user creation
func (m *MQUser) onUserCreate() {
	ch, err := m.MQ.Channel()
	util.FailOnErrorMQ(err, "onUserCreate")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"create_user_queue", // name
		true,                // durable
		true,                // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	util.FailOnErrorMQ(err, "onUserCreate")

	// Bind queue to exchange
	err = ch.QueueBind(
		q.Name,   // Queue name
		"create", // Binding key
		"user",   // Exchange name
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnErrorMQ(err, "onUserCreate")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}

// onUserChange Listen to user changes
func (m *MQUser) onUserChange() {
	ch, err := m.MQ.Channel()
	util.FailOnErrorMQ(err, "onUserChanged")
	defer ch.Close()

	// Create queue to consume incoming messages
	q, err := ch.QueueDeclare(
		"update_user_queue", // name
		false,               // durable
		true,                // delete when unused
		true,                // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	util.FailOnErrorMQ(err, "onUserChanged")

	// Bind to topic user.update route pattern
	ch.QueueBind(
		q.Name,   // Queue name
		"update", // Binding Key
		"user",   // exchange
		false,
		nil,
	)

	// Consume messages from queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnErrorMQ(err, "onUserChanged")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}
