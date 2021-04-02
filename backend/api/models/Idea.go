package models

// Idea struct representation of a video idea
type Idea struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Votes       uint   `json:"votes" gorm:"default:0"`
}
