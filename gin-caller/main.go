package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Animal struct {
	ID   string `json: "id"`
	Name string `json: "name"`
	Type string `json: "type"`
}

func callLocalHost() {
	n := 100

	for n > 0 {

		response, err := http.Get("http://localhost:8080/albums")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
		n--
	}
}

func structureAnimalData(records [][]string) {
	var animal Animal
	var animals []Animal

	for _, record := range records {
		animal.ID = record[0]
		animal.Name = record[1]
		animal.Type = record[2]
		animals = append(animals, animal)
	}
	fmt.Println(animals)
	animalJson, err := json.Marshal(animals)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(animalJson))

}

func parseAndPrintCSV(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = '|'
	record, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Mess up in file")
	}
	structureAnimalData(record)
}

func main() {
	// callLocalHost()
	parseAndPrintCSV("./example.csv")
}
