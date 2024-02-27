package main

import (
	"fmt"
)

func main() {
	var p *int32 = new(int32) // Similar to malloc
	var i *int32 = new(int32)
	*p = 5
	*i = 6
	fmt.Printf("Original pointer (p): %v, original pointer (i): %v", *p, *i)
	swap(p, i)
	fmt.Printf("Swapped pointer (p): %v, swapped pointer (i): %v", *p, *i)

}

func swap(p *int32, i *int32) {
	var temp int32 = *p
	*p = *i
	*i = temp
}
