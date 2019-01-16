package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	message, alternate := CreateMessage("Ken","Dear")
	fmt.Println(message)
	fmt.Println(alternate)
}
/*This is a named return type here and here in the return we are not mentioning anything
so Go will automatically take the variable message and alternate as mentioned while func declaration line
so return will automatically take it.*/
func CreateMessage(name, greeting string)(message string, alternate string){
	message = greeting + "  " + name
	alternate = "Hi" + name
	return
}

