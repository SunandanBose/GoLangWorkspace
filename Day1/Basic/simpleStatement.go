package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import "fmt"

func main(){
	
//by this way we can five values to both the variables but there scope will only be limited
//only to the if else loop only
	if first_rank, secondRank := 40,650; first_rank < secondRank {
		fmt.Println("firstRank is better")
	}else if first_rank > secondRank{
		fmt.Println("secondRank is fine")
	}else{
		fmt.Println("Both are fine")
	}
	/*fmt.Println(first_rank) yjis stateent eill give error undefined: first_rank as our scope of first_rank 
	was only limited to if else*/
}