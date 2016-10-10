package main

import (
	//"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/serm-tape/todo/core"
	"github.com/serm-tape/todo/controller"
	"github.com/serm-tape/todo/model"

	"fmt"
	"flag"
	"net/http"
)

func main(){
	fmt.Println("hello world")

	configFile := flag.String("config", "config.json", "Path to configuration file")
	port := flag.String("port", "1323", "port to be listening")
	flag.Parse()

	core.Load(*configFile)
	core.App.Database.AutoMigrate(&model.Task{})

	e := echo.New()
	e.Get("/", func(c echo.Context) error {return c.String(http.StatusOK,"hello echo")})

	api := e.Group("/api")	//grouping for future middleware eg.authen
	api.Get("/task", controller.GetAllTask)
	api.Get("/task/:id", controller.GetTaskById)
	api.Post("/task", controller.CreateTask)
	api.Put("/task/:id", controller.EditTask)
	api.Delete("/task/:id", controller.DeleteTask)
	
	e.Run(standard.New(":"+*port))
}
