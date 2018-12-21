package main

func main() {
	c := make(chan string)
	c <- "Test"
	//fmt.Println(<-c)
}

func testMethod(c chan string)  {
	c <- "Test"
}



