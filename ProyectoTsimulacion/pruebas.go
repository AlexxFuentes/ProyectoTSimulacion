package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(num_ramdon(5, 11))

}

func num_ramdon(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func crear_arr(estaciones int) map[interface{}]interface{} {
	est := make(map[interface{}]interface{})
	for i := 1; i <= estaciones; i++ {
		est[i] = i
	}
	return est
}
