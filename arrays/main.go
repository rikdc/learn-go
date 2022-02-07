package main

import "fmt"

func main() {

	// Basic array usage
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Short syntax using implicit declaration
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

}
