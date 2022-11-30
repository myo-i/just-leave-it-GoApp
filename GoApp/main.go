package main

import (
	"GoApp/app/controllers"
	"GoApp/app/models"
	// "GoApp/bitflyer"
	"GoApp/config"
	"GoApp/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
}