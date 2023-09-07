package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	amountRequest := 20000
	wg.Add(amountRequest)
	channel := make(chan error)
	amountErrors := 0
	for i:=0; i<amountRequest;i++{
		go func(c chan<- error) {
			defer wg.Done()
			err := sendRequest()
			if err != nil {
				c <- err
			}
		}(channel)
	}
	wg.Wait()
	for err := range channel {
		amountErrors++
		log.Printf("Error received: %v", err)
	}
	close(channel)
	log.Printf("Amount of errors recived: %v",amountErrors)
}

func sendRequest() error {
	url := "http://localhost:8080"
	endpoint := "/healthcheck"
	req, err := http.NewRequest("GET", url+endpoint,nil)
	if err != nil {
		fmt.Println("Error on request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	start := time.Now()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on send request:", err)
		return err
	}
	defer resp.Body.Close()
	duration := time.Since(start).Milliseconds()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error on read response:", err)
		return err
	}
	log.Printf("Response: %s (Time: %d ms)\n", body, duration)
	return nil
}