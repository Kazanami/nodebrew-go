package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func main() {
	var url string = os.Args[1]
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", response.Status)
	
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, filename := path.Split(url)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer func(){
		file.Close()
	}()

	file.Write(body)
}
