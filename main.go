package main

import (
	"log"
	"os"

	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/handler"
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql"
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		Skipper: func(c echo.Context) bool {
			return c.Request().Method == echo.OPTIONS
		},
	}))
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	rw, err := mysql.InitMysql()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ro, err := mysql.InitMysql()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	roomRepo := repository.NewRoom(rw, ro)
	roomHandler := handler.NewRoomHandler(roomRepo)
	e.GET("/room/:id", roomHandler.GetRoom)

	e.Logger.Fatal(e.Start(":8080"))
}
