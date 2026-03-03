package main

import "fmt"

func main() {
	hobbies := []string{"Learning", "Board Games", "Cooking"}
	fmt.Println("ALL: ", hobbies)

	fmt.Println("First: ", hobbies[0])

	myLeastHobbies := hobbies[1:]
	fmt.Println("To Least: ", myLeastHobbies)
}

// Time to practice what you learned!

// DONE: 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// DONE: 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// TODO: 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// TODO: 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// TODO: 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// TODO: 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// TODO: 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
