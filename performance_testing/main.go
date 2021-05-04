package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
Each API will under go 100 performance test cycles against the same performance test script.
A test cycle will include the following requests, where n is the number of the cycle;

> Create n Video Game(s)
> Update last Video Game
> Get last Video Game
> Get all Video Games
> Delete last Video Game

A wait of half a second will be given before re-applying the requests above in the next cycle.
*/

func main() {

	hostPtr := flag.String("host", "localhost", "hostname of the target api")
	portPtr := flag.String("port", "8080", "port of target api")

	flag.Parse()

	baseUrl := fmt.Sprintf("http://%s:%s", *hostPtr, *portPtr)

	postBytes := readSample("post")

	for i := 0; i < 3; i++ {
		r := bytes.NewReader(postBytes)
		postRequest(baseUrl, r)
	}

	// for i := 0; i < 2; i++ {
	// 	var buf bytes.Buffer
	// 	tee := io.TeeReader(postReader, &buf)

	// 	postRequest(baseUrl, tee)
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter name: ")
	// name, _ := reader.ReadString('\n')
	// fmt.Printf("Hello %s", name)

	// r := read("sample.json")

	// resp, err := http.Post("http://localhost:8080/video_games", "application/json", r)

	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// bs, _ := ioutil.ReadAll(resp.Body)

	// fmt.Printf("Status: %s\nBody: %s", resp.Status, bs)
}

// Reads sample json for post and put requests
func readSample(sampleName string) []byte {
	path := fmt.Sprintf("sample/%s.json", sampleName)

	b, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return b
}

// Makes a post request to the api to create a video game
func postRequest(baseUrl string, r io.Reader) {
	url := fmt.Sprintf("%s/video_games", baseUrl)
	resp, err := http.Post(url, "application/json", r)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(bs))
}
