package conexion

import (
	"github.com/streadway/amqp"
	"log"
)

// Función de conexión RabbitMQ
func RabbitMQConn() (conn *amqp.Connection, err error) {
	// Nombre de usuario
	var user string = "guest"
	// Contraseña
	var pwd string = "guest"
	// host
	var host string = "localhost"
	// Puerto
	var port string = "5672"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"
	// Crea una nueva conexión
	conn, err = amqp.Dial(url)
	// devuelve conexión y error
	return
}

// Función de manejo de errores
func ErrorHanding(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
