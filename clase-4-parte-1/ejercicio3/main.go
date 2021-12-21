package main

import (
	"fmt"
)

func main() {
	var salary = 401
	checkSalary(salary)
}

func checkSalary(salary int) {
	if salary < 150000 {
		fmt.Println(fmt.Errorf("error: el mínimo imponible es de 150000 y el salaio ingresado es de: [%v]", salary))

	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
