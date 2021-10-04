package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"text/template"
)

var err error

func main() {
	dataFileNamePtr := flag.String("data", "data.json", "Name of the data json file")

	flag.Parse()

	templateFileName := flag.Args()[0]

	t1 := template.New(templateFileName)

	t1, err = t1.ParseFiles(templateFileName)
	if err != nil {
		log.Fatal(err)
	}

	dataFile, err := os.Open(*dataFileNamePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer dataFile.Close()

	decoder := json.NewDecoder(dataFile)

	var x interface{}

	decoder.Decode(&x)

	err = t1.Execute(os.Stdout, x)
	if err != nil {
		log.Fatal(err)
	}
}
