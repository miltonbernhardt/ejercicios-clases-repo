package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Client struct {
	file      int
	firstname string
	lastname  string
	dni       string
	phone     string
	address   string
}

func main() {
	defer fmt.Println("Fin de la ejecución")
	defer fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	defer fmt.Println("No han quedado archivos abiertos")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()


	client := Client{}
	file, err := assignFile()
	if err != nil {
		panic(err)
	}
	client.file = file



	existsClient := checkIfClientExist(client)
	fmt.Printf("the client exists: %v", existsClient)

	validClient, err := validateClient(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("the client is valid: %v", validClient)

	fmt.Printf("Client: %v", client)
}

func assignFile() (int, error) {
	rand.Seed(time.Now().UnixNano())
	file := rand.Intn(100-1) + 1

	if file < 30 {
		return 0, errors.New("error: invalid user file")
	}

	return file, nil
}

func checkIfClientExist(client Client) bool {
	content, err := os.ReadFile("customers.txt")

	if err != nil {
		panic(fmt.Errorf("el archivo indicado no fue encontrado o está dañado: %v", err))
	}

	for _, cli := range content {
		if client.equals(string(cli)) {
			return true
		}

	}

	return false
}

func validateClient(client Client) (bool, error){
	if client.firstname != "" {
		return false, errors.New("the attribute 'firstname' contains a invalid value")
	}

	if client.lastname != "" {
		return false, errors.New("the attribute 'lastname' contains a invalid value")
	}

	if client.address != "" {
		return false, errors.New("the attribute 'address' contains a invalid value")
	}

	if client.dni != "" {
		return false, errors.New("the attribute 'dni' contains a invalid value")
	}

	if client.phone != "" {
		return false, errors.New("the attribute 'phone' contains a invalid value")
	}

	return true, nil
}

// https://stackoverflow.com/questions/50660200/golang-convert-bytes-array-to-struct-with-a-field-of-type-byte
func (client Client) equals(dni string) bool {
	return client.dni != dni
	//if client.firstname != client2.firstname {
	//	return false
	//}
	//
	//if client.lastname != client2.lastname {
	//	return false
	//}
	//
	//if client.address != client2.address {
	//	return false
	//}
	//
	//if client.dni != client2.dni {
	//	return false
	//}
	//
	//if client.phone != client2.phone {
	//	return false
	//}
	//
	//return true
}
