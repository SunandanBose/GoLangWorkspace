package main

import (
	"fmt"
)

func main(){
	dictData := make(map[string]int)
	dictData["Ken"] = 50
	dictData["Robert"] = 45
	dictData["Robert"] = 55
//composite iteral form
	nameAge := map[string]int{
		"Akash" :22,
		"Bhushan" : 25,
		"Amar" : 28,  
	}
	fmt.Printf("Dict data %v",dictData)
	fmt.Println("name Age : ",nameAge)

	for key,value := range nameAge{
		fmt.Printf("Key = %v ,  value = %v", key,value)
	}

	nameAge["Sunandan"] = 23
	nameAge["Akash"] = 42
	delete(nameAge,"Bhushan")
	fmt.Println(nameAge)

	fmt.Println("After Manipulation")
	 
	for key,value := range nameAge{
		fmt.Printf("Key = %v ,  value = %v", key,value)
	}

}