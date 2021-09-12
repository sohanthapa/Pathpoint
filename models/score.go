package models

//ScoreRecord represents the score info being saved
type ScoreRecord struct {
	ValidJSON bool   `json:"-"`
	ID        string `json:"id"`
	Score     int    `json:"score"`
}

type ScoreMap map[int]ScoreRecord

// NOTE: add validation for the fields in the structs above for future
