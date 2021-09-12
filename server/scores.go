package server

import (
	"Pathpoint/models"
	"fmt"
)

type ScoreOutput struct {
	Score int    `json:"score"`
	ID    string `json:"id"`
}

// getScoreRecord sets the score record value from
// the unmarshalled value of score info.
// ScoreRecord value keeps track of the Score, if the JSON value is valid for the
// ScoreRecord and the Id of the ScoreRecord
func getScoreRecord(scoreInfo string) models.ScoreRecord {
	var sr models.ScoreRecord
	//check if it is valid json format
	if IsJSON(scoreInfo) {
		dataFile, err := UnMarshallJSON(scoreInfo)
		if err != nil {
			fmt.Printf("error unmarshalling %v", err)
			return models.ScoreRecord{}
		}
		sr = models.ScoreRecord{
			ValidJSON: true,
			ID:        dataFile.ID,
		}
	} else {
		sr = models.ScoreRecord{
			ValidJSON: false,
		}
	}
	return sr
}

//GetScoreRecordList sets all the ScoreOutput into an output list and returns that value.
func GetScoreRecordList(numOfScores int, scoreRecords []models.ScoreRecord) ([]ScoreOutput, error) {
	var j int

	outputList := make([]ScoreOutput, 0)
	//loop through the scoreRecordList and append the value to the outputList
	for _, sr := range scoreRecords {
		if j < numOfScores {
			if sr.ValidJSON {
				o := ScoreOutput{
					Score: sr.Score,
					ID:    sr.ID,
				}
				outputList = append(outputList, o)
				j++
			} else {
				return []ScoreOutput{}, ErrInvalidJSON
			}
		}
	}
	return outputList, nil
}
