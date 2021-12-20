package main

import "fmt"

type Matrix struct {
	values      []float64
	height      int
	width       int
	isQuadratic bool
	maxValue    float64
}

func main() {
	matrix1 := Matrix{height: 4, width: 2}
	matrix2 := Matrix{height: 2, width: 5}
	matrix3 := Matrix{}

	matrix1.Set(76,89,2,3)
	matrix1.Print()
	matrix2.Set(12,98 ,3,7 ,1,982,739,8,12,731)
	matrix2.Print()
	matrix3.Print()

}

func (matrix *Matrix) Set(values ...float64) {
	matrix.values = values
}

func (matrix Matrix) Print() {

	if matrix.height == 0 {
		matrix.height = 3 //default value of height
	}

	if matrix.width == 0 {
		matrix.width = 3 //default value of width
	}

	for height := 0; height < matrix.height; height++ {
		for width := 0; width < matrix.width; width++ {
			index := height*matrix.width + width
			if index < len(matrix.values) {
				fmt.Printf("%v ", matrix.values[index])
			} else {
				fmt.Printf("%v ", 0)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
