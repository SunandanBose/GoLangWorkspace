package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import "fmt"

func main(){
	first_rank := "39"
	secondRank := "614"

	if first_rank < secondRank {
		fmt.Println("firstRank is better")
	}else if first_rank > secondRank{
		fmt.Println("secondRank is fine")
	}else{
		fmt.Println("Both are fine")
	}
}