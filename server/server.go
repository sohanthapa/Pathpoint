package server

import (
	"Pathpoint/models"
	"Pathpoint/utils/helpers"
	"bufio"
	"os"
	"strconv"
	"strings"
)

//ReadFromFile function reads each line from the provided file and appends the necessary
// values to the scoreRecord list
func ReadFromFile(file *os.File) ([]models.ScoreRecord, error) {
	fscanner := bufio.NewScanner(file)
	scoreMap := make(models.ScoreMap)
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
					return []models.ScoreRecord{}, ErrConvertStringToInt
				}
				sr := getScoreRecord(secondPart)
				scoreMap[score] = sr
			}
		}
	}
	scoreRecords := helpers.SortInReverse(scoreMap)

	return scoreRecords, nil
}
