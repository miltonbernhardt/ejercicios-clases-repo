package main

import (
	"fmt"
	"io"
	"os"
)

type Producto struct {
	id       int
	precio   float32
	cantidad int
}

func main() {
	p1 := Producto{23, 45.50, 41}
	p2 := Producto{52, 56.35, 21}
	p3 := Producto{31, 230.75, 45}
	var productos = []Producto{p1, p2, p3}

	stringProd := "ID,Precio,Cantidad\n"
	stringProd += fmt.Sprint(p1.id, ",", p1.precio, ",", p1.cantidad, "\n")
	stringProd += fmt.Sprint(p2.id, ",", p2.precio, ",", p2.cantidad, "\n")
	stringProd += fmt.Sprint(p3.id, ",", p3.precio, ",", p3.cantidad, "\n")

	_ = os.WriteFile("./productos-1.csv", []byte(stringProd), 777)


	buffer2, _ := os.Create("./productos-2.csv")
	_, _ = io.WriteString(buffer2, "ID,Precio,Cantidad\n")
	for _, producto := range productos {
		_, _ = io.WriteString(buffer2, fmt.Sprint(producto.id, ",", producto.precio, ",", producto.cantidad, "\n"))
	}
	_ = buffer2.Close()

}
