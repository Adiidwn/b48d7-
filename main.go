package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	conect "vodlab/com/connection"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Author       string
	ProjectName  string
	Durations    string
	StartDate    time.Time
	EndDate      time.Time
	Description  string
	Technologies []string
	ReactJs      bool
	Golang       bool
	NodeJs       bool
	Javascipt    bool
	Image        string
}

var dataProject = []Project{
	// {
	// 	Author:       "Adiwidiawan",
	// 	ProjectName:  "Project 1",
	// 	StartDate:    "20July2023",
	// 	EndDate:      "21July2024",
	// 	Description:  "ini Project 1",
	// 	Technologies: "Javascipt , Node Js",
	// },
}

func main() {
	e := echo.New()
	conect.ConectDatabase()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/addProject", addProject)
	e.GET("/testimonials", testimonials)
	e.GET("/about", about)
	e.GET("/projectDetail/:id", projectDetail)
	e.POST("/addmyProject", addmyProject)
	e.POST("/deletemyProject/:id", deletemyProject)

	e.GET("/updateProject/:id", editProject)
	e.POST("/updatedProject/:id", updatedProject)

	e.Logger.Fatal(e.Start("localhost:666"))
}

// handler
func home(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/index.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	data1, dataerror := conect.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	if dataerror != nil {
		return x.JSON(500, err.Error())
	}

	dataProject = []Project{}
	for data1.Next() {
		var each = Project{}

		data1.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		// if dataerror != nil {
		// 	return x.JSON(500, err.Error())
		// }
		each.Author = "Adiwidiawan"
		fmt.Println("id:", each.Id, "namaP;", each.ProjectName)
		each.Durations = Duration(each.StartDate, each.EndDate)

		if checkValue(each.Technologies, "ReactJs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "Golang") {
			each.Golang = true
		}
		if checkValue(each.Technologies, "NodeJs") {
			each.NodeJs = true
		}
		if checkValue(each.Technologies, "Javascript") {
			each.Javascipt = true
		}
		// t1 := each.StartDate
		// t2 := each.EndDate
		// diff:=t1.Sub(t2)

		dataProject = append(dataProject, each)
	}

	data := map[string]interface{}{
		"Project": dataProject,
	}

	return tmplate.Execute(x.Response(), data)
}

func Duration(StartDate time.Time, EndDate time.Time) string {

	diff := EndDate.Sub(StartDate)
	day := int(diff.Hours() / 24)
	week := day / 7
	month := day / 30
	year := month / 12
	if day < 7 {
		return strconv.Itoa(day) + " Day"
	}
	if week < 4 {
		return strconv.Itoa(week) + " Week"
	}
	if month < 12 {
		return strconv.Itoa(month) + " Month"
	}
	return strconv.Itoa(year) + " Year"

}

func contact(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/contact.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}

func addProject(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/project1.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}
func about(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/about.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}

func testimonials(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/testimonials.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id := c.Param("id") // misal : 1

	tmpl, err := template.ParseFiles("htmls/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Id, _ := strconv.Atoi(id)

	dataProject := Project{}

	err1 := conect.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", Id).Scan(&dataProject.Id, &dataProject.ProjectName, &dataProject.StartDate, &dataProject.EndDate, &dataProject.Description, &dataProject.Technologies, &dataProject.Image)

	dataProject.Durations = Duration(dataProject.StartDate, dataProject.EndDate)

	if err1 != nil {
		return c.JSON(500, err.Error())
	}

	data := map[string]interface{}{
		"Project":    dataProject,
		"Id":         id,
		"startDateS": dataProject.StartDate.Format("2006-01-02"),
		"endDateS":   dataProject.EndDate.Format("2006-01-02"),
	}

	return tmpl.Execute(c.Response(), data)
}

func addmyProject(c echo.Context) error {
	// id := c.Param("id")
	projectName := c.FormValue("projectName")
	StartDate := c.FormValue("startdate")
	EndDate := c.FormValue("enddate")
	description := c.FormValue("description")
	golang := c.FormValue("golang")
	javascript := c.FormValue("javascript")
	reactjs := c.FormValue("reactjs")
	nodejs := c.FormValue("nodejs")
	technologies := []string{golang, javascript, reactjs, nodejs}
	image := c.FormValue("file")
	// Id, _ := strconv.Atoi(id)
	// if

	// if checkValue(technologies, "on") {
	// 	icon.NodeJs = true
	// }
	// if checkValue(technologies, "on") {
	// 	icon.Javascipt = true
	// }

	_, err := conect.Conn.Exec(context.Background(), "INSERT INTO tb_projects(p_name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)", projectName, StartDate, EndDate, description, technologies, image)

	if err != nil {
		fmt.Println("error guys")
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deletemyProject(c echo.Context) error {
	id := c.Param("id")

	Id, _ := strconv.Atoi(id)

	_, err1 := conect.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", Id)
	if err1 != nil {
		return c.JSON(500, err1.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(x echo.Context) error {
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("htmls/updateProject.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	dataProject := Project{}
	err1 := conect.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id=$1", Id).Scan(&dataProject.Id, &dataProject.ProjectName, &dataProject.StartDate, &dataProject.EndDate, &dataProject.Description, &dataProject.Technologies, &dataProject.Image)

	dataProject.Durations = Duration(dataProject.StartDate, dataProject.EndDate)

	if err1 != nil {
		return x.JSON(500, err.Error())
	}

	data := map[string]interface{}{
		"Project": dataProject,
		"Id":      id,
		// "startDateS": dataProject.StartDate.Format("2006-01-02"),
		// "endDateS":   dataProject.EndDate.Format("2006-01-02"),
	}
	return tmpl.Execute(x.Response(), data)
}

func updatedProject(x echo.Context) error {
	projectName := x.FormValue("projectName")
	StartDate := x.FormValue("startdate")
	EndDate := x.FormValue("enddate")
	description := x.FormValue("description")
	golang := x.FormValue("golang")
	javascript := x.FormValue("javascript")
	reactjs := x.FormValue("reactjs")
	nodejs := x.FormValue("nodejs")
	technologies := []string{golang, javascript, reactjs, nodejs}
	image := x.FormValue("file")
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	_, err := conect.Conn.Exec(context.Background(), "UPDATE tb_projects SET p_name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5 ,image = $6 WHERE id = $7", projectName, StartDate, EndDate, description, technologies, image, Id)

	if err != nil {
		fmt.Println("error guys")
		x.JSON(http.StatusInternalServerError, err.Error())
	}

	// UpdatedProject := Project{
	// 	Id:          Id,
	// 	Author:      "Adiwidiawan",
	// 	ProjectName: projectName,
	// 	// StartDate:    StartDate,
	// 	// EndDate:      EndDate,
	// 	Description:  description,
	// 	Technologies: []string{checklis1, checklis2, checklis3, checklis4},
	// 	Image:        "adi.jpg",
	// }

	// dataProject[Id] = UpdatedProject

	return x.Redirect(http.StatusMovedPermanently, "/")
}
func checkValue(x []string, checked string) bool {
	for _, data := range x {
		if data == checked {
			return true
		}
	}
	return false
}
