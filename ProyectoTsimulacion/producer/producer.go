package producer

import (
	"amqp/conexion"
	"github.com/streadway/amqp"
)

func Producer(name_queue string, msg string) {
	conn, err := conexion.RabbitMQConn()
	conexion.ErrorHanding(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	conexion.ErrorHanding(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(name_queue, false, false, false, false, nil)
	conexion.ErrorHanding(err, "Failed to declare a queue")

	err = ch.Publish("", queue.Name, false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	//log.Println(" [x] Enviado")
	conexion.ErrorHanding(err, "Error al publicar mensaje")
}