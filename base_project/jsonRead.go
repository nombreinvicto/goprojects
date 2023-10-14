package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}
type Users struct {
	Users []User `json:"users"`
}

func readJson(filename string) {

	// get a pointer to the file. its *File. File is a struct
	jsonFilePtr, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFilePtr.Close()

	// read whole file in memory as byte array
	byteSliceRead, err := io.ReadAll(jsonFilePtr)
	if err != nil {
		panic(err)
	}

	// init the Users slice
	var users Users

	// unmarshall = loads = json string to Go struct
	err = json.Unmarshal(byteSliceRead, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println(users.Users[1].Name)

}
