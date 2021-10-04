package main

import "fmt"

//add function
func add(x, y int) int  {
	return x + y
}

//Swap Function
func swap(x,y string)(string, string) {
	return y,x
}

func main(){
	//Declare Variable
	a, b := swap("Hello", "World")
	fmt.Println("I love Burhan")
	fmt.Println(add(42, 24))
	fmt.Println(a,b)
}

