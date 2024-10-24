package main

type Map[K comparable, V any] mapImplementation

type Set[K comparable] = Map[K, bool]

type integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type IntSet[K integers] = Set[K]

func main() {
	var m Map[int, bool]
	var s IntSet[int]
	s = m
	m = s
}

type mapImplementation struct{}
