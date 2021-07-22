package main

import (
	"encoding/json"
	"log"

)

const data = `[
	{"name": "smogy", "age": 4, "location": "heart"},
	{"name": "nata", "age": 4, "location": "heart"}
]`

func GetData() (people []Person) {
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		log.Fatal(err)
	}
	return people
}
