package main

import (
	"fmt"
	"os"
	"strings"
)

type Producto struct {
	id       int
	precio   float32
	cantidad int
}

func main() {
	productos, err := os.ReadFile("./productos-1.csv")

	if err != nil {
		fmt.Println("Ejecutar primero el ejercicio 1")
		os.Exit(1)
	}

	tablaProductos := strings.Split(string(productos),"\n")

	for _, filaString := range tablaProductos {
		if filaString != ""{
			fila := strings.Split(filaString,",")
			fmt.Print(fila[0] + " \t ")
			fmt.Printf("\t%v", fila[1])
			fmt.Printf("\t%v\n", fila[2])
		}
	}
}
