package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func DetectTypes(values []any) []string {
	types := make([]string, 0, len(values))
	for _, v := range values {
		types = append(types, fmt.Sprintf("%T", v))
	}
	return types
}

func ToString(values []any) []string {
	result := make([]string, 0, len(values))
	for _, v := range values {
		result = append(result, fmt.Sprintf("%v", v))
	}
	return result
}

func JoinStrings(values []string) string {
	result := ""
	for _, v := range values {
		result += v
	}
	return result
}

func StringToRunes(s string) []rune {
	return []rune(s)
}

func HashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	withSalt := string(runes[:mid]) + salt + string(runes[mid:])

	hash := sha256.Sum256([]byte(withSalt))
	return hex.EncodeToString(hash[:])
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	names := []string{
		"numDecimal",
		"numOctal",
		"numHexadecimal",
		"pi",
		"name",
		"isActive",
		"complexNum",
	}

	values := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	fmt.Println("Типы переменных:")
	types := DetectTypes(values)
	for i, v := range values {
		fmt.Printf(
			"%s = %v | type = %s\n",
			names[i],
			v,
			types[i],
		)
	}

	stringValues := ToString(values)
	joined := JoinStrings(stringValues)
	fmt.Println("\nОбъединённая строка:")
	fmt.Println(joined)

	runes := StringToRunes(joined)
	fmt.Println("\nСрез рун:")
	fmt.Println(runes)

	hash := HashWithSalt(runes, "go-2024")
	fmt.Println("\nSHA256 хэш с солью:")
	fmt.Println(hash)
}
