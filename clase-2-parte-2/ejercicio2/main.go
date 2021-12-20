package main

import "fmt"

type TipoProducto string

const (
	Pequeno TipoProducto = "Pequeño"
	Mediano TipoProducto = "Mediano"
	Grande  TipoProducto = "Grande"
)

type tienda struct {
	productos []Producto
}

type producto struct {
	tipoProducto TipoProducto
	nombre       string
	precio       float32
}

type Producto interface {
	CalcularCosto() float32
}

type Ecommerce interface {
	Total() float32
	Agregar(producto Producto)
}

func main() {
	tienda1 := NuevaTienda()
	tienda2 := NuevaTienda()

	tienda1.Agregar(NuevoProducto(Grande, "producto1", 4600))
	tienda1.Agregar(NuevoProducto(Mediano, "producto2", 5000))
	tienda1.Agregar(NuevoProducto(Pequeno, "producto3", 2200))

	tienda2.Agregar(NuevoProducto(Grande, "producto1", 3500))
	tienda2.Agregar(NuevoProducto(Mediano, "producto2", 1500))
	tienda2.Agregar(NuevoProducto(Pequeno, "producto3", 1200))

	fmt.Printf("Precio total de Tienda1: $%v\n", tienda1.Total())
	fmt.Printf("Precio total de Tienda2: $%v\n", tienda2.Total())
}

func NuevaTienda() Ecommerce {
	return &tienda{}
}

func NuevoProducto(tipoProducto TipoProducto, nombre string, precio float32) Producto {
	return producto{tipoProducto, nombre, precio}
}

func (tienda tienda) Total() float32 {
	var total float32 = 0
	for _, prod:= range tienda.productos {
		total += prod.CalcularCosto()
	}
	return total
}

func (tienda *tienda) Agregar(prod Producto) {
  tienda.productos = append(tienda.productos, prod)
}

func (producto producto) CalcularCosto() float32 {
	switch producto.tipoProducto {
	case Pequeno:
		return producto.precio
	case Mediano:
		return producto.precio * 1.03
	case Grande:
		return producto.precio * 1.06 + 2500
	}
	return 0
}
