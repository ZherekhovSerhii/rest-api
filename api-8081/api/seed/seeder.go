package seed

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/ZherekhovSerhii/http-rest-api/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = [1000]models.Post{}

/*	models.Post{
		Title: "Ticket 1",
		From:  "Poland",
		To:    "Ukraine",
		Price: 100,
	},
	models.Post{
		Title: "Ticket 2",
		From:  "Turkey",
		To:    "France",
		Price: 200,
	},
	models.Post{
		Title: "Ticket 3",
		From:  "Turkey",
		To:    "France",
		Price: 300,
	},*/

// Load ...
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range posts {
		posts[i].Title = "Ticket " + strconv.Itoa(i+1)
		posts[i].From = "Turkey"
		posts[i].To = "Ukraine"
		posts[i].Price = rand.Uint32() //uint32((i + 1))
		posts[i].AuthorID = 1          //users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
