package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Config struct{
	DBHost string `json:"DBHost"`
	DBPort int `json:"DBPort"`
	DBUsername string `json:"DBUsername"`
	DBPassword string `json:"DBPassword"`
	DBName string `json:"DBName"`
}

type Application struct{
	Database *gorm.DB
	Config Config
}

var App Application

func Load(configFile string) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &App.Config)
	if err != nil {
		panic(err)
	}

	App.Database, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			App.Config.DBUsername,
			App.Config.DBPassword,
			App.Config.DBHost,
			App.Config.DBPort,
			App.Config.DBName,
		),
	)
	App.Database.LogMode(true)
}
