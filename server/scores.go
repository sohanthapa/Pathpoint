package server

import (
	"Pathpoint/models"
	"fmt"
)

// GetScoreRecord sets the score record value from
// the unmarshalled value of score info.
// ScoreRecord value keeps track of the Score, if the JSON value is valid for the
// ScoreRecord and the Id of the ScoreRecord
func getScoreRecord(scoreInfo string, score int) models.ScoreRecord {
	var sr models.ScoreRecord
	//check if it is valid json format
	if IsJSON(scoreInfo) {
		dataFile, err := UnMarshallJSON(scoreInfo)
		if err != nil {
			fmt.Printf("error unmarshalling %v", err)
			return models.ScoreRecord{}
		}
		sr = models.ScoreRecord{
			Score:     score,
			ValidJSON: true,
			Id:        dataFile.Id,
		}
	} else {
		sr = models.ScoreRecord{
			Score:     score,
			ValidJSON: false,
		}
	}
	return sr
}

//GetScoreRecordList sets all the ScoreOutput into an output list and returns that value.
func GetScoreRecordList(numOfScores int, scoreRecordList []models.ScoreRecord) ([]models.ScoreOutput, error) {
	var j int
	outputList := make([]models.ScoreOutput, 0)
	//loop through the scoreRecordList and append the value to the outputList
	for _, sr := range scoreRecordList {
		if j < numOfScores {
			if sr.ValidJSON {
				o := models.ScoreOutput{
					Score: sr.Score,
					Id:    sr.Id,
				}
				outputList = append(outputList, o)
				j++
			} else {
				return []models.ScoreOutput{}, ErrInvalidJSON
			}
		}
	}
	return outputList, nil
}
