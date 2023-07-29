package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to string conversion"
	fmt.Println(welcome)
	//craeting a buffer to take input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter pizza ratings")
	input, _ := reader.ReadString('\n')
	fmt.Println("the ratings are: ", input)
	fmt.Printf("the type is: %T \n", input)

	//now we are going to convert
	//as we know that the input type is string
	inputTrimed := strings.TrimSpace(input)         //removed \n
	num, err := strconv.ParseFloat(inputTrimed, 64) // Parsefloat returns 64bit float
	//NOTE THE BELOW COMMENTS WERE MADE BEFOR TRIMMING INPUT
	//while trying to convert to float64 we are getting this error strconv.ParseFloat: parsing "5\n": invalid syntax
	//this is because in readstring comand we gave termination condition as /n , hence /n is added to string

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("the type of converted input is: %T \n", num)
		addedNum := num + 1
		fmt.Println("the number is: ", addedNum)
	}
}
