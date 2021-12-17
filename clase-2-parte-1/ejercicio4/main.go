package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func main() {

	minFunc, err := operation(minimo)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	promFunc, err := operation(promedio)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	maxFunc, err := operation(maximo)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

	fmt.Printf("valorMinimo: %v\n", valorMinimo)
	fmt.Printf("valorPromedio: %v\n", valorPromedio)
	fmt.Printf("valorMaximo: %v\n", valorMaximo)
}

func operation(tipoOp string) (funcion func(notas ...int32) int32, err error) {
	switch tipoOp {
	case minimo:
		return minFunction, nil
	case maximo:
		return maxFunction, nil
	case promedio:
		return promFunction, nil
	default:
		return nil, errors.New("no se da definido la funcion")
	}
}

func minFunction(notas ...int32) int32 {
	var minNota int32 = math.MaxInt32
	for _, nota := range notas {
		if nota < minNota {
			minNota = nota
		}
	}

	return minNota
}

func maxFunction(notas ...int32) int32 {
	var minNota int32 = 0
	for _, nota := range notas {
		if nota > minNota {
			minNota = nota
		}
	}

	return minNota
}

func promFunction(notas ...int32) int32 {
	var sumaNotas int32 = 0
	for _, nota := range notas {
		sumaNotas += nota
	}
	return sumaNotas / int32(len(notas))
}
