package main

import (
	"fmt"
)

type Product struct {
	name     string
	price    float64
	quantity int
}

type Service struct {
	name          string
	price         float64
	minutesWorked int
}

type Maintenance struct {
	name  string
	price float64
}

func main() {

	products := []Product{{"producto1", 600, 40}, {"producto2", 450, 21}, {"producto3", 720, 32}}
	services := []Service{{"service1", 400, 200}, {"service2", 500, 230}, {"service3", 100, 400}}
	maintenances := []Maintenance{{"maintenance1", 400}, {"maintenance2", 500}, {"service3", 100}}

	priceProducts := make(chan float64)
	priceServices := make(chan float64)
	priceMaintenances := make(chan float64)

	go addProducts(products, priceProducts)
	fmt.Printf("sum of product prices = %v\n", <-priceProducts)

	go addServices(services, priceServices)
	fmt.Printf("sum of service prices = %v\n", <-priceServices)

	go addMaintenance(maintenances, priceMaintenances)
	fmt.Printf("sum of maintenance prices = %v\n", <-priceMaintenances)
}

func addProducts(products []Product, price chan float64) {
	var totalPrice float64 = 0
	for _, product := range products {
		totalPrice += product.price
	}
	price <- totalPrice
}

func addServices(services []Service, price chan float64) {
	var totalPrice float64 = 0
	for _, service := range services {
		if (service.minutesWorked / 30) < 30 {
			totalPrice += service.price
		} else {
			totalPrice += float64(service.minutesWorked/30) * service.price
		}
	}
	price <- totalPrice
}

func addMaintenance(maintenances []Maintenance, price chan float64) {
	var totalPrice float64 = 0
	for _, maintenance := range maintenances {
		totalPrice += maintenance.price
	}
	price <- totalPrice
}
