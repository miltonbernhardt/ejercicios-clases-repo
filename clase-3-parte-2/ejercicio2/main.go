package main

import "fmt"

type Product struct {
	Name     string
	Price    float32
	Quantity int
}

type User struct {
	Firstname string
	Lastname  string
	Email     string
	Products  []*Product
}

func main() {
	user := User{"Milton", "Bernhardt", "milton.bernhardt@mercadolibre.com", []*Product{}}
	var p1 = new(Product)
	p1 = newProduct("producto1", 456)

	addProduct(&user, p1, 78)
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", user.toString())
	fmt.Printf("%v\n", *user.Products[0])

	deleteProducts(&user)
	fmt.Printf("%v\n", user)
}

func newProduct(name string, price float32) *Product {
	p1 := Product{name, price, 0}
	return &p1
}

func addProduct(user *User, product *Product, quantity int) {
	product.Quantity = quantity
	user.Products = append(user.Products, product)

	product.Name = "Nombre cambiado"
}

func deleteProducts(user *User) {
	user.Products = []*Product{}
}

func (user User) toString() string {
	filaString := fmt.Sprintf("Nombre: %v %v - Email: %v", user.Firstname, user.Lastname, user.Email)
	for i, product := range user.Products {
		filaString += fmt.Sprintf("\n\tProducto %v: \n\t\tNombre producto: %v \n\t\tPrecio: %v \n\t\tCantidad: %v ", i, product.Name, product.Price, product.Quantity)
	}
	return filaString
}
