package main

import (
	"EDD/src/1_functions/calculator/operation"
	"fmt"
	"log"
)

func add(summandA int16, summandB int16) int16 {
	return summandA + summandB
}

func main() {
	sum := add(10, 20)
	fmt.Println("Result-Sum: ", sum)

	quotient := operation.Division(10, 20)
	fmt.Println("Result-Quotient: ", quotient)

	quotient2, err := operation.DivisionWithNullCheck(30, 0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Result-Quotient: ", quotient2)
}
