package models

type ScoreRecord struct {
	Score     int
	ValidJSON bool
	Id        string
}

type ScoreOutput struct {
	Score int    `json:"score"`
	Id    string `json:"id"`
}
