package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Video struct {
	Id          uuid.UUID      `json:"id" ,gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	PublishTime time.Time      `json:"publish_time"`
	Thumbnails  pq.StringArray `gorm:"type:[]text"`
	// Add any other fields you require
}
