package controller

import (
	"github.com/gcharalla/url-shortener/database"
	"github.com/gcharalla/url-shortener/models"
)

func GetAllGolies() ([]models.Goly, error) {
	var golies []models.Goly

	tx := database.DB.Db.Find(&golies)
	if tx.Error != nil {
		return []models.Goly{}, tx.Error
	}

	return golies, nil
}

func GetGoly(id uint64) (models.Goly, error) {
	var goly models.Goly

	tx := database.DB.Db.Where("id = ?", id).First(&goly)

	if tx.Error != nil {
		return models.Goly{}, tx.Error
	}

	return goly, nil
}

func CreateGoly(goly models.Goly) error {
	tx := database.DB.Db.Create(&goly)
	return tx.Error
}

func UpdateGoly(goly models.Goly) error {

	tx := database.DB.Db.Save(&goly)
	return tx.Error
}

func DeleteGoly(id uint64) error {

	tx := database.DB.Db.Unscoped().Delete(&models.Goly{}, id)
	return tx.Error
}

func FindByGolyUrl(url string) (models.Goly, error) {
	var goly models.Goly
	tx := database.DB.Db.Where("goly = ?", url).First(&goly)
	return goly, tx.Error
}
