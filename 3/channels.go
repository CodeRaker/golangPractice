package main

import "fmt"
import "./src/abdul"

func main() {
	channel1 := make (chan string)
	go abdul.Abdulize(channel1)
	printFromChannel(channel1)
}

func printFromChannel(ch chan string) {
	for i := 0; i < 100; i++ {
		msg := <- ch
		fmt.Printf(msg)

	}

}