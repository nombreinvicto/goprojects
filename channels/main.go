package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a string channel
	c := make(chan string)

	//startTime := time.Now()

	for _, website := range websites {

		// pass the channel to the function
		// being run by a goroutine
		go checkLink(website, c)
	}

	// the moment we start waiting for messages
	// in a channel, its a blocking call until the
	// message is received
	//for i := 0; i < len(websites); i++ {
	//	fmt.Println(<-c)
	//}

	// iterating over channel yields messages as they come
	for l := range c {
		fmt.Println("value from channel: ", l)
		go func(lnk string) {
			time.Sleep(1 * time.Second)
			checkLink(lnk, c)
		}(l)
	}

	// calculate time elapsed
	//endTime := time.Now()
	//elapsed := endTime.Sub(startTime)
	//fmt.Printf("Took %v for all websites", elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v ERROR \n", link)
		return
	}
	fmt.Printf("%v OK \n", link)
	c <- link
}
