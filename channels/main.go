package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	links := []string{
		"https://www.google.co.in",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.gitlab.com",
		"https://www.android.com",
	}

	c := make(chan string)

	for _, link := range links {
		go check(link, c)
	}

	for i := 0;i < len(links); i++ {
		fmt.Println(<- c)
	}


	//for {
	//	//fmt.Println(<-c)
	//	go checkLink(<-c, c)
	//}

	//for l := range c {
	//	go func(link string) {
	//		time.Sleep(5 * time.Second)
	//		checkLink(l, c)
	//	}(l)
	//}

}

func test() {
	c := make(chan string)
	c <- "hi"
}

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link
		os.Exit(1)
	}
	c <- link
	fmt.Println(link, " is up")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(err)
		c <- link
		os.Exit(1)
	}
	fmt.Println(link, " is up")
	c <- link
}
