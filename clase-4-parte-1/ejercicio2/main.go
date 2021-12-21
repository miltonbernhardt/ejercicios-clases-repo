package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary = 401
	checkSalary(salary)
}

func checkSalary(salary int) {
	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))

	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
