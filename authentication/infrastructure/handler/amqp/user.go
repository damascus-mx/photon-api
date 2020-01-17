package handler

import (
	env "github.com/damascus-mx/photon-api/authentication/common/config"
	"github.com/streadway/amqp"
	"log"
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
	m.onUserChange()
}

// onUserCreate Listen to user creation
func (m *MQUser) onUserCreate() {
	ch, err := m.MQ.Channel()
	if err != nil {
		log.Fatalf(env.FailedMQError, "onUserCreate", err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"user.create", // name
		true,          // durable
		true,          // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal(env.FailedMQError, "onUserChanged", err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(env.FailedMQError, "onUserChanged", err.Error())
	}

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
	if err != nil {
		log.Fatal(env.FailedMQError, "onUserChanged", err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"user.update", // name
		true,          // durable
		true,          // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal(env.FailedMQError, "onUserChanged", err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(env.FailedMQError, "onUserChanged", err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages.")
	<-forever
}
