package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	for timer := 10; timer >0; timer-- {
		if(timer %2 ==0){
			continue
		}
		fmt.Println(timer)
	}
}