package db

import (
	"api/models"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a databsase connection and performs migrations
func Connect() *gorm.DB {
	// use postgres for prod
	dsn := os.Getenv("DB_URI")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Panicf("failed to connect to db. Reason : %s", err.Error())
	}
	logrus.Warn("SUCESSFULLY CONNECTED TO DB")
	db.AutoMigrate(&models.Idea{})
	return db
}

// Insert creates a new video idea
func Insert(idea models.Idea, dbcon *gorm.DB) {
	inserterr := dbcon.Create(&idea).Error
	if inserterr != nil {
		logrus.Fatalf("failed to insert %v to database. Reason %s", idea, inserterr.Error())
	}
}

func UpdateIdea(idea *models.Idea, dbcon *gorm.DB) error {
	updatedidea := models.Idea{
		ID:          idea.ID,
		Description: idea.Description,
		Votes:       idea.Votes,
	}
	updaterr := dbcon.First(&models.Idea{}, "id = ?", idea.ID).Save(updatedidea).Error
	if updaterr != nil {
		logrus.Fatalf("unable to upate video idea, reason %s", updaterr.Error())
		return updaterr
	}
	return nil
}

func IncVote(ideaID uint, dbcon *gorm.DB) error {
	idea := &models.Idea{}
	searcherr := dbcon.First(&idea, ideaID).Error
	if searcherr != nil {
		if searcherr == gorm.ErrRecordNotFound {
			return searcherr
		}
		logrus.Fatal("could not retrive record. Reason %s", searcherr.Error())
	}
	idea.Votes += 1
	updaterr := dbcon.Save(&idea).Error
	if updaterr != nil {
		logrus.Fatal("could not incement vote count reason %s", updaterr.Error())
	}
	return nil
}

func GetIdeas(dbcon *gorm.DB) []models.Idea {
	var ideas []models.Idea
	searcherr := dbcon.Find(&ideas).Error
	if searcherr != nil {
		logrus.Fatalf("could not retrieve ideas. Reason %s", searcherr.Error())
	}
	return ideas
}
