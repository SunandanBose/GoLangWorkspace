package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
	"os"
)

func main(){
/*os.Open actually gives the handle of the file present in the particular path
but if the path is not there the error will be save in err named variable
but if we need the handle of the file we can give handle,err
but here we are not interested in the handle so we gave _ so that it is ignored*/

	_,err := os.Open("c:\\go\\calc.go")
	if err != nill{
		fmt.Println(err)
	}
}