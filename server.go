package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func GetAllFriends(screenName string) {
	anaconda.SetConsumerKey("")
	anaconda.SetConsumerSecret("")

	api := anaconda.NewTwitterApi("", "")

	friendIds := []anaconda.User{}
	v := url.Values{}
	v.Set("count", "100")
	v.Set("screen_name", screenName)
	nextCursor := "-1"
	for nextCursor != "0" {
		v.Set("cursor", nextCursor)
		c, err := api.GetFriendsList(v)
		if err != nil {
			log.Fatal(err)
		}
		friendIds = append(friendIds, c.Users...)
		fmt.Print(c.Users[0].ScreenName)
		nextCursor = c.Next_cursor_str
	}

}

func search(c echo.Context) error {
	screenName := c.FormValue("query")
	GetAllFriends(screenName)

	return c.HTML(http.StatusOK, fmt.Sprintf("%s", screenName))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.Static("/", "public")
	e.POST("/", search)

	e.Logger.Fatal(e.Start(":1323"))
}
