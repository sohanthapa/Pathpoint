package helpers

import (
	"Pathpoint/models"
	"sort"
)

//SortInReverse sorts the given ScoreMap in descending order
func SortInReverse(scoreMap models.ScoreMap) []models.ScoreRecord {
	scoreRecords := make([]models.ScoreRecord, 0)
	keys := make([]int, 0, len(scoreMap))
	for k := range scoreMap {
		keys = append(keys, k)
	}

	//sort in descending order
	sort.Slice(keys, func(i, j int) bool {
		return keys[j] < keys[i]
	})

	for _, k := range keys {
		sr := models.ScoreRecord{
			Score:     k,
			ValidJSON: scoreMap[k].ValidJSON,
			ID:        scoreMap[k].ID,
		}
		scoreRecords = append(scoreRecords, sr)
	}

	return scoreRecords
}
