package main

import (
	"crudproduct/attribute"
	"crudproduct/config"
	"crudproduct/model"
	"crudproduct/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func main() {
	var arguments, messageLog string

	if len(os.Args) > 1 {
		arguments = os.Args[1]
	} else {
		arguments = "development"
	}

	config.GenerateConfiguration(arguments)
	attribute.SetServerAttribute()

	//sql migration gorm
	err := migrateDatabase(attribute.ServerAttribute.GormClient)
	if err != nil {
		messageLog = "Failed to Migrate Database : " + err.Error()
		fmt.Println(messageLog)
		os.Exit(3)
	}

	messageLog = "App Start in port : " + strconv.Itoa(config.ApplicationConfiguration.GetServerPort())
	fmt.Println(messageLog)

	routes := gin.Default()
	router.APIController(routes)
}

func migrateDatabase(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return
	}
	return
}
