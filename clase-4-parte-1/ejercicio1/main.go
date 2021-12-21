package main

import (
	"fmt"
)

type MyError struct{}

func main() {
	var salary = 401

	fmt.Printf(checkSalary(salary))
}

func checkSalary(salary int) string {
	err := MyError{}
	if salary < 150000 {
		return err.Error()
	} else {
		return "Debe pagar impuesto"
	}
}

func (myError *MyError) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}
