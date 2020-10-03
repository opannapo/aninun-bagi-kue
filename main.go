package main

import (
	"fmt"
)

func main() {
	apples := 20
	cakes := 20
	fpb := CalculateBoxFactor(apples, cakes)
	totalItemApple, totalItemCake := CalculatePartOfItem(apples, cakes, fpb)
	fmt.Printf("Dari total apel sebanyak %d dan cake sebanyak %d , bisa dibagi menjadi %d kotak masing-masing terdiri dari %d aple dan %d cake\n",
		apples, cakes, fpb, totalItemApple, totalItemCake)
}

func CalculateBoxFactor(apples int, cakes int) int {
	if cakes > apples {
		apples = cakes
		cakes = apples
	}

	for i := cakes; i > 0; i-- {
		if cakes%i == 0 && apples%i == 0 {
			return i
		}
	}

	return 0
}

func CalculatePartOfItem(apples int, cakes int, fpb int) (totalItemApple int, totalItemCake int) {
	totalItemApple = apples / fpb
	totalItemCake = cakes / fpb
	return
}
