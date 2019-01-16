package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import "fmt"
func main(){
	fmt.Println("Enter a number");
	var input float64
	fmt.Scanf("%f",&input)
	output := input*2
	fmt.Println("Output : ",output)
}