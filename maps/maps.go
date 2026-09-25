package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps is associative data type
	// --hash, objectin java, dictionary in python
	// m := make(map[string]string)

	// setting elements
	// m["name"] = "golang"
	// m["area"] = "backend"
	// get element
	// fmt.Println(m["name"], m["area"])
	// if key does not exist in map, it returns zero

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// delete(m, "age")
	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 50, "phone": 3}
	// fmt.Println(m)

	// m := map[string]int{"price": 50, "phone": 3}
	// k, ok := m["phone"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not okay")
	// }

	m1 := map[string]int{"price": 50, "phone": 3}
	m2 := map[string]int{"price": 50, "phone": 8}
	fmt.Println(maps.Equal(m1, m2))
}
