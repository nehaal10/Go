package main

import "fmt"

var message1 string = "hfgg"

const token float32 = 1.023

func main() {
	var username string = "abcdefghijklmnop"
	fmt.Println(username)
	fmt.Printf("the type is : %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("the type is : %T \n", isLogged)

	var num int = 20
	fmt.Println(num)

	message := "nehaal"
	fmt.Println(message)

	fmt.Println(message1)

	fmt.Println(token)
	message1 = "xyz"
	fmt.Println(message1)

}
