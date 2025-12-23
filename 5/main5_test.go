package main

import (
	"reflect"
	"testing"
)

func TestCheckUnion(t *testing.T) {
	a := []int{65, 3, 58, 678, 64, 64, 64, 43}
	b := []int{64, 2, 3, 43}
	expected := []int{64, 3, 43}

	ok, result := CheckUnion(a, b)

	if !ok {
		t.Fatalf("expected ok=true, got false")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestCheckUnionEmpty(t *testing.T) {
	a := []int{65, 3, 58, 678, 64, 64, 64, 43}
	b := []int{66, 2, 4, 42}
	expected := []int{}

	ok, result := CheckUnion(a, b)

	if ok {
		t.Fatalf("expected ok=false, got true")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}

}
