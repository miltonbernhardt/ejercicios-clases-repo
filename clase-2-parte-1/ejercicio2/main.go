package main

import (
	"errors"
	"fmt"
)

func main() {

	promedio, errorPromedio := calcularPromedio(10, 5, 7, 8)

	if errorPromedio != nil {
		fmt.Printf("%v", errorPromedio)
	} else {

		fmt.Printf("promedio: %f", promedio)
	}
}

func calcularPromedio(calificaciones ...int32) (impuesto float32, error error) {
	if len(calificaciones) == 0 {
		return 0, errors.New("no se pasaron notas para calcular el promedio")
	}

	var sumaNotas int32 = 0
	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("la calcificación no puede ser negativa")
		} else {
			sumaNotas += calificacion
		}
	}


	return float32(sumaNotas) / float32(len(calificaciones)), nil
}
