package main

import "fmt"

func main() {
	arrays()
	maps()

}

func arrays() {

	// Basic array usage
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Short syntax using implicit declaration
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Slice concept
	slice := []int{1, 2, 3}
	slice = append(slice, 4)

	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)
}

func maps() {
	m := map[string]int{"foo": 42, "bar": 47}

	fmt.Println(m)

	delete(m, "bar")

	fmt.Println(m)
}
