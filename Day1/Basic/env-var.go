package main

import (
	"fmt"
	"os"
)

func main(){
	/* To declare a variable inside a function we can declare it like this
	name := "Sunandan"
	As we know we need not declare the datatype in GO Language it is inferred
	automatically
	! := will only work inside a function and not on a package level
	if you want to give a data type to a variable just write it like this
	name string
	*/

	name := os.Getenv("USERNAME")
	fmt.Println("User name: ",name)
	root := os.Getenv("GOROOT")
	fmt.Println("root: ", root)
}