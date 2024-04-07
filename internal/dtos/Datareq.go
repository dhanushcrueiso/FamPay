package dtos

import (
	"time"

	"github.com/google/uuid"
)

type DataFilter struct {
	Page int    `json:"pg"`
	Q    string `json"q"`
}
type Video struct {
	Id          uuid.UUID `json:"id" `
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishTime time.Time `json:"publish_time"`
	Thumbnails  []string  `json:"thumbnails"`
}
