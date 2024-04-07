package services

import (
	"FAMPAY/config"
	"FAMPAY/internal/database/daos"
	"FAMPAY/internal/database/db"
	"FAMPAY/internal/database/models"
	"FAMPAY/internal/dtos"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
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
	env := "dev"

	var file *os.File
	var err error

	file, err = os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}
	//parsing json with the config and passing the dev.json values from here
	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)

	db.Init(&db.Config{
		URL:       cnf.DatabaseURL,
		MaxDBConn: cnf.MaxDBConn,
	})
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
	Videos := []models.Video{}
	for _, item := range response.Items {
		video := models.Video{
			Id:          uuid.New(),
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishTime: parseTime(item.Snippet.PublishedAt),
			Thumbnails:  make([]string, 0),
		}
		if item.Snippet.Thumbnails != nil {
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.High.Url)
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.Medium.Url)
			video.Thumbnails = append(video.Thumbnails, item.Snippet.Thumbnails.Default.Url)
		}
		Videos = append(Videos, video)
		fmt.Println("result ", video)
	}
	err = daos.UpsertVideoData(Videos)
	if err != nil {
		log.Fatalln("unable to upsert data to db")
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

func GetDataPaginated(req dtos.DataFilter) ([]dtos.Video, error) {

	res, err := daos.GetAllPaginated(req)
	if err != nil {
		return nil, err
	}

	resDtos := DaoToDtos(res)
	return resDtos, err

}

func DaoToDtos(req []models.Video) []dtos.Video {
	Videos := []dtos.Video{}
	for _, val := range req {
		video := dtos.Video{
			Id:          val.Id,
			Title:       val.Title,
			Description: val.Description,
			PublishTime: val.PublishTime,
			Thumbnails:  val.Thumbnails,
		}
		Videos = append(Videos, video)
	}
	return Videos
}
