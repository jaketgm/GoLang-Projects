package main

import "fmt"

/*
	Consider 15
	Nums coprime to 15: 1,2,4,7,8,11,13,14 (gcd(15, num)=1)
	Modular inverses mod 15: 1,8,4,13,2,11,7,14
	Because
	1 * 1 = 1 mod 15 = 1
	2 * 8 = 16 mod 15 = 1
	4 * 4 = 16 mod 15 = 1
	7 * 13 = 91 mod 15 = 1
	11 * 11 = 121 mod 15 = 1
	14 * 14 = 196 mod 15 = 1

	I(n) is the largest m \in Z^{+} where m < n-1 such that the mod inv. of m mod n = m
	So I(15) = 11
	Also I(100) = 51 and I(7) = 1
	Find I(n) for 3\leq n \leq 2 x 10^7

	My approach:
		- seq of primes [i]
		- Need to find i * inv. = 1 mod (numOne)
	
	numCoprimeNums[i] * inv. \equiv 1 (mod numOne)
	numCoprimeNums[i] * inv = 1 + numOne y 
*/

func main() {
	var numOne int
	fmt.Println("Enter number one (Desired number to find coprimes for): ")
	fmt.Scanln(&numOne)

	coPrimeNums := []int{} 

	// Loop from 1 to numOne to find all coprime numbers within this range
	for numTwo := 1; numTwo <= numOne; numTwo++ {
		if findGCD(numOne, numTwo) == 1 {
			coPrimeNums = append(coPrimeNums, numTwo) // Append if coprime
		}
	}

	fmt.Println("Coprime numbers:", coPrimeNums)
}

func findGCD(numOne int, numTwo int) int {
	if numTwo == 0 {
		return numOne
	}
	return findGCD(numTwo, numOne%numTwo)
}

func findModInv(coPrimeNums []int, numOne int, numTwo int) []int {
	for numTwo := 1; numTwo <= numOne; numTwo++ {
		if 
	}
}
