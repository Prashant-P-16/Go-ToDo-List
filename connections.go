package main

// required package imports
import (
  "github.com/garyburd/redigo/redis"
  "github.com/streadway/amqp"
)

// function to manage connections to Redis & return Redis connection object
func RedisConnect() redis.Conn {
        conn, err := redis.Dial(configuration.REDIS_NETWORK, configuration.REDIS_ADDR)
        HandleError(err)
        return conn
}

// function to manage connections to RabbitMQ & return RabbitMQ connection object
func RabbitMQConnect() *amqp.Connection {
        conn, err := amqp.Dial(configuration.RABBIT_CONN_URL)
        HandleError(err)
        return conn
}
