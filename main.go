package main

import (
	"fmt"

	module "github.com/Anprazt/simple_module_project/v2"
)

func main() {
	//function
	module.About("SD Gandasari 1")

	//struct method
	var stdn1 = module.Student{"Andi Prasetyo Wicaksono", 12}
	stdn1.Greeting()

	var name = stdn1.GetNameAt(3)
	fmt.Println("Nama terakhir: ", name)

	//anonymous function
	var status = func(name string) bool {
		switch name {
		case "Andi":
			return true
		case "Prasetyo":
			return true
		case "Wicaksono":
			return true
		default:
			return false
		}
	}
	module.User(name, status)

	//struct
	module.ProfileTeacher()

	//anonymous struct
	module.AnonymousOrtu()

	//function return value
	var x = 5
	var y = 9
	fmt.Printf("Nilai x = %d,\nNilai y = %d.\nHasil Penjumlahannya adalah = %d\n", x, y, module.Sum(x, y))

	//function multiple return value
	var diameter float64 = 21
	var luas, keliling = module.CalculateCircle(diameter)
	fmt.Printf("Luas lingkaran \t\t: %2.f \n", luas)
	fmt.Printf("Keliling lingkaran \t: %2.f \n", keliling)

	//interface
	var bangunDatar module.Hitung = module.Persegi{10.0}
	fmt.Println("Hitung Persegi")
	fmt.Println("Luas        :", bangunDatar.Area())
	fmt.Println("Keliling    :", bangunDatar.Circumference())
}
