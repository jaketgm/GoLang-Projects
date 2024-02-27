package main

import (
	"fmt"
	"strings"
)

func main() {
	var myStr = []rune("rÉsumÉ")
	var index = myStr[1]
	fmt.Printf("%v, %T\n", index, index)
	for i, v := range myStr {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of the string is: %v", len(myStr))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = []string{"s", "u", "b", "s"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
