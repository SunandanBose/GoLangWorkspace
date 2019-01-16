package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){

	/*In go we don't need to write break after each case 
	it is applied automatically in Go Lang*/
	switch "docker"{
		case "linux":
			fmt.Println("linux installed")
		case "windows":
			fmt.Println("windows installed")
			falthrough
	/*In other programming languages if we don't write break after case then it will keep on executing
	the next case unless it finds break but as we know Go Lang provide by default break statement 
	we need not put it. So to get the fallthrough feature like other language then we can use
	the keyword falthrough*/		
		case "mac":
			fmt.Println("mac installed")
		case "docker":
			fmt.Println("docker installed")
		default:
			fmt.Println("no match found")
	}
}