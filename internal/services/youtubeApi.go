package services

import (
	"FAMPAY/config"
	"FAMPAY/internal/database/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	youtubeService *youtube.Service
)
var lastPublishTime time.Time

func init() {
	// Initialize YouTube API client
	httpClient := &http.Client{
		Transport: &transport.APIKey{Key: config.APIKey},
	}
	var err error
	youtubeService, err = youtube.New(httpClient)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

}

func FetchAndStoreVideos() {
	query := "Football"
	call := youtubeService.Search.List([]string{"snippet"}).
		Q(query).
		MaxResults(5).
		Order("date").
		Type("video").
		PublishedAfter(lastPublishTime.Format(time.RFC3339))
	response, err := call.Do()
	if err != nil {
		log.Printf("Error fetching videos from YouTube: %v", err)
		return
	}

	for _, item := range response.Items {
		video := models.Video{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishTime: item.Snippet.PublishedAt,
			Thumbnails:  make([]string, 0),
		}
		if item.Snippet.Thumbnails != nil {
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.High.Url)
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.Medium.Url)
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.Default.Url)
		}

		fmt.Println("result ", video)
	}
	if len(response.Items) > 0 {
		lastPublishTime = parseTime(response.Items[0].Snippet.PublishedAt)
	}

}

func parseTime(publishedAt string) time.Time {
	t, err := time.Parse(time.RFC3339, publishedAt)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
	}
	return t
}
