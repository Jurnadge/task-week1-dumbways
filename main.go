package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"task-web-dev-with-bootstrap/connection"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ID          int
	Title       string
	StartDate   string
	EndDate     string
	Deadline    string
	Description string
	NodeJS      bool
	NextJS      bool
	ReactJS     bool
	TypeScript  bool
	Image       string
}

var ProjectData = []Project{
	{
		Title:       "Title of the project",
		StartDate:   "12 November 2023",
		EndDate:     "13 January 2024",
		Deadline:    "Deadline: 3 Month",
		Description: "Desc of the project",
		NodeJS:      true,
		NextJS:      true,
		ReactJS:     true,
		TypeScript:  true,
	},
	{
		Title:       "Title of the project 2",
		StartDate:   "12 November 2023",
		EndDate:     "13 January 2024",
		Deadline:    "Deadline: 3 Month",
		Description: "Desc of the project 2",
		NodeJS:      true,
		NextJS:      true,
		ReactJS:     true,
		TypeScript:  true,
	},
}

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addProject)
	e.POST("/added-project", addedProject)
	e.POST("/project-delete/:id", deleteProject)
	e.GET("/my-testimonials", myTestimonials)
	e.GET("/project-detail/:id", projectDetail)

	// quary string
	// e.GET(project-detail/id:)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, start_date, end_date, deadline, description, nodejs, nextjs, reactjs, typescript FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.ID, &each.Title, &each.StartDate, &each.EndDate, &each.Deadline, &each.Description, &each.NodeJS, &each.NextJS, &each.ReactJS, &each.TypeScript)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		result = append(result, each)
	}
	projects := map[string]interface{}{
		"Projects": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	return tmpl.Execute(c.Response(), projects)
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

	//getting date input
	startDateInput := c.FormValue("inputStartDate")
	endDateInput := c.FormValue("inputEndDate")
	//
	getStartDt, _ := time.Parse("2006-01-02", startDateInput)
	startDate := getStartDt.Format("2 Jan 2006")
	//
	getEndDt, _ := time.Parse("2006-01-02", endDateInput)
	endDate := getEndDt.Format("2 Jan 2006")
	//
	var finalDiffTime string
	timeDiff := getEndDt.Sub(getStartDt)
	//
	if timeDiff.Hours()/24 < 30 {
		fdt := strconv.FormatFloat(timeDiff.Hours()/24, 'f', 0, 64)
		finalDiffTime = "Deadline: " + fdt + "Day"
	} else if timeDiff.Hours()/24/30 < 12 {
		fdt := strconv.FormatFloat(timeDiff.Hours()/24/30, 'f', 0, 64)
		finalDiffTime = "Deadline: " + fdt + "Month"
	} else {
		fdt := strconv.FormatFloat(timeDiff.Hours()/24/30/12, 'f', 0, 64)
		finalDiffTime = "Deadline: " + fdt + "Year"
	}

	//start techtype
	var nodeJs bool
	if c.FormValue("NodeJs") == "yes" {
		nodeJs = true
	}

	var nextJs bool
	if c.FormValue("NextJs") == "yes" {
		nextJs = true
	}

	var reactJs bool
	if c.FormValue("ReactJs") == "yes" {
		reactJs = true
	}

	var typeScript bool
	if c.FormValue("TypeScript") == "yes" {
		typeScript = true
	}
	//end techtype

	var newData = Project{
		Title:       title,
		StartDate:   startDate,
		EndDate:     endDate,
		Deadline:    finalDiffTime,
		Description: description,
		NodeJS:      nodeJs,
		NextJS:      nextJs,
		ReactJS:     reactJs,
		TypeScript:  typeScript,
	}

	ProjectData = append(ProjectData, newData)

	fmt.Println(ProjectData)

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

	// data := map[string]interface{}{
	// 	"id":      id,
	// 	"Tittle":  "Judul Project",
	// 	"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Inventore nam neque dicta sint, laudantium commodi iure sed! Ipsam corrupti at mollitia obcaecati maxime alias sequi laborum et! Maiores, nulla libero?",
	// }

	var ProjectDetail = Project{}
	for i, data := range ProjectData {
		if id == i {
			ProjectDetail = Project{
				Title:       data.Title,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Deadline:    data.Deadline,
				Description: data.Description,
				NodeJS:      data.NodeJS,
				NextJS:      data.NextJS,
				ReactJS:     data.ReactJS,
				TypeScript:  data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Index : ", id)

	ProjectData = append(ProjectData[:id], ProjectData[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
