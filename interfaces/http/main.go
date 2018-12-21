package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (l logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)

}

func print(a ...interface{})  {
	fmt.Println(a)
}