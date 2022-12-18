package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

var sl []Note

func main() {
	e := echo.New()
	e.POST("/save_note", SaveNote)
	e.GET("/get_notes", GetNotes)
	log.Fatalln(e.Start("127.0.0.1:4000"))
}

func SaveNote(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}

	n := Note{}

	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}

	sl = append(sl, n)

	fmt.Println("Name:", n.name)
	fmt.Println("Surname", n.surname)
	fmt.Println("Info", n.note)
	return c.NoContent(200)
}

func GetNotes(c echo.Context) error {
	return c.JSON(http.StatusOK, &sl)
}
