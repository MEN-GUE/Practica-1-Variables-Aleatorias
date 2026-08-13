package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GeneradorX representa la función que ejecuta el paso (a)
type GeneradorX func() int

// ProbabilidadCondicional representa la evaluación de P(Y=j | X=i) del paso (d)
type ProbabilidadCondicional func(i int) float64

// GenerarVariableW implementa el algoritmo exacto descrito en el problema
func GenerarVariableW(generarX GeneradorX, probYDadoX ProbabilidadCondicional) int {
	for {
		// Pasos (a) y (b): Generar el valor de una variable aleatoria con la distribución de X
		i := generarX()

		// Paso (c): Generar un número aleatorio U
		u := rand.Float64()

		// Paso (d): Evaluar la condición de aceptación
		// Si U < P(Y=j | X=i), haga W = i y deténgase
		if u < probYDadoX(i) {
			return i
		}

		// Paso (e): Si la condición no se cumple, el ciclo for hace que regrese a (a)
	}
}

func main() {
	// Refrescar la semilla aleatoria
	rand.Seed(time.Now().UnixNano())

	// --- EJEMPLO DE USO ---
	// Para probar que el algoritmo funciona, simularemos un escenario abstracto:
	
	// Simulamos que la distribución de X corresponde a lanzar un dado de 6 caras.
	simularX := func() int {
		return rand.Intn(6) + 1
	}

	// Simulamos una probabilidad condicional ficticia.
	// Haremos que la probabilidad de aceptar el número dependa de su propio valor.
	// Ej: si i=6, la probabilidad de aceptación es 6/6 (100%). Si i=1, es 1/6.
	probabilidadY := func(i int) float64 {
		return float64(i) / 6.0
	}

	fmt.Println("Generando valores para la variable W utilizando el método de rechazo:")
	fmt.Println("-------------------------------------------------------------------")
	
	for k := 1; k <= 10; k++ {
		w := GenerarVariableW(simularX, probabilidadY)
		fmt.Printf("Ejecución %2d: W = %d\n", k, w)
	}
}