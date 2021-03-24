package db

import (
	"api/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect established a databsase connection and performs migrations
func Connect() *gorm.DB {
	// use postgres for prod
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		logrus.Panicf("failed to connect to db. Reason : %s", err.Error())
	}
	db.AutoMigrate(&models.Idea{})
	db.AutoMigrate(&models.Vote{})
	return db
}

// Insert creates a new video idea
func Insert(idea models.Idea, dbcon *gorm.DB) {
	inserterr := dbcon.Create(&idea).Error
	if inserterr != nil {
		logrus.Fatalf("failed to insert %v to database. Reason %s", idea, inserterr.Error())
	}
}

func UpdateIdea(idea *models.Idea, dbcon *gorm.DB) {
	updatedidea := models.Idea{
		ID:          idea.ID,
		Description: idea.Description,
		Votes:       idea.Votes,
	}
	updaterr := dbcon.First(&models.Idea{}, "id = ?", idea.ID).Save(updatedidea).Error
	if updaterr != nil {
		logrus.Fatalf("unable to upate video idea, reason %s", updaterr.Error())
	}
}

func IncVote(idea *models.Idea, dbcon *gorm.DB) {
	idea.Votes.Count += 1
	incerr := dbcon.First(idea).Save(&idea).Error
	if incerr != nil {
		logrus.Fatal("could not increment vote. Reason %s", incerr.Error())
	}
}

func GetIdeas(dbcon *gorm.DB) []models.Idea {
	var ideas []models.Idea
	searcherr := dbcon.Find(&ideas).Error
	if searcherr != nil {
		logrus.Fatalf("could not retrieve ideas. Reason %s", searcherr.Error())
	}
	return ideas
}
