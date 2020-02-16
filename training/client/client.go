package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

var url = "http://localhost:8080"

func main() {
	//timeOutWithHTTPClient()
	//timeOutWithTransport()
	//timeOutWithContext()
	clientCancelReq()
}

func timeOutWithHTTPClient() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func timeOutWithTransport() {
	tr := &http.Transport{
		DialContext: (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
		//MaxIdleConns: 10,
		//IdleConnTimeout:    5 * time.Second,
		//DisableCompression: true,
	}

	client := http.Client{
		Transport: tr,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func timeOutWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error when init request", err)
	}
	resp, reqErr := http.DefaultClient.Do(req.WithContext(ctx))
	if reqErr != nil {
		fmt.Println("Error when perform request", reqErr)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func clientCancelReq() {
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	req = req.WithContext(ctx)

	go func() {
		// if google doesn't reponse after 10ms, I will cancel request
		time.Sleep(10 * time.Microsecond)
		cancel()
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
	fmt.Printf("%s", resp.Body)

}

/*
There are 2 ways to define the timeout in client side:
	- Inside the http client
	- Inside the request
	- In httpTransport (this is part of httpClient)

Cancellation in backend: https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/
*/
