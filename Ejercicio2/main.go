package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generarVariableX contiene la lógica optimizada solicitada en el ejercicio 2
func generarVariableX() int {
	u := rand.Float64()

	// Evaluamos en orden de probabilidad descendente para reducir saltos lógicos
	if u < 0.35 {
		return 3
	} else if u < 0.65 {
		return 1
	} else if u < 0.85 {
		return 2
	}
	
	// Si no cayó en ninguno de los rangos anteriores, por descarte es 4
	return 4
}

func main() {
	// Refrescamos la semilla para que los números cambien en cada ejecución
	rand.Seed(time.Now().UnixNano())

	// Corremos la simulación 1,000,000 de veces para probar que funciona
	iteraciones := 1000000
	conteos := make(map[int]int)

	for i := 0; i < iteraciones; i++ {
		x := generarVariableX()
		conteos[x]++
	}

	fmt.Println("Resultados de la simulación (1,000,000 de iteraciones):")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("%-10s | %-20s | %-20s\n", "Valor X", "Prob. Simulada", "Prob. Teórica")
	fmt.Println("---------------------------------------------------------")

	// Definimos el orden para imprimir la tabla limpiamente (1 a 4)
	ordenResultados := []int{1, 2, 3, 4}
	probTeoricas := map[int]float64{
		1: 0.30,
		2: 0.20,
		3: 0.35,
		4: 0.15,
	}

	for _, valor := range ordenResultados {
		probSimulada := float64(conteos[valor]) / float64(iteraciones)
		fmt.Printf("X = %-6d | %-20.4f | %-20.4f\n", valor, probSimulada, probTeoricas[valor])
	}
	fmt.Println("---------------------------------------------------------")
}