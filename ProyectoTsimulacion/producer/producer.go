package producer

import (
	"amqp/conexion"
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func Producer(name_queue string, msg string) {
	conn, err := conexion.RabbitMQConn()
	conexion.ErrorHanding(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	conexion.ErrorHanding(err, "Failed to open a channel")
	defer ch.Close()

	//crear una cola para guardar mensajes
	queue, err := ch.QueueDeclare(name_queue, false, false, false, false, nil)
	conexion.ErrorHanding(err, "Failed to declare a queue")

	err = ch.Publish("", queue.Name, false, false,
		amqp.Publishing{
			Headers:     nil,
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	log.Println(" [x] Enviado")
	conexion.ErrorHanding(err, "Error al publicar mensaje")
}

func Enviar_mensajes(name_queue string, user string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(user, "Escriba mensaje: \n--EXIT")
	for true {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		Producer(name_queue, user+": "+msg)
		if strings.Contains(msg, "--EXIT") {
			break
		}
	}
}
