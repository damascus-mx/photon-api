package env

const (
	// ServiceConnected Service started
	ServiceConnected string = "SERVICE: %s service started ..."
	// FailedService Service failed
	FailedService string = "SERVICE: Failed to start %s service"
	// FailedMQError Fail to do MQ action
	FailedMQError string = "MQ BROKER: Failed to do operation on %s-		Error: %s"
	// MQExchange MQ Broker Exchange
	MQExchange string = "user_events"
)
