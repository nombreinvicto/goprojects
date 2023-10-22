package main

import "fmt"

func main() {
	jsonFilePath := "testdata\\bookworms.json"
	fmt.Println(loadBookworms(jsonFilePath))

}
