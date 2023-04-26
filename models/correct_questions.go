package models

import "gorm.io/gorm"

type CorrectQuestions struct {
	gorm.Model
	Question string `json:"question"`
	Wrong string `json:"wrong"`
	Correct string `json:"correct"`
}