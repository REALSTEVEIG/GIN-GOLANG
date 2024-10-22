package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name"`
	Date string `json:"date"`
	Cost int `json:"cost"`
	IsCompleted bool `json:"iscompleted"`
}