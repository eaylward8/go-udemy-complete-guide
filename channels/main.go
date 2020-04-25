package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string) // new channel that can communicate strings

	for _, l := range links {
		go checkLink(l, c) // allows main routine and child routine to communicate
	}

	// simple solution
	// listen for msg from channel and log it
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) // this is a blocking call
	// }

	// infinite loop, but with blocking HTTP requests
	// range on channel says, wait for value to be received, then assign it to var l
	for l := range c {
		// go checkLink(l, c) // first arg is allowed because it's a channel that produces strings

		// function literal that is called immediately
		// important to pass l in as value so the function literal is working with a copy of l (pass by value)
		// don't share variables between routines
		go func(lnk string) {
			time.Sleep(3 * time.Second)
			checkLink(lnk, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// don't care about response in this program
	_, err := http.Get(link) // http.Get is blocking code

	if err != nil {
		fmt.Println(link, "might be down!")
		// simple solution
		// c <- "Looks like it's down" // send msg to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// simple solution
	// c <- "It's up!" // send msg to channel

	c <- link
}
