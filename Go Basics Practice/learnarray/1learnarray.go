package learnarray

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.


type product struct {
	title string
	id int
	price float64
}

func Learnarray() {
	//1. Lab show above.
	hobbies := []string{"dancing", "singing", "reading"}
	fmt.Println(hobbies)
	
	//2.
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3.
	// sliceHobbies := hobbies[0:2]
	sliceHobbies := hobbies[:2]
	fmt.Println(sliceHobbies)


	//4.
	newSliceHobbies := sliceHobbies[0:cap(sliceHobbies)]
	fmt.Println(newSliceHobbies)

	//5
	goals := []string{"Developer", "Philanthropist"}
	goals[1] = "Enterprenuer"
	
	//6.
	newGoals := append(goals, "Teacher");
	fmt.Println(newGoals)

	//7.
	product1 := product{
		"Toothbrush",
		1,
		56.32,
	}

	product2 := product{
		"Toothpaste",
		2,
		55.43,
	}

	productDetails := [] product{product1, product2}
	
	product3 := product{
		"BrushCap",
		3,
		20.43,
	}

	newProductDetails := append(productDetails, product3)
	fmt.Println(newProductDetails)
}