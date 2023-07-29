package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")
	currentTime := time.Now()
	//the output will be something like "2023-07-29 18:10:58.765797 +0530 IST m=+0.000198793"
	fmt.Println(currentTime)
	//to get specifics like only date
	date := currentTime.Format("01-02-2006")
	fmt.Println("current date is: ", date)
	time := currentTime.Format("15:04:05")
	fmt.Println("current time is: ", time)
	//both date and time combined
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 monday"))
	fmt.Printf("date type: %T \n ", date)
	fmt.Println(date + " " + time)
}
