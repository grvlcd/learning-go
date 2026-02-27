package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age Pointer: ", agePointer)
	fmt.Println("Age Value: ", age)
	deductAge(agePointer)
	fmt.Println("Age New Value: ", age)
}

func deductAge(age *int) {
	*age = *age - 18
}
