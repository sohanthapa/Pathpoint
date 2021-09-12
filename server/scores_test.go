package server

import (
	"Pathpoint/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestGetScoreRecordList tests the functionality of getScoreRecordList
func TestGetScoreRecordList(t *testing.T) {
	t.Run("error - invalid json", func(t *testing.T) {
		numOfScores := 1
		scoreRecordList := []models.ScoreRecord{
			{
				Score:     12344,
				ValidJSON: false,
				ID:        "123445",
			},
		}
		outputList, err := GetScoreRecordList(numOfScores, scoreRecordList)
		assert.Equal(t, 0, len(outputList))
		assert.Equal(t, err, ErrInvalidJSON)
	})

	t.Run("success", func(t *testing.T) {
		numOfScores := 2
		scoreRecordList := []models.ScoreRecord{
			{
				Score:     12344,
				ValidJSON: true,
				ID:        "123445",
			},
			{
				Score:     12347,
				ValidJSON: true,
				ID:        "123445gttg",
			},
		}
		outputList, err := GetScoreRecordList(numOfScores, scoreRecordList)
		assert.Equal(t, 2, len(outputList))
		assert.Nil(t, err)
	})
}
