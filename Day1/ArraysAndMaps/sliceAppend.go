package main

/**
 * Created by sunandanbose on 16/01/19.
 */


import (
	"fmt"
)

func main(){
	mySlice := make([]int,1,4)
	fmt.Println(mySlice)

	for i:=1;i<17;i++{
		mySlice = append(mySlice,i)
		fmt.Printf("Capacity = %d Value = %d at index = %d", cap(mySlice),mySlice[i],i)
		
	}
}