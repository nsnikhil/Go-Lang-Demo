package main

import (
	"io"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fileName := dir + "/interfaces/assignmentTwo/" + os.Args[1]
	f, _ := os.Open(fileName)
	io.Copy(os.Stdout,f)
	//ile, _ := ioutil.ReadFile(fileName)
	//fmt.Println(string(file))
}
