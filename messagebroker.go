package main

// required package imports
import (
  "log"
  "encoding/json"
  "strconv"
  "github.com/streadway/amqp"
)

// function to add messages to Rabbit messaging queues
func callMsgBroker(id int){
  // connect to RabbitMQ server
  conn := RabbitMQConnect()
  defer conn.Close()

  // connect to RabbitMQ connection channel
  amqpChannel, err := conn.Channel()
	HandleError(err)

	defer amqpChannel.Close()

  // declaration of message queue in case it doesn't already exist
	queue, err := amqpChannel.QueueDeclare(configuration.RABBIT_QNAME, true, false, false, false, nil)
	HandleError(err)

  // marshal todoItems id to JSON and add it to message body
  body, err := json.Marshal(strconv.Itoa(id))
	if err != nil {
		HandleError(err)
	}

  // publish message to queue
	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	//log.Printf("AddTask: %d", body)
}
