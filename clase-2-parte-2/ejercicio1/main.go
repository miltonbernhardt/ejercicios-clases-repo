package main

import "fmt"

type Fecha struct {
	dia, mes, anio int
}

type Alumno struct {
	nombre       string
	apellido     string
	dni          string
	fechaIngreso Fecha
}

func main() {
	p1 := Alumno{"Milton", "Bernhardt", "9999999", Fecha{12, 01, 2021}}
	p2 := Alumno{nombre: "Nombre2", apellido: "Apellido2", dni: "88888888", fechaIngreso: Fecha{14, 02, 2021}}
	p3 := Alumno{"nombre3", "apellido3", "7777777", Fecha{13, 01, 2021}}

	alumnos := []Alumno{p1, p2, p3}

	for index, alumno := range alumnos {
		fmt.Printf("Alumno %v\n", index)
		fmt.Printf("\tNombre: %s\n", alumno.nombre)
		fmt.Printf("\tApellido: %s\n", alumno.apellido)
		fmt.Printf("\tDNI: %s\n", alumno.dni)
		fmt.Printf("\tFecha: %v\n\n", alumno.fechaIngreso.toString())
	}
}

func (fecha Fecha) toString() string {
	return fmt.Sprintf("%v/%v/%v", fecha.dia, fecha.mes, fecha.anio)
}
