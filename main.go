package main

import (
	"fmt"

	module "github.com/Anprazt/simple_module_project"
)

func main() {
	var stdn1 = module.Student{"Andi Pras", 12}
	stdn1.Greeting()

	var name = stdn1.GetNameAt(2)
	fmt.Println("Nama terakhir: ", name)
}
