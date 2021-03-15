package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Record []struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		data := readfile(string(arg))
		fmt.Println(data)
	} else {
		fmt.Println("Please pass the argument...")
		os.Exit(1)
	}
}

func readfile(name string) string {
	details := ""
	r := Record{}
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	json.Unmarshal(data, &r)
	for _, record := range r {
		lowerfname, lowerWordfname := strings.ToLower(record.Firstname), strings.ToLower(name)
		lowerlname, lowerWordlname := strings.ToLower(record.Lastname), strings.ToLower(name)
		if strings.Contains(lowerfname, lowerWordfname) || strings.Contains(lowerlname, lowerWordlname) {
			details = details + "Name:" + record.Firstname + " " + record.Lastname + ", " + "Address:" + record.Address + "\n"
		}
	}
	return details
}
