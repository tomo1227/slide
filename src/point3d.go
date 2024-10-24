package main

import "fmt"

type Point3D[E any] = struct{ x, y, z E }

func main() {
	p := Point3D[float64]{1.0, 2.71828, 3.14159}
	fmt.Println(p) // {1 2.71828 3.14159}
}
