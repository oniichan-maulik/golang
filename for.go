package main

import "fmt"

//for loop- only method for looping in go

func main() {

	//while loop with for
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// for loop
	// for i := 0; i <= 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	for i := range 3 {
		fmt.Println(i)
	}

}
