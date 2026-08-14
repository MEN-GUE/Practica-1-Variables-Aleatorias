package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Inicializamos una fuente aleatoria local (mejor práctica en versiones recientes de Go)
	fuente := rand.NewSource(time.Now().UnixNano())
	r := rand.New(fuente)
	
	iteraciones := 1000000

	fmt.Println("Resultados de Estimación Monte Carlo (1,000,000 iteraciones)")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-10s | %-15s | %-15s | %-25s\n", "Ejercicio", "Estimación", "Valor Exacto", "Notas")
	fmt.Println("--------------------------------------------------------------------------------")

	// Ejercicio 3: Integral de 0 a 1 de exp(e^x)
	// Solución exacta aproximada: 6.31506 (Ei(e) - Ei(1))
	var suma3 float64
	for i := 0; i < iteraciones; i++ {
		u := r.Float64()
		suma3 += math.Exp(math.Exp(u))
	}
	est3 := suma3 / float64(iteraciones)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 3", est3, 6.31506, "Sin analítica simple")

	// Ejercicio 4: Integral de 0 a 1 de (1 - x^2)^(3/2)
	// Solución exacta: 3*PI/16 (aprox 0.58905)
	var suma4 float64
	for i := 0; i < iteraciones; i++ {
		u := r.Float64()
		suma4 += math.Pow(1.0-u*u, 1.5)
	}
	est4 := suma4 / float64(iteraciones)
	exacto4 := (3.0 * math.Pi) / 16.0
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 4", est4, exacto4, "Exacta: 3pi/16")

	// Ejercicio 5: Integral de -2 a 2 de exp(x + x^2)
	// Limites a y b. Intervalo = 4. Transformación: x = -2 + 4u
	// Valor numérico exacto aprox 93.16275
	var suma5 float64
	for i := 0; i < iteraciones; i++ {
		u := r.Float64()
		x := -2.0 + 4.0*u
		// Se multiplica por la longitud del intervalo (4)
		suma5 += 4.0 * math.Exp(x+x*x)
	}
	est5 := suma5 / float64(iteraciones)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 5", est5, 93.16275, "Sin analítica simple")

	// Ejercicio 6: Integral de 0 a infinito de x * (1 + x^2)^-2
	// Sustitución: x = 1/u - 1. Multiplicador diferencial: 1/u^2
	// Solución exacta: 1/2
	var suma6 float64
	for i := 0; i < iteraciones; i++ {
		u := r.Float64()
		x := (1.0 / u) - 1.0
		val := x / math.Pow(1.0+x*x, 2)
		suma6 += val / (u * u)
	}
	est6 := suma6 / float64(iteraciones)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 6", est6, 0.50000, "Exacta: 1/2")

	// Ejercicio 7: Integral de -inf a inf de exp(-x^2)
	// Por simetría, es 2 * integral de 0 a infinito. Sustitución x = 1/u - 1
	// Solución exacta: Raíz de PI (aprox 1.77245)
	var suma7 float64
	for i := 0; i < iteraciones; i++ {
		u := r.Float64()
		x := (1.0 / u) - 1.0
		val := 2.0 * math.Exp(-x*x) // Multiplicado por 2 por la simetría
		suma7 += val / (u * u)
	}
	est7 := suma7 / float64(iteraciones)
	exacto7 := math.Sqrt(math.Pi)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 7", est7, exacto7, "Exacta: Raíz de pi")

	// Ejercicio 8: Integral doble de 0 a 1 de exp((x+y)^2)
	// Requiere dos variables uniformes U1 y U2
	// Valor numérico exacto aprox 4.89916
	var suma8 float64
	for i := 0; i < iteraciones; i++ {
		u1 := r.Float64()
		u2 := r.Float64()
		suma8 += math.Exp(math.Pow(u1+u2, 2))
	}
	est8 := suma8 / float64(iteraciones)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 8", est8, 4.89916, "Doble. Sin analítica")

	// Ejercicio 9: Integral doble de 0 a inf, y de 0 a x de exp(-(x+y))
	// Integración con indicadora. Sustitución en ambas: x = 1/u1 - 1, y = 1/u2 - 1
	// Solución exacta: 1/2
	var suma9 float64
	for i := 0; i < iteraciones; i++ {
		u1 := r.Float64()
		u2 := r.Float64()

		x := (1.0 / u1) - 1.0
		y := (1.0 / u2) - 1.0

		// Función indicadora: solo sumamos si y < x
		if y < x {
			val := math.Exp(-(x + y))
			suma9 += val / (u1 * u1 * u2 * u2) // Multiplicadores de los diferenciales
		}
	}
	est9 := suma9 / float64(iteraciones)
	fmt.Printf("%-10s | %-15.5f | %-15.5f | %-25s\n", "Ej. 9", est9, 0.50000, "Exacta: 1/2")
	fmt.Println("--------------------------------------------------------------------------------")
}