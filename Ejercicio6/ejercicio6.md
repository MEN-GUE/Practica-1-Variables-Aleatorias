# Ejercicio 6: Estimación de suma de valores únicos

## La idea principal
Cuando una lista es demasiado grande y tiene muchos elementos repetidos, encontrar y sumar cada valor una sola vez puede ser un proceso muy lento. Podemos usar números aleatorios para "adivinar" (estimar) la suma total de forma rápida.

## ¿Cómo funciona el método?
El truco está en tomar en cuenta el valor de un elemento **solo si es la primera vez que aparece** en la lista. 

Los pasos son:
1. Elegimos una posición al azar de la lista completa (de tamaño `n`).
2. Miramos si el valor de esa posición ya apareció en alguna posición anterior. 
3. Si es la primera vez que aparece, anotamos su valor. Si ya estaba repetido antes, anotamos un 0.
4. Repetimos este proceso varias veces. Al final, calculamos el promedio de nuestras anotaciones y lo multiplicamos por el tamaño total de la lista (`n`).

## ¿Por qué funciona?
Porque cada valor distinto tiene exactamente **una sola** "primera aparición" en toda la lista. Al elegir posiciones al azar y descartar las repeticiones, nos aseguramos de no sumar el mismo valor dos veces.

## Ejemplo
```go
package main

import (
	"math/rand"
)

func estimarSumaUnicos(lista []float64, iteraciones int) float64 {
	n := len(lista)
	sumaMuestras := 0.0

	for i := 0; i < iteraciones; i++ {
		// 1. Elegir una posición al azar
		pos := rand.Intn(n)
		valor := lista[pos]

		// 2. Revisar si ese valor ya apareció antes en la lista
		esPrimeraVez := true
		for j := 0; j < pos; j++ {
			if lista[j] == valor {
				esPrimeraVez = false
				break
			}
		}

		// 3. Si es la primera vez que aparece, sumamos su valor
		if esPrimeraVez {
			sumaMuestras += valor
		}
	}

	// 4. Multiplicamos el promedio de las muestras por el tamaño total de la lista
	promedio := sumaMuestras / float64(iteraciones)
	return promedio * float64(n)
}
```
