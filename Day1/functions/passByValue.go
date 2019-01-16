package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
)

func main(){
	language := "Go Lang"
	name := "Ken Thompson"

	fmt.Println("Hi", name,"you have developed ", language)
	/*Her we are passing the Address*/
	changeLanguage(&language)
	fmt.Println("Hi", name,"you have developed ", language)
}
/*he we are fetching it with the help of pointers*/
func changeLanguage(language *string) string{
	*language = "C language"
	fmt.Println("Changing language to", *language)
	return *language
}