package main

import "fmt"

func main() {
	var sueldo float32 = 60000

	impuesto := calcularImpuesto(sueldo)

	fmt.Printf("%f",impuesto)
}

func calcularImpuesto(sueldo float32) (impuesto float32) {
	if sueldo >= 50000 && sueldo < 150000 {
		return sueldo * 0.17
	} else if sueldo >= 150000 {
		return sueldo * 0.27
	}
	return 0
}
