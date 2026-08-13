package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simularLanzamientos ejecuta un experimento hasta obtener todas las sumas posibles
func simularLanzamientos() int {
	// Arreglo para marcar las sumas vistas.
	// Los índices van del 0 al 12, aprovecharemos las posiciones del 2 al 12.
	vistos := make([]bool, 13)
	sumasUnicas := 0
	lanzamientos := 0

	// Continuar hasta haber visto las 11 sumas posibles (del 2 al 12)
	for sumasUnicas < 11 {
		// Simular dos dados (valores del 1 al 6)
		dado1 := rand.Intn(6) + 1
		dado2 := rand.Intn(6) + 1
		suma := dado1 + dado2

		lanzamientos++

		// Si es la primera vez que vemos esta suma, actualizar registros
		if !vistos[suma] {
			vistos[suma] = true
			sumasUnicas++
		}
	}

	return lanzamientos
}

func main() {
	rand.Seed(time.Now().UnixNano())

	iteraciones := 100000
	sumaTotalLanzamientos := 0

	// Ejecutar la simulación múltiples veces para encontrar el valor esperado
	for i := 0; i < iteraciones; i++ {
		sumaTotalLanzamientos += simularLanzamientos()
	}

	// Calcular el promedio esperado
	esperanza := float64(sumaTotalLanzamientos) / float64(iteraciones)

	fmt.Println("Simulación de Lanzamiento de Dados")
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Iteraciones ejecutadas          : %d\n", iteraciones)
	fmt.Printf("Esperanza (Promedio de intentos): %.2f lanzamientos\n", esperanza)
}