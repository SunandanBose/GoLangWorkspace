package main


/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	lowest := findLowest(10,2,-1,0,1,14)
	fmt.Println(lowest)
}

// ... heps us to pass zero or n number of arguments
func findLowest(listNum...int)int{
	low :=listNum[0]
	for _,i := range listNum{
		if i< low{
			low = i
		}
	}
	return low
}