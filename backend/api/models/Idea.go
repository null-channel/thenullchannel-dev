package models

// Idea struct representation of a video idea
type Idea struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Votes       Vote   `json:"votes" gorm:"foreignKey:Count"`
}

type Vote struct {
	Count int64
}
