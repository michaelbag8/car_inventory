package main

import (
	"fmt"
)

type Car struct {
	Brand string
	Model string
	Year  int
	Price float64
}

func allCars(inventory []Car) {
	for _, cars := range inventory {
		fmt.Printf("%s %s (%d)-$%.2f\n", cars.Brand, cars.Model, cars.Year, cars.Price)
	}

}

func filterByCarBrand(inventory []Car, brand string) []Car {
	var result []Car
	for _, cars := range inventory {
		if cars.Brand == brand {
			result = append(result, cars)
		}
	}
	return result
}
