package main

import (
	"fmt"
)

func main() {
	fmt.Printf("salario empleado 1: %f\n", calcularSalario(60, "A"))
	fmt.Printf("salario empleado 2: %f\n", calcularSalario(60, "B"))
	fmt.Printf("salario empleado 3: %f\n", calcularSalario(60, "C"))
}

func calcularSalario(minutosTrabajados int32, categoria string) (salario float32) {
	var salarioPorHora float32 = 0
	var aumento float32 = 0
	switch categoria {
	case "A":
		salarioPorHora = 1000
	case "B":
		salarioPorHora = 1500
		aumento = 0.20
	case "C":
		salarioPorHora = 3000
		aumento = 0.50
	}

	return (float32(minutosTrabajados) / 60) * salarioPorHora * (1 + aumento)
}
