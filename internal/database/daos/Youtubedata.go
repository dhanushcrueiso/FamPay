package daos

import (
	"FAMPAY/config"
	"FAMPAY/internal/database/db"
	"FAMPAY/internal/database/models"
	"FAMPAY/internal/dtos"
	"errors"
	"fmt"
	"log"
)

func UpsertVideoData(req []models.Video) error {
	log.Println("checking daos upsert level")
	err := db.DB.Unscoped().Table("videos").Save(&req).Error
	if err != nil {
		return errors.New("unable to place order")
	}
	return nil
}

func GetAllPaginated(req dtos.DataFilter) ([]models.Video, error) {
	res := []models.Video{}
	query := db.DB.Debug().Table("videos")

	if req.Q != "" {
		query.Where("id = (?) ", "60ddde02-d350-4bcb-99f5-c25abd7c6966")
	}

	offset := config.PageSize * (req.Page - 1)
	if offset > 0 {
		query.Offset(req.Page * config.PageSize)
	}
	query.Limit(config.PageSize)
	err := query.Find(&res).Error
	if err != nil {
		fmt.Println("error while fetching GetShipmentCounts")
		return nil, err
	}
	return res, nil
}
