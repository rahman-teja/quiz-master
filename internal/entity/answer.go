package entity

type Answers struct {
	ID        string    `json:"id"`
	Questions Questions `json:"questions"`
	Answer    string    `json:"answer"`
	IsCorrect bool      `json:"isCorrect"`
	Point     int64     `json:"point"`
}
