package entity

type Questions struct {
	ID        string   `json:"id"`
	Questions string   `json:"questions"`
	Answers   []string `json:"answers"`
	Point     int64    `json:"point"`
}
