package migrate

import (
	"go-gql/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(model.User{})
	if err != nil {
		panic(err)
	}
	log.Println("db migrate success")
}
