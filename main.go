package main

import "fmt"

func main() {
	hyundaiBuilder := GetBuilder("hyundai")
	bmwBuilder := GetBuilder("bmw")

	director := NewDirector(hyundaiBuilder)
	hyundaiCar := director.buildCar()

	fmt.Printf("Hyundai Car Name : %s\n", hyundaiCar.Name)
	fmt.Printf("Hyundai Car Producer : %s\n", hyundaiCar.Producer)
	fmt.Printf("Hyundai Car Product Region : %d\n", hyundaiCar.Price)

	director.setBuilder(bmwBuilder)
	bmwCar := director.buildCar()

	fmt.Printf("BMW Car Name : %s\n", bmwCar.Name)
	fmt.Printf("BMW Car Producer : %s\n", bmwCar.Producer)
	fmt.Printf("BMW Car Product Region : %d\n", bmwCar.Price)

}
