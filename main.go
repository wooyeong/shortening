package main

import "fmt"

func server(ctx <-chan string, control chan bool) {
Server:
	for {
		select {
		case msg := <-ctx:
			if msg == "bye" {
				break Server
			} else {
				fmt.Printf("Server: msg: %s\n", msg)
			}

		default:
		}
	}

	control <- true
}

func client(ctx chan string, control chan bool) {
	var input string
	for {
		fmt.Printf("Client: Enter string: ")
		fmt.Scanln(&input)
		ctx <- input
		if input == "bye" {
			break
		}
	}

	control <- true
}

func main() {
	//q := make(chan message)

	q := make(chan string)
	control := make(chan bool)
	go server(q, control)
	go client(q, control)

	for count := 0; count < 2; count++ {
		<-control
	}
}
