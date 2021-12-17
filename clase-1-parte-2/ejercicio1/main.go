package main

import "fmt"

func main() {
	var palabra = "palabra"

	fmt.Println(len(palabra))

	for index := range palabra {
		fmt.Print(string(palabra[index])+ " - ")
	}

	fmt.Println("")
	fmt.Println("---------")
	//otra forma
	for _, letra := range palabra {
		fmt.Print(string(letra) + " - ")
	}
}
