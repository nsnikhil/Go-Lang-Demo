package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//noinspection GoStructTag
type responseGithub struct {
	UserId    int    `json:"userId'`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//noinspection GoUnhandledErrorResult
func main() {

	fileName := "temp.txt"

	url := fmt.Sprintf("%s%d", "https://jsonplaceholder.typicode.com/todos/", getInput())

	c := make(chan *http.Response)

	go getResponse(url, c)

	u := new(responseGithub)

	json.NewDecoder((<-c).Body).Decode(&u)

	bytes, _ := json.Marshal(u)

	saveToFile(fileName, bytes)

	readFromFile(fileName)

}

func getInput() int {
	var i int
	_, err := fmt.Scanf("%d", &i)
	errorHandler(err)
	return i
}

func getResponse(url string, c chan *http.Response) {
	resp, _ := http.Get(url)
	c <- resp
}

//noinspection GoReservedWordUsedAsName
func print(a ...interface{}) {
	fmt.Println(a)
}

func saveToFile(fileName string, data []byte) {
	ioutil.WriteFile(fileName, data, 0777)
}

func readFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	errorHandler(err)
	print(string(data))
}

func errorHandler(err error) {
	if err != nil {
		print(err)
		os.Exit(1)
	}
}
