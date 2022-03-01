package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Names struct {
	Names []string `json:"names"`
}

var nameStruct Names
var names = []string{}

func RandomName() string {
	return names[rand.Intn(len(names))]
}

func Init() error {
	fileBytes, err := ioutil.ReadFile("names.json")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("an error occurred reading the 'names.json' file. are you sure it exists?")
	}
	fmt.Println("successfully opened names.json")
	json.Unmarshal(fileBytes, &nameStruct)
	names = nameStruct.Names
	rand.Seed(time.Now().UnixNano())
	return nil
}
