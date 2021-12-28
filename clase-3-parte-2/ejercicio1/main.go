package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Age       int
	Email     string
	Password  string
}

func main() {
	user := User{"Milton", "Bernhardt", 24, "milton.bernhardt@mercadolibre.com", "password"}
	user.changeName("nombre", "apellido")
	user.changeAge(25)
	user.changeEmail("mail@mail.com")
	user.changePassword("password@password")

	fmt.Printf("%v\n", user)
}

func (user *User) changeName(firstname string, lastname string) {
	user.Firstname = firstname
	user.Lastname = lastname
}

func (user *User) changeAge(age int) {
	user.Age = age
}

func (user *User) changeEmail(email string) {
	user.Email = email

}

func (user *User) changePassword(password string) {
	user.Password = password
}
