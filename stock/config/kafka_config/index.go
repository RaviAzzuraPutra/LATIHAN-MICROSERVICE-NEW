package kafka_config

import "os"

var Kafka_Broker string
var Kafka_Topic string
var GroupID string

func KAFKA_CONFIG() {
	Kafka_Broker = os.Getenv("KAFKA_BROKER")
	Kafka_Topic = os.Getenv("KAFKA_TOPIC")
	GroupID = os.Getenv("GROUP_ID")
}
