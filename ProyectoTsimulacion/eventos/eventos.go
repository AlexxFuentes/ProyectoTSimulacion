package eventos

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Funcion que retorna la cantidad de personas que fueron atendidas
durante un intervalos de tiempo a una frecuencia especifica

parametros: intervalo de tiempo en min, frecuencia por min
*/
func Frecuencia_personas(intervalo_tiempo int, frecuencia_min float64) int {
	cont_personas := 0
	i := 0
	rand.Seed(time.Now().Unix())

	for i < intervalo_tiempo {
		j := rand.Float64()
		if j > (1 - frecuencia_min) {
			cont_personas++
		}
		i++
	}
	return cont_personas
}

/*
Funcion que retorna el total de personas atendidas en todo un dia
*/
func Personas_llegaron() int {

	intervalo_tiempo := []int{180, 179, 89, 90, 239}          //Intervalo de tiempo en minutos
	frecuencia_min := []float64{0.31, 0.46, 0.55, 0.23, 0.73} // frecuencia por minuto
	frecuencia := 0
	i := 0

	for i < 5 {
		frecuencia += Frecuencia_personas(intervalo_tiempo[i], frecuencia_min[i])
		fmt.Println("\n", frecuencia)
		i++
	}
	fmt.Printf("Cantidad total de personas que llegaron: %d", frecuencia)
	return frecuencia
}

/*
Funcion que devuelve un numero aleatorio entre [5, 10] minutos
simulando el tiempo de atencion de una persona.
*/
func tiempo_atencion(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/*
Fucion que simula la cantidad de personas atendidas segun
la cantidad de estaciones y numero de recursos asignados

retorna cantidad de personas atendidas
*/
func Personas_atendidas(estaciones, recursos int) int {
	var (
		min_dia           = 777 //Minutos de atencion en un dia
		min_transcurridos = 0   //Contador de minutos transcurridos
		cant_recursos     = 0   //asigna la cantidad recursos
		num_atendidos     = 0   //numero de personas atendidas
	)

	cant_recursos = estaciones
	for i := 0; i <= min_dia; i++ {
		if min_transcurridos <= 360 { //Primeras 6 horas = 360 min
			i += tiempo_atencion(5, 11)
			min_transcurridos = i
			num_atendidos += cant_recursos
		} else if min_transcurridos <= 720 { // cambio de turno desde 360 min hasta 720 min
			i += tiempo_atencion(5, 11)
			min_transcurridos = i
			num_atendidos += recursos - cant_recursos
		} else if min_transcurridos <= min_dia { // ultimas horas desde 720 min hasta 777 min
			i += tiempo_atencion(5, 11)
			min_transcurridos = i
			num_atendidos += cant_recursos
		}
	}

	fmt.Printf("Posibles personas atendidas %d", num_atendidos)
	return num_atendidos
}
