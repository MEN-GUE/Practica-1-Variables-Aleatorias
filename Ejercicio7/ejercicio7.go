package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// GenerarVariableX aplica el algoritmo para simular una variable aleatoria con tasas de riesgo.
// Recibe el límite superior lambdaMax y una función que calcula la tasa para cualquier paso n.
func GenerarVariableX(lambdaMax float64, tasaRiesgo func(int) float64) int {
	// PASO 1: Inicializamos S en 0
	S := 0

	for {
		// PASO 2: Generar U y calcular el salto geométrico Y
		U1 := rand.Float64()
		
		// Utilizamos la fórmula de la transformada inversa para la distribución geométrica
		Y := int(math.Log(U1)/math.Log(1.0-lambdaMax)) + 1

		// PASO 3: Acumulamos el salto para encontrar el siguiente instante candidato
		S = S + Y

		// PASO 4: Generar un nuevo número aleatorio para la prueba de aceptación
		U2 := rand.Float64()

		// PASO 5: Condición de aceptación y rechazo
		lambdaS := tasaRiesgo(S)
		probabilidadAceptacion := lambdaS / lambdaMax

		if U2 <= probabilidadAceptacion {
			// El evento es aceptado. X toma el valor de S y el algoritmo se detiene.
			return S
		}
		// En caso contrario (rechazo), el ciclo se repite regresando al Paso 2 automáticamente.
	}
}

func main() {
	// Refrescamos la semilla de aleatoriedad
	rand.Seed(time.Now().UnixNano())

	// Definimos el límite superior lambda (debe ser mayor a 0 y menor a 1)
	lambdaMax := 0.6

	// Definimos las tasas de riesgo discretas (lambda_n).
	// Para propósitos de esta simulación, creamos un riesgo que aumenta progresivamente 
	// pero que nunca supera a lambdaMax.
	tasaRiesgo := func(n int) float64 {
		tasa := float64(n) * 0.05
		if tasa > lambdaMax {
			return lambdaMax
		}
		return tasa
	}

	fmt.Println("Simulación de variable con tasas de riesgo discretas:")
	fmt.Println("------------------------------------------------------")
	
	// Generamos 10 ejemplos para demostrar el funcionamiento
	for i := 1; i <= 10; i++ {
		resultadoX := GenerarVariableX(lambdaMax, tasaRiesgo)
		fmt.Printf("Prueba %2d: El evento ocurrió en X = %d\n", i, resultadoX)
	}
}