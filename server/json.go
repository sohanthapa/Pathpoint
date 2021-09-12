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
	ID        string `json:"id"`
}

//IsJSON checks if the string is a valid json format
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

//UnMarshallJSON unmarshalls the string s into DataFileInput struct
func UnMarshallJSON(s string) (DataFileInput, error) {
	var dataFileInput DataFileInput
	// unmarschal JSON
	err := json.Unmarshal([]byte(s), &dataFileInput)

	//check for error when unmarshalling
	if err != nil {
		return DataFileInput{}, err
	}
	return dataFileInput, nil
}

// WriteJSON encode data as JSON and prints the data in JSON format
func WriteJSON(data interface{}) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
