package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	message, _ := CreateMessage("Ken","Dear")
	Greet(message, display)
	Greet(message, displayLine)
	Test(message, displayLine)
}
func display(s string){
	fmt.Print("Calling display")
	fmt.Print(s)
}

func displayLine(s string){
	fmt.Println("Calling Display Line")
	fmt.Println(s)
}
/*METHOD 1*/
/*In this method Greet we are taking function as a parameter
*The function which is to be called is to be mentioned while calling the function in
Greet we will just give the function a name "do" we can give any name possible
and inside the Greet function we are calling the function we got as a argument */
func Greet(name string, do func(string)){
	do(name)
}

/*METHOD 2*/
type Printer func(string)()
/* Here type is a keyword used to create user defined data type
printer becomes the data type for a function which takes an argument*/
func Test(name string, done Printer){
	done(name);
}


func CreateMessage(name, greeting string)(message string, alternate string){
	message = greeting + "  " + name
	alternate = "Hi" + name
	return
}

