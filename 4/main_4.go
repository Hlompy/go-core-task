package main

import "fmt"

func Union1(slice1 []string, slice2 []string) []string {
	toRemove := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		toRemove[v] = struct{}{}
	}
	write := 0

	for _, v := range slice1 {
		if _, exists := toRemove[v]; !exists {
			slice1[write] = v
			write++
		}
	}

	return slice1[:write]
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Union1(slice1, slice2))
}
