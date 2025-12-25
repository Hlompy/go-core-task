package main

import (
	"reflect"
	"testing"
)

func TestUnion1(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := Union1(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}

}

func TestUnion1NoRemovals(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"x", "y"}

	result := Union1(slice1, slice2)

	if !reflect.DeepEqual(result, slice1) {
		t.Fatalf("expected %v, got %v", slice1, result)
	}
}

func TestUnion1AllRemoved(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b"}

	result := Union1(slice1, slice2)

	if len(result) != 0 {
		t.Fatalf("expected empty slice, got %v", result)
	}
}
