package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import "fmt"

func main(){
	module := "Go Lang"
	ptr := &module

	fmt.Println("Memory address of module is ",ptr,"and its value is",*ptr)
	//& -> "address of" *ptr -> de-referencing ptr
	// the output of the above expression will be as given below
	//Memory address of module is  0xc00002e1d0 and its value is Go Lang
}