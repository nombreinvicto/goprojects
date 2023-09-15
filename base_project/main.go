package main

import (
	"fmt"
	"os"
)

func main() {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

}
