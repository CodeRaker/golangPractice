package main

import (
	"../../pkg/restapi"
	"fmt"
	"time"
)

server1Port := 8080
server2Port := 8081
//server1Chan := make (chan int, 2) // create buffered channel with capacity of two
//server1Chan := make (chan int)

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
	go restapi.StartRouter(server1Port)
	restapi.StartRouter(server2Port)
}

func printTime() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hallo! This is ABDUL!\n")
		time.Sleep(2 * time.Second)
	}
}

// https://devhints.io/go