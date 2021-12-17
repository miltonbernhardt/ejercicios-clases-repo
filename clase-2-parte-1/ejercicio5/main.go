package main

import (
	"errors"
	"fmt"
)

const (
	perro = "perro"
	gato  = "gato"
)

func main() {

	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var cantidad float64
	cantidad += animalPerro(5)
	cantidad += animalGato(8)

	fmt.Printf("total alimento a comprar: %v\n", cantidad)
}

func Animal(tipoAnimal string) (funcion func(cantidadAnimales int32) float64, err error) {
	switch tipoAnimal {
	case perro:
		return Perro, nil
	case gato:
		return Gato, nil
	default:
		return nil, errors.New("no se da el tipo de animal")
	}
}

func Gato(cantidadGatos int32) float64 {
	return float64(cantidadGatos * 5)

}

func Perro(cantidadPerros int32) float64 {
	return float64(cantidadPerros * 10)
}
