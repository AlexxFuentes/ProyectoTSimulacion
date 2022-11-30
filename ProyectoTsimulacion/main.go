package main

import (
	"amqp/eventos"
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
		fmt.Println("Ingrese la duracion de la simulacion (en dias): ")
		fmt.Scanln(&duracion)
		fmt.Println("Ingrese el numero de recursos: ")
		fmt.Scanln(&recursos)

		if estaciones <= 15 && recursos >= estaciones {
			control_estaciones = false
		} else if estaciones > 15 || recursos < estaciones {
			fmt.Println("El numero de estaciones debe de ser menor o igual a 15 y menor o igual que los recursos.")
		}
	}

	personas_llegaron_dias := eventos.Personas_llegaron(duracion)
	eventos.Personas_atendidas(estaciones, recursos, duracion, personas_llegaron_dias) //personas_atendias_dias

}
