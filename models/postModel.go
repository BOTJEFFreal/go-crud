package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
	Date  time.Time
}
