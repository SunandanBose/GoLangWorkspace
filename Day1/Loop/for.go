package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
	"time"
)

func main(){
	for timer := 10; timer >0; timer-- {
		if(timer ==5){
			break
			fmt.Println("Done...")
		}
		fmt.Println(timer)
		//time.Sleep(2*time.Second) will actually take 2 seconds and then execute the code
		time.Sleep(2*time.Second)
	}
}