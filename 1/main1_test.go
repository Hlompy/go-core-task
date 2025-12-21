package main

import "testing"

func TestDetectTypes(t *testing.T) {
	input := []any{42, 3.14, "go", true}
	types := DetectTypes(input)

	if len(types) != 4 {
		t.Fatalf("expected 4 types, got %d", len(types))
	}
}

func TestToString(t *testing.T) {
	input := []any{42, "go"}
	result := ToString(input)

	if result[0] != "42" || result[1] != "go" {
		t.Fatal("string conversion failed")
	}
}

func TestJoinStrings(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := JoinStrings(input)

	if result != "abc" {
		t.Fatalf("expected abc, got %s", result)
	}
}

func TestStringToRunes(t *testing.T) {
	runes := StringToRunes("go")

	if len(runes) != 2 {
		t.Fatal("rune conversion failed")
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("test")
	hash1 := HashWithSalt(runes, "go-2024")
	hash2 := HashWithSalt(runes, "go-2024")

	if hash1 != hash2 {
		t.Fatal("hash must be deterministic")
	}
}
