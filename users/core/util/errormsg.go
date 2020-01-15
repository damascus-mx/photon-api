package core

import "log"

// FailOnError Publish a error message
func FailOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("\n%s: %s\n", msg, err)
	}
}
