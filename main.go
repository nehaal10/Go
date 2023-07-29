package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")
	//comma , of / err ok
	input, _ := reader.ReadString('\n') //we can assume that input is like try and _ is like catch
	fmt.Println("the ratings are : ", input)
	fmt.Printf("the type is : %T ", input)

	//buffio is the package --> bufio temperorly accumlates the the value buffer transferring it to the variable
	//NewReader is like scanner is java
	//os.stdin is saying u must take input
}
