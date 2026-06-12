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


func addCar(brand, model string, year int, price float64) Car {
	return Car{
		Brand: brand,
		Model: model,
		Year:  year,
		Price: price,
	}
}
func main() {

	var brand string
	var model string
	var year int
	var price float64
	var myChoice string
	fmt.Println("======Car Inventory System=====")

	inventory := []Car{
		{"Toyota", "Corolla", 2020, 20000},
		{"Jeep", "Wrangler", 2021, 34000},
		{"Ford", "Mustang", 2021, 35000},
		{"Honda", "Civic", 2019, 18000},
		{"Mazda", "CX-5", 2022, 28000},
		{"Toyota", "Camry", 2022, 26000},
		{"BMW", "3 Series", 2021, 38000},
		{"Ford", "F-150", 2020, 38000},
		{"Hyundai", "Elantra", 2021, 19000},
		{"Honda", "Accord", 2021, 27000},
		{"Nissan", "Altima", 2019, 17500},
		{"Tesla", "Model Y", 2023, 48000},
		{"Toyota", "RAV4", 2021, 29000},
		{"Jeep", "Grand Cherokee", 2022, 43000},
		{"Ford", "Explorer", 2022, 42000},
		{"Honda", "CR-V", 2020, 28000},
		{"Mazda", "3", 2021, 23000},
		{"Chevrolet", "Suburban", 2021, 60000},
		{"BMW", "X5", 2023, 65000},
		{"Toyota", "Tacoma", 2023, 35000},
		{"Hyundai", "Tucson", 2022, 26500},
		{"Ford", "Escape", 2019, 22000},
		{"Nissan", "Rogue", 2022, 27000},
	}

	for {
		fmt.Println(`Menu: 1) Display Cars
      2) Add Car
      3) Search Cars
      4) Exit`)
		fmt.Scanln(&myChoice)

		switch myChoice {
		case "1":
			fmt.Println("--- List of all Cars ---")
			allCars(inventory)
			fmt.Println()
		case "2":
			fmt.Println("Enter Brand: ")
			fmt.Scanln(&brand)
			fmt.Println("Enter Model: ")
			fmt.Scanln(&model)
			fmt.Println("Enter Year: ")
			fmt.Scanln(&year)
			fmt.Println("Enter Price: ")
			fmt.Scanln(&price)
			newCar := addCar(brand, model, year, price)
			inventory = append(inventory, newCar)
			fmt.Println("Car added successfully!")
			fmt.Println()
		case "3":
			fmt.Println("Enter Brand: ")
			fmt.Scanln(&brand)
			fmt.Println("--- Search Results ---")
			results := filterByCarBrand(inventory, brand)
			if len(results) == 0 {
				fmt.Println("No cars found for that brand.")
			} else {
				for _, car := range results {
					fmt.Printf("%s %s (%d) - $%.2f\n", car.Brand, car.Model, car.Year, car.Price)
				}
			}
			fmt.Println()
		case "4":
			fmt.Println()
			fmt.Println("Exiting... Goodbye!")
			return

		}
	}

}