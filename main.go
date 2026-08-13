package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inicializar la semilla para generar números aleatorios distintos en cada corrida
	rand.Seed(time.Now().UnixNano())

	diasSimulacion := 15
	inventarioDiario := 3

	// Variables para acumular las métricas
	totalVentas := 0
	totalSobrantes := 0
	totalVentasPerdidas := 0

	fmt.Println("Simulación de Inventario - Tienda de Conveniencia")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("%-5s | %-12s | %-8s | %-8s | %-10s | %-15s\n", "Día", "Num Aleat", "Demanda", "Ventas", "Sobrantes", "Ventas Perdidas")
	fmt.Println("--------------------------------------------------------------------------------")

	for dia := 1; dia <= diasSimulacion; dia++ {
		// Generar número aleatorio entre 0.0 y 1.0
		numAleatorio := rand.Float64()
		demanda := 0

		// Mapear el número aleatorio a la demanda según la distribución acumulada
		if numAleatorio < 0.05 {
			demanda = 0
		} else if numAleatorio < 0.20 {
			demanda = 1
		} else if numAleatorio < 0.50 {
			demanda = 2
		} else if numAleatorio < 0.80 {
			demanda = 3
		} else if numAleatorio < 0.95 {
			demanda = 4
		} else {
			demanda = 5
		}

		// Calcular resultados del día
		ventas := 0
		sobrantes := 0
		ventasPerdidas := 0

		if demanda <= inventarioDiario {
			ventas = demanda
			sobrantes = inventarioDiario - demanda
		} else {
			ventas = inventarioDiario
			ventasPerdidas = demanda - inventarioDiario
		}

		// Acumular métricas
		totalVentas += ventas
		totalSobrantes += sobrantes
		totalVentasPerdidas += ventasPerdidas

		// Imprimir la fila del día
		fmt.Printf("%-5d | %-12.4f | %-8d | %-8d | %-10d | %-15d\n", dia, numAleatorio, demanda, ventas, sobrantes, ventasPerdidas)
	}

	fmt.Println("--------------------------------------------------------------------------------")
	
	// Resumen de métricas
	fmt.Println("\nResumen de la Simulación (15 días):")
	fmt.Printf("Total de botellas vendidas: %d\n", totalVentas)
	fmt.Printf("Total de botellas sobrantes (exceso): %d\n", totalSobrantes)
	fmt.Printf("Total de ventas perdidas (escasez): %d\n", totalVentasPerdidas)

	// Imprimir la conclusión directamente desde el programa
	fmt.Println("\nConclusión del Análisis:")
	fmt.Println("El inventario fijo de 3 botellas cubre la demanda en el 80% de los escenarios probables.")
	fmt.Println("Sin embargo, el valor esperado de la demanda es de 2.5 botellas diarias.")
	fmt.Println("Mantener 3 botellas asegura pocas ventas perdidas, pero genera exceso de inventario")
	fmt.Println("cerca de la mitad de los días (cuando la demanda es 0, 1 o 2). Si el costo de")
	fmt.Println("almacenamiento es despreciable, la política de 3 botellas es muy sólida.")
}