// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	// Declare the pointer. This will be unintialized.
	//var firstName *string

	// This will error out.
	// firstName = "Arthur"

	// Dereferencing the pointer will also generate an error.
	// *firstName = "Arthur"

	// Correct: Initialize at create
	// Declare the pointer. This will be initialized.
	var firstName *string = new(string)
	*firstName = "Arthur"

	fmt.Println(*firstName) // Remember to dereference!

}
