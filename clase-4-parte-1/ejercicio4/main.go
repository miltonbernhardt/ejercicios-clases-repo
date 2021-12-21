package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var salaries []float32

	monthsWorked := rand.Intn(12 - 9) + 9

	for i := 0; i < monthsWorked; i++ {
		hoursWorked := rand.Intn(110 - 79) + 80
		hourValue := rand.Intn(95 - 80) + 80
		salary, err := calculateSalary(hoursWorked, hourValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		salaries = append(salaries, salary)
	}

	aguinaldo, err := calculateAguinaldo(salaries)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("The employee has next salaries => %v\n", salaries)
	fmt.Printf("The employee has aguinaldo => %v\n", aguinaldo)

}

func calculateSalary(hoursWorked int, hourValue int) (float32, error) {
	if hoursWorked < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales - horas trabajadas: %v", hoursWorked)
	}

	var salary = float32(hourValue * hoursWorked)
	if salary <= 150000 {
		salary = salary * 0.9
	}

	return salary, nil
}

func calculateAguinaldo(salaries []float32) (float32, error) {

	if len(salaries) == 0 {
		return 0, errors.New("error: el trabajador debe de haber trabajado un mes para percibir el aguinaldo")
	}

	var maxSalary float32 = 0
	for _, salary := range salaries {
		if maxSalary < 0 {
			return 0, fmt.Errorf("error: el salario del trabajador no debe ser negativo - %v", salary)
		}

		if salary > maxSalary {
			maxSalary = salary
		}
	}

	return (maxSalary / 12) * float32(len(salaries)), nil
}
