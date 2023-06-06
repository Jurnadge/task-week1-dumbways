package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addProject)
	e.POST("/added-project", addedProject)
	e.GET("/my-testimonials", myTestimonials)
	e.GET("/project-detail/:id", projectDetail)

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

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

func addedProject(c echo.Context) error {
	title := c.FormValue("inputProjectName")
	description := c.FormValue("inputDescription")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")

	println("Tittle : " + title)
	println("Description : " + description)
	println("Start date : " + startDate)
	println("End date :" + endDate)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func myTestimonials(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/my-testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"Tittle":  "Judul Project",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Inventore nam neque dicta sint, laudantium commodi iure sed! Ipsam corrupti at mollitia obcaecati maxime alias sequi laborum et! Maiores, nulla libero?",
	}
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}
