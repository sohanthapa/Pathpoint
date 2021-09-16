package server

import (
	"Pathpoint/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		file, _ := os.Open("test/badFile.txt")
		scoreRecords, err := ReadFromFile(file)
		assert.Equal(t, 0, len(scoreRecords))
		assert.Equal(t, ErrConvertStringToInt, err)
	})
	t.Run("success", func(t *testing.T) {
		file, _ := os.Open("test/goodFile.txt")
		scoreRecords, err := ReadFromFile(file)
		expectedBody := []models.ScoreRecord{
			{
				ValidJSON: true,
				ID:        "085a11e1b82b441184f4a193a3c9a13c",
				Score:     13214012,
			},
			{
				ValidJSON: true,
				ID:        "84a0ccfec7d1475b8bfcae1945aea8f0",
				Score:     11446512,
			},
			{
				ValidJSON: true,
				ID:        "3c867674494e4a7aac9247a9d9a2179c",
				Score:     10622876,
			},
		}
		assert.Equal(t, expectedBody, scoreRecords)
		assert.Nil(t, err)
	})
}
