package main

import "fmt"
import "reflect"

func main() {
	var version float32 = 0.1
	name := "John Doe"
	age := 20

	fmt.Println("Version:", version)

	fmt.Println("Hello, Sr.", name)
	fmt.Println("You are", age, "years old")

	fmt.Println("Type of var version is", reflect.TypeOf(version))
	fmt.Println("Type of var name is", reflect.TypeOf(name))
	fmt.Println("Type of var age is", reflect.TypeOf(age))

	fmt.Println("0 - Exit")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")

	var command int

	fmt.Scan(&command) // or fmt.Scanf("%d", &command)

	fmt.Println("The chosen command was", command)
}