package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Card struct {
	Number string
}

func (c *Card) Get() interface{} {
	url := fmt.Sprintf("%s/%s", BaseUrl, c.Number)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	var bodyContent []byte
	bodyContent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	return string(bodyContent)
}
