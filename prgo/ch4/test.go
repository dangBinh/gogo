package main

import (
	"fmt"
)

type composer struct {
	name string
	birthYear int
}

type Product struct {
	name string
	price float64
}

type Point struct {
	x, y, z int
}

func (product Product) String() string {
	return fmt.Sprintf("%s (%.2f)", product.name, product.price)
}

func (point Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", point.x, point.y, point.z)
}

func main()  {
	// Con tro
	aton := composer{"Antonio", 1707}
	agnes := new(composer)
	agnes.name, agnes.birthYear = "Agnes", 1991
	julia := &composer{}
	julia.name, julia.birthYear = "Julia", 1819
  fmt.Println(aton, agnes, julia);

	// Duyet slice
	products := []*Product{{"Spanner", 3.99}, {"Wrench", 2.49}}
	fmt.Println(products)
	for _, product := range products {
		product.price += 0.50
	}
	fmt.Println(products)

	// Map
	populationForCity := map[string]int{"Hanoi" : 100000, "BacGiang" : 500}
	for city, population := range populationForCity {
		fmt.Printf("%-10s %8d\n", city, population)
	}
	// Map and struct
	triangle := make(map[*Point]string, 3);
	triangle[&Point{1, 2, 3}] = "alpha"
	triangle[&Point{4, 5, 7}] = "beta"
	fmt.Println(triangle);
}
