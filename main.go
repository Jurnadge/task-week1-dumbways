package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"task-web-dev-with-bootstrap/connection"
	"task-web-dev-with-bootstrap/middleware"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Project struct {
	Id         int
	Title      string
	Content    string
	StartDate  string
	EndDate    string
	Duration   string
	NodeJS     bool
	NextJS     bool
	ReactJS    bool
	TypeScript bool
	Image      string
	Author     string
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// var dataProject = []Project{
// 	{
// 		Title:      "Dumbways Web Apps 2023",
// 		Content:    "Content",
// 		StartDate:  "2023-05-08",
// 		EndDate:    "2023-06-08",
// 		Duration:   "1 month",
// 		NodeJS:     true,
// 		NextJS:     true,
// 		ReactJS:    true,
// 		TypeScript: true,
// 	},
// 	{
// 		Title:      "Dumbways Web Apps 2023",
// 		Content:    "Content 2",
// 		StartDate:  "2023-05-08",
// 		EndDate:    "2023-06-08",
// 		Duration:   "1 month",
// 		NodeJS:     true,
// 		NextJS:     true,
// 		ReactJS:    true,
// 		TypeScript: true,
// 	},
// }

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// Serve a static files form "public" directory
	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	// To use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/my-testimonials", testimonials)
	e.GET("/add-project", addProjectForm)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/update-project/:id", updateProjectForm)
	e.POST("/add-project", middleware.UploadFile(addProject))
	e.POST("/project-delete/:id", deleteProject)
	e.POST("/update-project/:id", middleware.UploadFile(updatedProject))

	//SIGNUP
	e.GET("/signup-form", signupForm)
	e.POST("/signup", singup)

	//LOGIN
	e.GET("/login-form", loginForm)
	e.POST("/login", login)

	//LOGOuT
	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	//selecting table from postgres
	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_project.id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image, tb_user.name AS author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id ORDER BY tb_project.id DESC")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Title, &each.StartDate, &each.EndDate, &each.Duration, &each.Content, &each.NodeJS, &each.NextJS, &each.ReactJS, &each.TypeScript, &each.Image, &each.Author)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"Projects":    result,
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, errTemplate = template.ParseFiles("views/index.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), projects)

}

func contact(c echo.Context) error {

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func testimonials(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/my-testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesasge": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func addProjectForm(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	projects := map[string]interface{}{
		"DataSession": userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func calculateDuration(startDateInput string, endDateInput string) string {
	startTime, _ := time.Parse("2006-01-02", startDateInput)
	endTime, _ := time.Parse("2006-01-02", endDateInput)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func addProject(c echo.Context) error {
	title := c.FormValue("inputProjectName")
	content := c.FormValue("inputDescription")
	startDateInput := c.FormValue("inputStartDate")
	endDateInput := c.FormValue("inputEndDate")
	duration := calculateDuration(startDateInput, endDateInput)
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
	image := c.Get("dataFile").(string)

	sess, _ := session.Get("session", c)
	author := sess.Values["id"].(int)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", title, startDateInput, endDateInput, duration, content, nodeJs, nextJs, reactJs, typeScript, image, author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_project.id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image, tb_user.name AS author FROM tb_project JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_project.id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.NodeJS, &ProjectDetail.NextJS, &ProjectDetail.ReactJS, &ProjectDetail.TypeScript, &ProjectDetail.Image, &ProjectDetail.Author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Project":     ProjectDetail,
		"DataSession": userData,
	}

	var tmpl, errTemplate = template.ParseFiles("views/project-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Index : ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func updateProjectForm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, start_date, end_date, duration, content, nodejs, nextjs, reactjs, typescript, image FROM tb_project WHERE id=$1", id).Scan(
		&ProjectDetail.Id, &ProjectDetail.Title, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Content, &ProjectDetail.NodeJS, &ProjectDetail.NextJS, &ProjectDetail.ReactJS, &ProjectDetail.TypeScript, &ProjectDetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/update-project.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updatedProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("inputProjectName")
	content := c.FormValue("inputDescription")
	startDateInput := c.FormValue("inputStartDate")
	endDateInput := c.FormValue("inputEndDate")
	duration := calculateDuration(startDateInput, endDateInput)
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

	image := c.Get("dataFile").(string)

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET title=$1, start_date=$2, end_date=$3, duration=$4, content=$5, nodejs=$6, nextjs=$7, reactjs=$8, typescript=$9, image=$10 WHERE id=$11", title, startDateInput, endDateInput, duration, content, nodeJs, nextJs, reactJs, typeScript, image, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func signupForm(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/signup-form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func singup(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Signup failed, please try again!", false, "/signup-form")
	}

	return redirectWithMessage(c, "Signup success!", true, "/login-form")
}

func loginForm(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/login-form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/login-form")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect", false, "/login-form")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}
