package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information."
	credit := 12.45
	fmt.Printf("%T %T", text, credit) // output: string float64
}

func main() {
	tellMeTypes()
}
