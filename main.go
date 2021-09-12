package main

import (
	"Pathpoint/server"
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
	scoreRecords, _ := server.ReadFromFile(file)

	outputList, err := server.GetScoreRecordList(numOfScores, scoreRecords)
	if err != nil {
		log.Println(err)
		return
	}
	server.WriteJSON(outputList) // write to json format

}
