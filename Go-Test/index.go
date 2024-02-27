package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // Adds new element to the end
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8) // Keys are of type string, values are type uint8
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range myMap2 {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// No while loops in go, rather:
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// OR:
	var j int = 0
	for {
		if j >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	// Last way:
	for k := 0; k < 10; k++ {
		fmt.Println(k)
	}
}
