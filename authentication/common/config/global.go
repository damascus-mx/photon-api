package env

const (
	// ServiceConnected Service started
	ServiceConnected string = "\nSERVICE: %s service started ...\n"
	// FailedService Service failed
	FailedService string = "\nSERVICE: Failed to start %s service\n"
	// FailedMQError Fail to do MQ action
	FailedMQError string = "\nMQ BROKER: Failed to do operation on %s\n-		Error: %s\n"
	// MQExchange MQ Broker Exchange
	MQExchange string = "auth"
)
