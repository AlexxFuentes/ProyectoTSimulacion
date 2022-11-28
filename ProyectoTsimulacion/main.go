package main

import (
	"amqp/consumer"
	"amqp/eventos"
	"amqp/producer"
	"fmt"
)

/*
Proyecto: Centro especializado en atención al cliente  RNP HN
*/
func main() {

	//Declaracion de variables
	var (
		estaciones, duracion, recursos = 0, 0, 0
		control_estaciones             = true
	)

	//Entrada de datos
	for control_estaciones {
		fmt.Println("Ingrese el numero de estaciones: ")
		fmt.Scanln(&estaciones)
		fmt.Println("Ingrese el numero de recursos: ")
		fmt.Scanln(&recursos)
		fmt.Println("Ingrese la duracion de la simulacion (en dias): ")
		fmt.Scanln(&duracion)

		if estaciones <= 15 && estaciones <= recursos {
			control_estaciones = false
		} else if estaciones > 15 || estaciones > recursos {
			fmt.Println("El numero de estaciones debe de ser menor o igual a 15 y menor o igual que los recursos.")
		}
	}

	personas_atendias := eventos.Personas_atendidas(estaciones, recursos)
	personas_llegaron := eventos.Personas_llegaron()

	if personas_atendias > personas_llegaron {
		fmt.Println("Todos las personas fueron atendidas!!")
	} else {
		fmt.Println("Pacientes No atendidos: ", personas_llegaron-personas_atendias)
	}

	consumer.Consumer("")
	for i := 0; i < duracion; i++ {
		if personas_atendias > personas_llegaron {
			for i := 0; i < personas_llegaron; i++ {
				producer.Producer("", "Llega")
			}
			consumer.Consumer("")
		} else if personas_atendias < personas_llegaron {
			for i := 0; i < personas_llegaron; i++ {
				if personas_atendias < i {
					producer.Producer("", "Llega")
					consumer.Consumer("")
				} else if i > personas_atendias {
					producer.Producer("", "Sin atender")
				}
			}
		}
	}
}
