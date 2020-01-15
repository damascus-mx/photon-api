package infrastructure

import (
	"database/sql"
	"github.com/streadway/amqp"
)

// Broker AMQP Broker
type Broker struct {
	MQ *amqp.Connection
	DB *sql.Conn
}

// NewBroker Get a Broker instance
func NewBroker(mq *amqp.Connection, db *sql.Conn) *Broker {
	return &Broker{MQ: mq, DB: db}
}

// getChan Create and get a new AMQP Channel
func (b *Broker) getChan() *amqp.Channel {
	// ch, err := b.MQ.Channel()
	return nil
}
