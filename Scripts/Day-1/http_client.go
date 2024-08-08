package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiUrl := "http://localhost"

	// For get call
	res, err := http.Get(apiUrl + "/get")
	if err != nil {
		fmt.Println("Error getting respose from the api server: ", err)
		return
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	fmt.Println("Response from the API server:", string(resBody))
}
