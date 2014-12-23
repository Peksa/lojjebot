package main

import (
	"fmt"
	irc "github.com/fluffle/goirc/client"
)

func main() {
	fmt.Println("Booting up!")

	// Creating a simple IRC client is simple.
	c := irc.SimpleClient("lojjebot")

	// Add handlers to do things here!
	// e.g. join a channel on connect.
	c.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			fmt.Println("Joining #lojjebot")
			conn.Join("#lojjebot")
		})

	// And a signal on disconnect
	quit := make(chan bool)
	c.HandleFunc("disconnected",
		func(conn *irc.Conn, line *irc.Line) {
			quit <- true
		})

	// Connect!
	c.ConnectTo("irc.motherland.nu")

	// Wait for disconnect
	<-quit
}
