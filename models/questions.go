package models

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	Question string `json:"question"`
	AnswerA string `json:"answerA"`
	AnswerB string `json:"answerB"`
	AnswerC string `json:"answerC"`
	AnswerD string `json:"answerD"`
	Correct string `json:"correct"`
}