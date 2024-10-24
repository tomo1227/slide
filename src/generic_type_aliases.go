package main

type Map[K comparable, V any] mapImplementation

type Set[K comparable] = Map[K, bool]

func main() {
	var m Map[int, bool]
	var s Set[int]
	s = m
	m = s
}

type mapImplementation struct{}
