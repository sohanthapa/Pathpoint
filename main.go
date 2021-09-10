package main

import (
	"Pathpoint/server"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Println("need to provide 2 args")
		return
	}
	numOfScores, _ := strconv.Atoi(args[1])
	file, err := os.Open(args[0])
	if err != nil {
		log.Println("error when opening file")
		return
	}
	server.ReadFromFile(file, numOfScores)

}
