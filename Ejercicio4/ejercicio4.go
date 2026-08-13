package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Inicializamos la semilla
	rand.Seed(time.Now().UnixNano())

	p := 0.8
	q := 1.0 - p
	totalVariables := 25
	
	secuencia := make([]int, totalVariables)
	numerosGenerados := 0
	
	// Paso 1: Asumir que todos son éxitos inicialmente para optimizar
	for i := 0; i < totalVariables; i++ {
		secuencia[i] = 1
	}

	indiceActual := 0

	// Paso 2: Iterar buscando los índices donde ocurren los fracasos
	for indiceActual < totalVariables {
		u := rand.Float64()
		numerosGenerados++

		// Calcular los ensayos hasta el siguiente fracaso (Distribución Geométrica)
		saltos := int(math.Log(u)/math.Log(q)) + 1
		
		indiceActual += saltos
		
		// Paso 3: Si el fracaso cayó dentro de las 25 variables, registrar el 0
		if indiceActual <= totalVariables {
			// Restamos 1 porque los índices en Go van de 0 a 24
			secuencia[indiceActual-1] = 0 
		}
	}

	fmt.Println("Secuencia de 25 variables de Bernoulli (p=0.8):")
	fmt.Println(secuencia)
	fmt.Printf("Números aleatorios uniformes utilizados: %d\n", numerosGenerados)
}