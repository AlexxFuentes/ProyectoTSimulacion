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
		estaciones, duracion, recursos, control_estaciones = 0, 0, 0, true
	)
	//Entrada de datos
	for control_estaciones {
		fmt.Println("Ingrese el numero de estaciones: ")
		fmt.Scanln(&estaciones)
		fmt.Println("Ingrese el numero de recursos: ")
		fmt.Scanln(&recursos)
		fmt.Println("Ingrese la duracion de la simulacion (en dias): ")
		fmt.Scanln(&duracion)

		if estaciones <= 15 && recursos >= estaciones && recursos <= 2*estaciones{
			control_estaciones = false
		} else if estaciones > 15 || recursos < estaciones || recursos > 2*estaciones {
			fmt.Println("El numero de estaciones debe de ser menor o igual a 15.")
			fmt.Println("Los recursos deben de ser mayores que las estaciones y menores que el doble de las estaciones.")
		}
	}
	personas_llegaron_dias := eventos.Personas_llegaron(duracion)
	eventos.Personas_atendidas(estaciones, recursos, duracion, personas_llegaron_dias) //personas_atendias_dias
}