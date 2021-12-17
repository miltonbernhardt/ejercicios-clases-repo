package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	cantEmpleadosViejos := 0

	for _, age := range employees {
		if age > 21{ cantEmpleadosViejos++ }
	}

	fmt.Println(employees)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

	fmt.Printf("cantidad de empleados mayores a 21 anos: %v", cantEmpleadosViejos)

}
