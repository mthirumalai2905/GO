package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This package is about time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 10, 4, 20, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("2006-01-02 Monday 15:04:05"))
}
