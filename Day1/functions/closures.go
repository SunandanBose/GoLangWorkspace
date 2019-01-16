package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	number := 10
	squareNum := func()(int){
		number *= number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())
}