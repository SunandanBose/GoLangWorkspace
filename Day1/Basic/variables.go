package main

/**
 * Created by sunandanbose on 16/01/19.
 */

import (
	"fmt"
	"reflect"
)
//reflect module is to know the type of the variable
/* variables needs to be declared using var keyword at a package level*/

var(
	//in Go first the variable name is written and then the type unlike other programming languages
	course string
	name string
	//this can also be written as course, name string
	module float64

	/* you can also declare the variable as mentioned below where you need
	not write the data type. It will be inferred by GO programming language*/
	first_name, age, salary = "Sunandan",25,1000000.0
	//here first_name will be by default taken as String, age as int and salary as float64
)

func main(){
	fmt.Println("Name is set to be ",name,"and is of type ",reflect.TypeOf(name))
	fmt.Println("Course is set to be ",course,"and is of type ",reflect.TypeOf(course))
	fmt.Println("Module is set to be ",module,"and is of type ",reflect.TypeOf(module))
	fmt.Println("First Name",first_name,"it is of type ",reflect.TypeOf(first_name))
	fmt.Println("Age",age,"it is of type ",reflect.TypeOf(age))
	fmt.Println("Salary",salary,"it is of type ",reflect.TypeOf(salary))
}