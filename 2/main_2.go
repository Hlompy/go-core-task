package main

import (
	"fmt"
	"math/rand"
)

func sliceExample(slice []int) []int {
	result := make([]int, 0)

	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result

}

func addElements(slice []int, n int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	result = append(result, n)
	return result
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func removeElement(slice []int, index int) []int {
	result := make([]int, 0)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func main() {

	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, 11))
	fmt.Println(copySlice(originalSlice))
	fmt.Println(removeElement(originalSlice, 0))
	fmt.Println(originalSlice)

}
