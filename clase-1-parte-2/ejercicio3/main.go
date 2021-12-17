package main

import "fmt"

func main() {
	var edad int8 = 23
	var antiguedad int8 = 2
	var isEmpleado = true
	var sueldo float32 = 120000


	if edad <= 22 || !isEmpleado || antiguedad <= 1 {
		fmt.Println("El cliente no puede acceder al prestamo porque no cumple una de las condiciones")
	}else if sueldo <= 100000 {
		fmt.Println("El cliente puede acceder al prestamo pero se le cobrara un interes porque su sueldo es MENOR al estipulado")
	}else{
		fmt.Println("El cliente puede acceder al prestamo pero NO se le cobrara un interes porque su sueldo es MAYOR al estipulado")
	}
}
