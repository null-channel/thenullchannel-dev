package db

import (
	"api/models"

	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Connect establishes a databsase connection and performs migrations
func Connect()  {
	// use postgres for prod

	dsn := os.Getenv("DB_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("failed to connect to db. Reason : %s", err.Error())
		os.Exit(1)
	}
	logrus.Warn("SUCESSFULLY CONNECTED TO DB")

	err = db.AutoMigrate(&models.Idea{})
	if err != nil {
		return
	}
	db = DB
}

// Insert creates a new video idea

func Insert(idea models.Idea) (models.Idea, error) {
	inserterr := DB.Create(&idea).Error
	if inserterr != nil {
		logrus.Fatalf("failed to insert %v to database. Reason %s", idea, inserterr.Error())
		return models.Idea{}, inserterr
	}
	return idea, nil

}

func UpdateIdea(idea *models.Idea,) error {
	updatedidea := models.Idea{
		ID:          idea.ID,
		Description: idea.Description,
		Votes:       idea.Votes,
	}
	updaterr := DB.First(&models.Idea{}, "id = ?", idea.ID).Save(updatedidea).Error
	if updaterr != nil {
		logrus.Fatalf("unable to upate video idea, reason %s", updaterr.Error())
		return updaterr
	}
	return nil
}

func IncVote(ideaID uint,) error {
	idea := &models.Idea{}
	searcherr := DB.First(&idea, ideaID).Error
	if searcherr != nil {
		if searcherr == gorm.ErrRecordNotFound {
			return searcherr
		}
		logrus.Fatal("could not retrive record. Reason %s", searcherr.Error())
	}
	idea.Votes += 1
	updaterr := DB.Save(&idea).Error
	if updaterr != nil {
		logrus.Fatal("could not incement vote count reason %s", updaterr.Error())
	}
	return nil
}

func GetIdeas() []models.Idea {
	var ideas []models.Idea
	searcherr := DB.Find(&ideas).Error
	if searcherr != nil {
		logrus.Fatalf("could not retrieve ideas. Reason %s", searcherr.Error())
	}
	return ideas
}

func DeleteIdea(idea models.Idea) error {
	err := DB.Delete(&idea).Error
	if err != nil {
		return err
	}
	return nil
}
