package utils

import (
	"dilya/collab/data"
	"encoding/json"
	"io/ioutil"
)

func WriteUser() {
	marshalledData, _ := json.Marshal(data.UsersSlice)
	ioutil.WriteFile("db.json", marshalledData, 0644)
}

func ReadUser() {
	readByte, _ := ioutil.ReadFile("db.json")
	json.Unmarshal(readByte, &data.UsersSlice)
}