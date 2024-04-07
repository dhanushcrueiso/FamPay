package models

type Video struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PublishTime string   `json:"publish_time"`
	Thumbnails  []string `json:"thumbnails"`
	// Add any other fields you require
}
