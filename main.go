package main

import (
	"Pathpoint/server"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// read args from command line
	args := os.Args[1:]
	if len(args) != 2 {
		log.Println(server.ErrMissingArgs)
		return
	}
	numOfScores, _ := strconv.Atoi(args[1])
	file, err := os.Open(args[0])
	if err != nil {
		log.Println(server.ErrFileOpen)
		return
	}
	scoreRecords, err := server.ReadFromFile(file)

	if err != nil {
		log.Println(err)
		return
	}

	outputList, err := server.GetScoreRecordList(numOfScores, scoreRecords)
	if err != nil {
		log.Println(err)
		return
	}
	finalOutput, err := server.WriteJSON(outputList) // write the final output to json format
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Printf("%s\n", finalOutput) // display the final result
	}

}

// NOTE: talk about assumptions, goodFile/badFile, all the corner cases, validation in models package for future
