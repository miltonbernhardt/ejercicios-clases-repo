package main

import "fmt"

func main() {
	var precio float32 = 4000
	var descuento float32 = 12

	fmt.Println(precio * (100-descuento) / 100)
}
