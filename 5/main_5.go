package main

import "fmt"

func CheckUnion(a []int, b []int) (bool, []int) {
	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	result := make([]int, 0)

	for _, v := range b {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}
	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64, 64, 64, 43}
	b := []int{64, 2, 3, 43}
	fmt.Println(CheckUnion(a, b))

}
