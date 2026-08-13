package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Se crea una fuente local de aleatoriedad para evitar dependencias globales y asegurar una buena semilla
	fuente := rand.NewSource(time.Now().UnixNano())
	r := rand.New(fuente)

	numCartas := 100
	simulaciones := 1000000

	cartas := make([]int, numCartas)

	var sumaAciertos float64
	var sumaAciertosCuadrado float64

	fmt.Println("Iniciando simulación...")

	for i := 0; i < simulaciones; i++ {
		// Llenar la baraja en orden inicial
		for j := 0; j < numCartas; j++ {
			cartas[j] = j + 1
		}

		// Mezclar aleatoriamente el arreglo
		r.Shuffle(numCartas, func(i, j int) {
			cartas[i], cartas[j] = cartas[j], cartas[i]
		})

		// Contar los aciertos de la baraja mezclada
		aciertos := 0
		for j := 0; j < numCartas; j++ {
			// El índice es j, pero visualmente las cartas son del 1 al 100 (j+1)
			if cartas[j] == j+1 {
				aciertos++
			}
		}

		// Almacenar los datos brutos de la iteración
		aciertosFloat := float64(aciertos)
		sumaAciertos += aciertosFloat
		sumaAciertosCuadrado += aciertosFloat * aciertosFloat
	}

	// Esperanza estimada: Promedio total de aciertos
	esperanzaEstimada := sumaAciertos / float64(simulaciones)

	// Varianza estimada: E[X^2] - (E[X])^2
	esperanzaCuadrados := sumaAciertosCuadrado / float64(simulaciones)
	varianzaEstimada := esperanzaCuadrados - (esperanzaEstimada * esperanzaEstimada)

	// Imprimir reporte de resultados
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Iteraciones ejecutadas : %d\n", simulaciones)
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-25s | %-10s | %-10s\n", "Métrica", "Estimado", "Exacto")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-25s | %-10.4f | %-10.4f\n", "Esperanza (Media)", esperanzaEstimada, 1.0000)
	fmt.Printf("%-25s | %-10.4f | %-10.4f\n", "Varianza", varianzaEstimada, 1.0000)
	fmt.Println("--------------------------------------------------")
}