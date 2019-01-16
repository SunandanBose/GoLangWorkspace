package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	//make([]Type, len, cap)
	myCourse := make([]string,5,10)
	fmt.Printf("Length = %d Capacity = %d", len(myCourse), cap(myCourse))
	fmt.Println(myCourse)

	mylang := []string{"C","Go","C++"}
	fmt.Printf("Length = %d Capacity = %d", len(mylang), cap(mylang))

	myArray := []int{1,2,4,5,6,7,8,9,0}
	fmt.Println(myArray)
	fmt.Println(myArray[5])
	myArray[2] = 100
	fmt.Println(myArray)
	mySlice := myArray[2:5]
	fmt.Println(mySlice)

	for _,i := range mySlice{
		fmt.Println(i)
	}

	//append slice to slice
	newSlice := []int{500,200,700}
	mySlice = append(mySlice,newSlice...)
	fmt.Println(mySlice)
}


/* A Slice is a segment of an array. Slices build on arrays and provide more power, flexibility, 
and convenience compared to arrays.
Just like arrays, Slices are indexable and have a length. But unlike arrays, they can be resized.
Internally, A Slice is just a reference to an underlying array. 
slice is created with the help of make function*/