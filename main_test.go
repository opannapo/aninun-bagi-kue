package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calculate_Box_Factor_Test_25_20(t *testing.T) {
	apples := 25
	cakes := 20
	expectFpb := 5
	fpb := CalculateBoxFactor(apples, cakes)
	assert.Equal(t, expectFpb, fpb)
}
func Test_Calculate_Part_Of_Item_Test_25_20(t *testing.T) {
	apples := 25
	cakes := 20
	fpb := CalculateBoxFactor(apples, cakes)

	expectAppleItem := 5
	expectCakeItem := 4
	totalItemApple, totalItemCake := CalculatePartOfItem(apples, cakes, fpb)

	assert.Equal(t, expectAppleItem, totalItemApple)
	assert.Equal(t, expectCakeItem, totalItemCake)
}

func Test_Calculate_Box_Factor_Test_80_20(t *testing.T) {
	apples := 80
	cakes := 20
	expectFpb := 20
	fpb := CalculateBoxFactor(apples, cakes)
	assert.Equal(t, expectFpb, fpb)
}
func Test_Calculate_Part_Of_Item_Test_80_20(t *testing.T) {
	apples := 80
	cakes := 20
	fpb := CalculateBoxFactor(apples, cakes)

	expectAppleItem := 4
	expectCakeItem := 1
	totalItemApple, totalItemCake := CalculatePartOfItem(apples, cakes, fpb)

	assert.Equal(t, expectAppleItem, totalItemApple)
	assert.Equal(t, expectCakeItem, totalItemCake)
}
