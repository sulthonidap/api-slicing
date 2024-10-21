package db

import (
	"gorm.io/gorm"
)

type Posting struct {
	gorm.Model
	Mind    string `gorm:"type:text" json:"mind"`
	Picture string `gorm:"type:varchar(100)" json:"picture"`
}

type Comment struct {
	gorm.Model
	Sentence  string `gorm:"type:text" json:"sentence"`
	Fullname  string `gorm:"NOT NULL" json:"fullname"`
	Job       string `gorm:"NOT NULL" json:"job"`
	PostingID int    `gorm:"type:int(10) unsigned;index" json:"postingID" form:"postingId"`

	Posting Posting
}

type Like struct {
	gorm.Model
	PostingID int `gorm:"type:int(10) unsigned;index" json:"postingID" form:"postingId"`

	Posting Posting
}
