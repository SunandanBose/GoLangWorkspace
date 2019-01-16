package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
	"strings"
)
/*function OverLoading is not possible in GO Lang*/
func main(){
	module := "Functions"
	author := "Ken Thompson"

	fmt.Println(converter(module,author))
}

/*Like in Java we write function then return type and then the function name
but in Go we write func keyword then the function name parameters and then the return type.
Note : we can return multiple values in Go Lang unlike Java*/
func converter(module string, author string)(s1, s2 string){
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module,author
}