package server

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ScoreRecord struct {
	Score     int
	ValidJSON bool
	Id        string
}

type ScoreOutput struct {
	Score int    `json:"score"`
	Id    string `json:"id"`
}

//ReadFromFile function reads each line from the provided file
func ReadFromFile(file *os.File, numOfScores int) {
	scoreRecordList := make([]ScoreRecord, 0)
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
				sr := setScoreRecord(secondPart, score)
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

func setScoreRecord(scoreInfo string, score int) ScoreRecord {
	var sr ScoreRecord
	if IsJSON(scoreInfo) {
		dataFile, err := UnMarshallJSON(scoreInfo)
		if err != nil {
			fmt.Printf("error unmarshalling %v", err)
			return ScoreRecord{}
		}
		sr = ScoreRecord{
			Score:     score,
			ValidJSON: true,
			Id:        dataFile.Id,
		}
	} else {
		sr = ScoreRecord{
			Score:     score,
			ValidJSON: false,
		}
	}
	return sr
}

func getScoreRecordList(numOfScores int, scoreRecordList []ScoreRecord) []ScoreOutput {
	var j int
	outputList := make([]ScoreOutput, 0)
	for _, y := range scoreRecordList {
		if j < numOfScores {
			if y.ValidJSON {
				o := ScoreOutput{
					Score: y.Score,
					Id:    y.Id,
				}
				outputList = append(outputList, o)
				fmt.Println(y.Score)
				j++
			} else {
				fmt.Println("invalid json for values")
				return []ScoreOutput{}
			}
		}
	}
	return outputList
}
