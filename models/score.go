package models

//ScoreRecord represents the score info being saved
type ScoreRecord struct {
	Score     int
	ValidJSON bool
	Id        string
}

//ScoreOutput represents the output of the score along with its id
type ScoreOutput struct {
	Score int    `json:"score"`
	Id    string `json:"id"`
}

// NOTE: add validation for the fields in the structs above for future
