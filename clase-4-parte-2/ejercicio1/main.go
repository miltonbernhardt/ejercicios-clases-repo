package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("ejecución finalizada")

	content, err := os.ReadFile("customers.txt")

	if err != nil {
		panic(fmt.Errorf("el archivo indicado no fue encontrado o está dañado: %v", err))
	}

	fmt.Printf("%v", content)

}
