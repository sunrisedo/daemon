package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println("now cnt:", i)
		log.Println("now cnt:", i)
		time.Sleep(time.Second * 10)
	}
	select {}
}
