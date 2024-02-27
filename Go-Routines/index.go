package main

import (
	"fmt"
	"sync"
)

func findMin(arr []int, minCh chan int) {
	var minVal int = arr[0]
	for _, v := range arr[1:] {
		if v < minVal {
			minVal = v
		}
	}
	minCh <- minVal
}

func findEven(arr2 []int, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range arr2 {
		if v%2 == 0 {
			evenCh <- v
		}
	}
}

func splitIntoThree(arr2 []int) ([]int, []int, []int) {
	length := len(arr2)
	spliceSize := length / 3

	lastThird := spliceSize * 2
	if length%3 != 0 {
		lastThird += 1
	}

	first := arr2[:spliceSize]
	second := arr2[spliceSize:lastThird]
	third := arr2[lastThird:]

	return first, second, third
}

func main() {
	arr := []int{7, 2, 8, -9, 4, 0}

	minCh := make(chan int, 2)
	go findMin(arr[:len(arr)/2], minCh)
	go findMin(arr[len(arr)/2:], minCh)

	min1, min2 := <-minCh, <-minCh
	close(minCh)

	min := min1
	if min2 < min {
		min = min2
	}

	fmt.Printf("The minimum in the array is: %v\n", min)

	// Size 21 => 7 processes for 3 goroutines
	arr2 := []int{8, 2, 1, 7, 9, 10, 16, 9, 11, 13, 21, 22, 89, 90, 33, 12, 77, 78, 101, 11, 23}
	first, second, third := splitIntoThree(arr2)

	var wg sync.WaitGroup
	evenCh := make(chan int)

	wg.Add(3)
	go findEven(first, evenCh, &wg)
	go findEven(second, evenCh, &wg)
	go findEven(third, evenCh, &wg)

	go func() {
		wg.Wait()
		close(evenCh)
	}()

	var evenNums []int
	for num := range evenCh {
		evenNums = append(evenNums, num)
	}

	fmt.Println("Even numbers:", evenNums)
}
