package main

import (
	"../../pkg/restapi"
	"fmt"
	"time"
)

/*type Pr struct {
	func objprint() {
		fmt.Printf("hello from struct")
	}
}*/

/*
   Always have one function running in foreground, otherwise the application will end!
   Below two webservers are created. Just like that.
*/

func main() {
	//l := Pr
	//l.objprint()
	go printTime()
	go restapi.StartRouter(":8080")
	restapi.StartRouter(":8081")
}

func printTime() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hallo! This is ABDUL!\n")
		time.Sleep(2 * time.Second)
	}
}

// comments