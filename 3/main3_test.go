package main

import "testing"

func TestAddAndGetValue(t *testing.T) {
	m := NewStringIntMap()

	m.Add("a", 1)

	value, ok := m.Get("a")
	if !ok {
		t.Fatalf("expected key 'a' to exist")
	}
	if value != 1 {
		t.Fatalf("expected value 1, got %d", value)
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()

	if m.Exists("a") {
		t.Fatalf("key should not exist yet")
	}

	m.Add("a", 10)

	if !m.Exists("a") {
		t.Fatalf("key should exist after Add")
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()

	m.Remove("a")

	_, ok := m.Get("a")

	if ok {
		t.Fatalf("key should not exist after Remove")
	}

}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	copyMap := m.Copy()

	if len(copyMap) != 2 {
		t.Fatalf("expected copy size 2, got %d", len(copyMap))
	}

	if copyMap["a"] != 1 || copyMap["b"] != 2 {
		t.Fatalf("copied values are incorrect")
	}

	copyMap["a"] = 100

	value, _ := m.Get("a")
	if value == 100 {
		t.Fatal("Copy() returned a map that shares memory with original")
	}
}
