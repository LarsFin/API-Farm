package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

/*
Each API will under go 100 performance test cycles against the same performance test script.
A test cycle will include the following requests, where n is the number of the cycle;

> Get all Video Games
> Create n Video Game(s)
> Update nth Video Game
> Get nth Video Game
> Delete nth Video Game

A wait of half a second will be given before re-applying the requests above in the next cycle.
*/

func main() {

	hostPtr := flag.String("host", "localhost", "hostname of the target api")
	portPtr := flag.String("port", "8080", "port of target api")

	flag.Parse()

	baseUrl := fmt.Sprintf("http://%s:%s", *hostPtr, *portPtr)

	if !isTargetReady(baseUrl) {
		fmt.Printf("The target api running at %s already has data.\nPlease clear before running performance tests.", baseUrl)
		return
	}

	client := &http.Client{}

	getAllReq := getAllRequest(baseUrl)
	postReq := postRequest(baseUrl)

	rt := makeRequest(client, getAllReq)

	for i := 0; i < 100; i++ {
		rt.aggregate(makeRequest(client, postReq))
	}

	fmt.Printf("Success Rate of Requests:\t%v\n", rt.SuccessRate)
	fmt.Printf("Time to First Byte:\t%v\n", rt.TimeToFirstByte)
	fmt.Printf("Full Response Time:\t%v\n", rt.FullResponseTime)
}

// Determines whether target api is clean and fair for performance testing
func isTargetReady(baseUrl string) bool {
	resp, err := http.Get(fmt.Sprintf("%s/video_games", baseUrl))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var data []interface{}

	err = json.Unmarshal(responseBody, &data)

	if err != nil {
		panic(err)
	}

	return len(data) == 0
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

// Sets video game id for a request by altering its path
func setRequestVgId(req *http.Request, id int) {
	req.URL.Path = fmt.Sprintf("/video_games/%d", id)
}

// Creates get request to retrieve a video game
func getRequest(baseUrl string) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/video_games/1", baseUrl), nil)

	if err != nil {
		panic(err)
	}

	return req
}

// Creates get request to retrieve all video games
func getAllRequest(baseUrl string) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/video_games", baseUrl), nil)

	if err != nil {
		panic(err)
	}

	return req
}

// Creates post request to create a new video game
func postRequest(baseUrl string) *http.Request {
	bs := readSample("post")
	r := bytes.NewReader(bs)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/video_games", baseUrl), r)

	if err != nil {
		panic(err)
	}

	return req
}

// Creates put request to update an existing video game
func putRequest(baseUrl string) *http.Request {
	bs := readSample("put")
	r := bytes.NewReader(bs)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/video_games/1", baseUrl), r)

	if err != nil {
		panic(err)
	}

	return req
}

// Creates delete request to delete an existing video game
func deleteRequest(baseUrl string) *http.Request {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/video_games/1", baseUrl), nil)

	if err != nil {
		panic(err)
	}

	return req
}

// Makes a request, records and responds response metrics
func makeRequest(client *http.Client, req *http.Request) *ResponseTrace {
	rt := &ResponseTrace{}

	start := time.Now()

	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			rt.TimeToFirstByte = time.Since(start)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		rt.SuccessRate = 1
	} else {
		rt.SuccessRate = 0
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	rt.FullResponseTime = time.Since(start)

	return rt
}

// Performance data regarding response
type ResponseTrace struct {
	// Success rate as float from 0 to 1
	SuccessRate float32

	// Time to first byte received from response
	TimeToFirstByte time.Duration

	// Full response time
	FullResponseTime time.Duration
}

// Averages data from passed response trace into called one
func (rt1 *ResponseTrace) aggregate(rt2 *ResponseTrace) {
	rt1.SuccessRate = (rt1.SuccessRate + rt2.SuccessRate) / 2
	rt1.TimeToFirstByte = (rt1.TimeToFirstByte + rt2.TimeToFirstByte) / 2
	rt1.FullResponseTime = (rt1.FullResponseTime + rt2.FullResponseTime) / 2
}
