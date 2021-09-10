package server

import (
	"Pathpoint/models"
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//ReadFromFile function reads each line from the provided file
func ReadFromFile(file *os.File, numOfScores int) {
	scoreRecordList := make([]models.ScoreRecord, 0)
	fscanner := bufio.NewScanner(file)
	//reading line by line
	for fscanner.Scan() {
		if len(fscanner.Text()) > 0 {
			st := strings.SplitN(fscanner.Text(), " ", 2)
			if len(st) > 0 {
				firstPart := st[0]  //gets the score
				secondPart := st[1] //gets the json part
				fp := strings.Trim(firstPart, ":")
				score, err := strconv.Atoi(fp)
				if err != nil {
					//print error message here
					return
				}
				sr := getScoreRecord(secondPart, score)
				scoreRecordList = append(scoreRecordList, sr)
			}
		}
	}

	sort.Slice(scoreRecordList, func(i, j int) bool {
		return scoreRecordList[j].Score < scoreRecordList[i].Score
	})

	outputList := getScoreRecordList(numOfScores, scoreRecordList)
	writeJSON(outputList)
}
