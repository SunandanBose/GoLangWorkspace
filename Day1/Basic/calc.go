package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
	"reflect"
)

func main(){
	a := 10.0
	b := 3
	

	/*c := a + b this expression will give error as a is of type float64 and b is of type int
	It will give the error 'Invalid Expression'
	Also Go says that if you declare a variable you have to use it or else it will give an error.
	*/
	c := int(a) + b
	fmt.Println("C = ",c,"it's type is",reflect.TypeOf(c))
}