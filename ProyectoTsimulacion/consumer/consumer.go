package consumer

import (
	"amqp/conexion"
	"log"
)

func Consumer(name_queue string) {
	conn, err := conexion.RabbitMQConn()
	conexion.ErrorHanding(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	conexion.ErrorHanding(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(name_queue, false, false, false, false, nil)
	conexion.ErrorHanding(err, "Failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	conexion.ErrorHanding(err, "Failed to register a consume")

	for msj := range msgs {
		log.Printf("%s", msj.Body)
		if err := msj.Ack(false); err != nil {
			log.Println("no se pudo reconocer el mensaje", err)
		}
	}
}