package server

import (
	"encoding/json"
	"fmt"
	"log"
)

type DataFileInput struct {
	Umbrella  int    `json:"umbrella"`
	Name      string `json:"name"`
	Value     int    `json:"value"`
	Payload   string `json:"payload"`
	DateStamp int    `json:"date_stamp"`
	Time      int    `json:"time"`
	Id        string `json:"id"`
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func UnMarshallJSON(s string) (DataFileInput, error) {
	var dataFile DataFileInput
	// unmarschal JSON
	err := json.Unmarshal([]byte(s), &dataFile)

	//check for error when unmarshalling
	if err != nil {
		return DataFileInput{}, err
	}
	return dataFile, nil
}

// writeJSON encode data as JSON and prints the data in JSON format
func writeJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
