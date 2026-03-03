package main

import "fmt"

func main() {
	hobbies := []string{"Learning", "Board Games", "Cooking"}
	fmt.Println("ALL: ", hobbies)

	fmt.Println("First: ", hobbies[0])

	fmt.Println("To Least: ", hobbies[1:])

	firstAndSecondHobbies := hobbies[:2]
	fmt.Println("First and Second: ", firstAndSecondHobbies)
	firstAndSecondHobbies = firstAndSecondHobbies[1:3]
	fmt.Println("Second and Last: ", firstAndSecondHobbies)

	courseGoals := []string{"Learn Go", "Write Backend"}
	fmt.Println("Course Goals: ", courseGoals)
	courseGoals[1] = "Write API"
	fmt.Println("New Second Index Goal: ", courseGoals)
	courseGoals = append(courseGoals, "Deploy to GCP")
	fmt.Println("Added Third Goal: ", courseGoals)
}

// Time to practice what you learned!

// DONE: 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// DONE: 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// DONE: 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// DONE: 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// DONE: 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// DONE: 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// TODO: 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
