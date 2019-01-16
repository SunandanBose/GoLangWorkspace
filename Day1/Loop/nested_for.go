package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	first := []string{"Go","C++","Python","Kotlin"}
	second := []string{"Go","Java","Python","Anaconda"}
	/*In Go array is declared in the above manner
	like in java we use forEach loop to iterate array of objects in a similar way we are using it in 
	GO lang for index,i := range courses here i contains the valueof the array courses and amd index 
	caontains the index location of that array.
	if we don't want the index then we can simply put _ in place of index*/


	for _,i := range first{
		for _,j := range second{
			if i==j{
				fmt.Println("You opted ", i , " language twice")
			}
		}
	}
}

/*Output
You opted  Go  language twice
You opted  Python  language twice
*/