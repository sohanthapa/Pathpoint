package server

import (
	"Pathpoint/models"
	"fmt"
)

// getScoreRecord sets the score record value from
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

//getScoreRecordList sets all the ScoreOutput into an output list and returns that value.
func getScoreRecordList(numOfScores int, scoreRecordList []models.ScoreRecord) []models.ScoreOutput {
	var j int
	outputList := make([]models.ScoreOutput, 0)
	for _, y := range scoreRecordList {
		if j < numOfScores {
			if y.ValidJSON {
				o := models.ScoreOutput{
					Score: y.Score,
					Id:    y.Id,
				}
				outputList = append(outputList, o)
				fmt.Println(y.Score)
				j++
			} else {
				fmt.Println("invalid json for values")
				return []models.ScoreOutput{}
			}
		}
	}
	return outputList
}
