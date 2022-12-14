package eventos

import (
	//"amqp/consumer"
	//"amqp/producer"
	"fmt"
	"math"
	"math/rand"
	"time"
	//"strconv"
)

/*
Funcion que retorna la cantidad de personas que fueron atendidas
durante un intervalos de tiempo a una frecuencia especifica

parametros: 
	intervalo de tiempo en min	: int
	frecuencia por min			: float64
*/
func Frecuencia_personas(intervalo_tiempo int, frecuencia_min float64) int {
	var (
		cont_personas, i = 0, 0
	)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < intervalo_tiempo {
		j := r.Float64()
		if j > (1 - frecuencia_min) {
			cont_personas++
		}
		i++
	}
	return cont_personas
}

/*
Funcion que retorna el total de personas atendidas en todo un dia

Parametros: 
	numero de dias : int
*/
func Personas_llegaron(dias int) map[int]int {
	var (
		intervalo_tiempo             = []int{180, 180, 90, 90, 240}            //Intervalo de tiempo en minutos
		frecuencia_min               = []float64{0.31, 0.46, 0.55, 0.23, 0.73} // frecuencia por minuto
		intervalo_horas              = []string{"04:30AM a 07:30AM", "07:31AM a 10:30AM", "10:31AM a 12:00M", "12:00AM a 01:30PM", "01:31PM a 05:30PM"}
		frecuencia, frecuencia_hora  = 0, 0
		frecuencia_personas_llegaron = make(map[int]int)
	)

	for dia := 1; dia <= dias; dia++ {
		frecuencia = 0
		for i := 0; i < 5; i++ {
			frecuencia_hora = Frecuencia_personas(intervalo_tiempo[i], frecuencia_min[i])
			fmt.Printf("Dia: %d, Intervalo de horas: %s frecuencia de %d personas. \n", dia, intervalo_horas[i], frecuencia_hora)
			frecuencia += frecuencia_hora
		}
		frecuencia_personas_llegaron[dia] = frecuencia
		fmt.Printf("Total de personas que llegaron el dia %d: %d \n\n", dia, frecuencia)
	}

	fmt.Printf("Cantidad total de personas que llegaron por dia: %d \n\n\n", frecuencia_personas_llegaron)
	return frecuencia_personas_llegaron
}

/*
Funcion que devuelve un numero aleatorio entre [5, 10] minutos
simulando el tiempo de atencion de una persona.

Parametros: 
	min : int
	max : int
*/
func Tiempo_atencion(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

/*
Fucion que simula la cantidad de personas atendidas segun las que hayan llegado
en un un dia, de acuerdo a la cantidad de estaciones y numero de recursos asignados

Parametros: 
	estaciones					  : int
	recursos					  : int
	dias 						  : int
	personas que llegaron por dia : map
*/
func Personas_atendidas(estaciones, recursos, dias int, personas_llegaron_dias map[int]int) {
	var (
		min_dia             = 780 //minutos de atencion en un dia
		min_transcurridos   = 0   //contador de minutos transcurridos
		cant_estaciones     = 0   //asigna la cantidad de estaciones
		cant_estaciones_esp = 0   //cantidad de recursos que quedaran en espera
	)

	if estaciones == recursos {
		cant_estaciones = estaciones
		cant_estaciones_esp = recursos
	} else if estaciones < recursos {
		cant_estaciones = estaciones
		cant_estaciones_esp = recursos - estaciones
	}

	for dia := 1; dia <= dias; dia++ { // Dias
		min_transcurridos = 0
		asignar_estacion_personas_dia := Asignar_personas_estaciones(estaciones, personas_llegaron_dias[dia])

		fmt.Printf("*** Dia: %d *** \n", dia)
		fmt.Printf("Numero de personas que llegaron el dia %d: %d \n", dia, personas_llegaron_dias[dia])
		fmt.Printf("Asignacion de estaciones a las personas antes de ser atendidas: \n")
		imprimir(asignar_estacion_personas_dia, "asignadas")

		for min := 0; min <= min_dia; min++ { // Minutos de un dia de trabajo => 13 horas = 780 minutos

			if min_transcurridos <= 360 { //Primeras 6 horas = 360 min
				min += Tiempo_atencion(5, 11)
				min_transcurridos = min
				
				for estacion := 1; estacion <= cant_estaciones; estacion++ {
					//producer.Producer("Personas_atendidas_dia_"+strconv.Itoa(dia), "Persona atendida en la estacion "+strconv.Itoa(estacion))
					if asignar_estacion_personas_dia[estacion] == 0 {
						asignar_estacion_personas_dia[estacion] = 0
					} else {
						asignar_estacion_personas_dia[estacion] -= 1
					}
				}
				//consumer.Consumer("Personas_atendidas_dia_" + strconv.Itoa(dia))

			} else if min_transcurridos <= 720 { // cambio de turno desde 360 min hasta 720 min
				min += Tiempo_atencion(5, 11)
				min_transcurridos = min
				
				for estacion := 1; estacion <= cant_estaciones_esp; estacion++ {
					//producer.Producer("Personas_atendidas_dia_"+strconv.Itoa(dia), "Persona atendida en la estacion "+strconv.Itoa(estacion))
					if asignar_estacion_personas_dia[estacion] == 0 {
						asignar_estacion_personas_dia[estacion] = 0
					} else {
						asignar_estacion_personas_dia[estacion] -= 1
					}
				}
				//consumer.Consumer("Personas_atendidas_dia_" + strconv.Itoa(dia))
			} else if min_transcurridos <= min_dia { // ultimas hora desde 720 min hasta 780 min
				min += Tiempo_atencion(5, 11)
				min_transcurridos = min
				
				for estacion := 1; estacion <= cant_estaciones; estacion++ {
					//producer.Producer("Personas_atendidas_dia_"+strconv.Itoa(dia), "Persona atendida en la estacion "+strconv.Itoa(estacion))
					if asignar_estacion_personas_dia[estacion] == 0 {
						asignar_estacion_personas_dia[estacion] = 0
					} else {
						asignar_estacion_personas_dia[estacion] -= 1
					}
				}
				//consumer.Consumer("Personas_atendidas_dia_" + strconv.Itoa(dia))
			}
		}
		fmt.Printf("Estaciones despues de atender a las personas durante el dia %d \n", dia)
		imprimir(asignar_estacion_personas_dia, "no atendidas")
		fmt.Printf("Numero de personas que llegaron el dia %d: %d \n", dia, personas_llegaron_dias[dia])
		fmt.Printf("Numero de personas que no se atendieron el dia %d: %d \n\n\n", dia, verificar(asignar_estacion_personas_dia))
	}
}

/*
Funcion que asigna de forma aleatoria las personas que llegaron
a las estaciones disponibles

Parametros: 
	numero de estaciones																	: int
	Objetivo: numero de personas que vamos a dividir de forma aleatoria en las estaciones	: int
*/
func Asignar_personas_estaciones(num_estaciones , objetivo int) map[int]int {
	var (
		est            = make(map[int]int)
		sum_est        int
		sum            float64
		sum_aleatorios = make(map[int]float64)
	)
	for i := 1; i <= num_estaciones; i++ {
		sum_aleatorios[i] = rand.Float64()
		sum += sum_aleatorios[i]
	}
	for i := 1; i <= num_estaciones; i++ {
		est[i] = int(math.Floor((sum_aleatorios[i] / sum) * float64(objetivo)))
		sum_est += est[i]
	}
	if sum_est < objetivo {
		est[1] = est[1] + objetivo - sum_est
		sum_est = 0
		for i := 1; i <= num_estaciones; i++ {
			sum_est += est[i]
		}
	}
	return est
}

/*
Funcion que imprime de forma ordenada como
fueron asignada las personas en las estaciones

Parametros: 
	m: map donde se guarda el numero de estacion y el numero de personas por estacion 
	s: cadena de texto con algun mensaje para personalizar la salida
*/
func imprimir(m map[int]int, s string) {
	var s1 = "Todos fueron atendidos en esta estacion."
	for index, value := range m {
		if value == 0 {
			fmt.Printf("Estacion %d: %s \n", index, s1)
		} else {
			fmt.Printf("Estacion %d: con %d personas %s. \n", index, value, s)
		}
	}
}

/*
Funcion que verificara la cantidad de personas que fueron atendidas

Parametros:
	m: map donde se guarda el numero de personas por estacion
*/
func verificar(m map[int]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}