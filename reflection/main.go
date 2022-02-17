package main

import (
	"fmt"
	"reflecting/model"
)

func main() {
	fmt.Println("My Favourite Movie")

	myFav := model.NewMovie("Top Gun", model.R, 43.2)

	fmt.Printf("My favourite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %v\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")
	fmt.Printf("My favourite movie is %s\n", myFav.GetTitle())
}
