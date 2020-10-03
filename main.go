package main

import "fmt"

func main() {
	var milikAinun = MilikAninun{
		Apples: 25,
		Cakes:  20,
	}

	fpb := calculateBoxFactor(&milikAinun)
	totalItemApple, totalItemCake := calculatePartOfItem(&milikAinun, fpb)
	fmt.Printf("Dari total apel sebanyak %d dan cake sebanyak %d , bisa dibagi menjadi %d kotak masing-masing terdiri dari %d aple dan %d cake\n",
		milikAinun.Apples, milikAinun.Cakes, fpb, totalItemApple, totalItemCake)
}

type MilikAninun struct {
	Apples int
	Cakes  int
}

func calculateBoxFactor(milikAinun *MilikAninun) int {
	apple := milikAinun.Apples
	cake := milikAinun.Cakes

	if cake > apple {
		apple = cake
		cake = apple
	}

	for i := cake; i > 0; i-- {
		if cake%i == 0 && apple%i == 0 {
			return i
		}
	}

	return 0
}

func calculatePartOfItem(milikAinun *MilikAninun, fpb int) (totalItemApple int, totalItemCake int) {
	totalItemApple = milikAinun.Apples / fpb
	totalItemCake = milikAinun.Cakes / fpb
	return
}
