package main

type StringIntMap struct {
	m map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		m: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.m[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.m, key)
}

func (m *StringIntMap) Copy() map[string]int {
	result := make(map[string]int)
	for k, v := range m.m {
		result[k] = v
	}
	return result
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.m[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.m[key]
	return v, ok
}
