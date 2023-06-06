package main

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addproject)
	e.GET("/my-testimonials", mytestimonials)

	// quary string
	// e.GET(project-detail/id:)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

func addproject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

func mytestimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

// func projectDetail(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := map[string]interface{}{
// 		"id": id,
// 		"Tittle": "aa"
// 		"Content": "aa"
// 	}
// 	var tmpl, err = template.ParseFiles("views/project-detail.html")

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}

// 	return tmpl.Execute(c.Response(), nil)
// }
