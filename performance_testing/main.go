package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter name: ")
	// name, _ := reader.ReadString('\n')
	// fmt.Printf("Hello %s", name)

	fmt.Println(readSample())
}

func readSample() string {
	b, _ := ioutil.ReadFile("sample.json")
	return string(b)
}
