package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2023, time.February, 01, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05"))
}
