package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Author      string
	ProjectName string
	Durations   string
	StartDate   string
	EndDate     string
	Description string
	NodeJs      bool
	ReactJs     bool
	Golang      bool
	Javascript  bool
	Image       string
}

var dataProject = []Project{
	{
		Author:      "Adiwidiawan",
		ProjectName: "Project 1",
		Durations:   Duration("2020-01-15", "2020-02-15"),
		StartDate:   "2020-01-15",
		EndDate:     "2020-02-15",
		Description: "HALOO GUYSSSSS",
		NodeJs:      true,
		ReactJs:     false,
		Golang:      false,
		Javascript:  false,
		Image:       "adi.jpg",
	},
	{
		Author:      "Adiwidiawan",
		ProjectName: "Project 2",
		Durations:   Duration("2020-01-15", "2020-02-15"),
		StartDate:   "2020-01-15",
		EndDate:     "2020-01-22",
		Description: "HALOO GUYSSSSS",
		NodeJs:      true,
		ReactJs:     true,
		Golang:      false,
		Javascript:  false,
		Image:       "adi2.jpg",
	},
	{
		Author:      "Adiwidiawan",
		ProjectName: "Project 3",
		Durations:   Duration("2020-01-15", "2020-02-15"),
		StartDate:   "2020-01-15",
		EndDate:     "2021-02-15",
		Description: "HALOO GUYSSSSS",
		NodeJs:      true,
		ReactJs:     true,
		Golang:      true,
		Javascript:  false,
		Image:       "adi3.jpg",
	},
}

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/addProject", addProject)
	e.GET("/testimonials", testimonials)
	e.GET("/projectDetail/:id", projectDetail)
	e.POST("/addmyProject", addmyProject)
	e.POST("/deletemyProject/:id", deletemyProject)
	e.GET("/updateProject/:id", editProject)
	e.POST("/updatedProject/:id", updatedProject)

	e.Logger.Fatal(e.Start("localhost:666"))
}

func home(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/index.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	data := map[string]interface{}{
		"Project": dataProject,
	}
	return tmplate.Execute(x.Response(), data)
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

	idInt, _ := strconv.Atoi(id)
	projectDetail := Project{}

	for index, data := range dataProject {
		if index == idInt {
			projectDetail = Project{
				Author:      data.Author,
				ProjectName: data.ProjectName,
				Durations:   data.Durations,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				NodeJs:      data.NodeJs,
				ReactJs:     data.ReactJs,
				Golang:      data.Golang,
				Javascript:  data.Javascript,
				Image:       data.Image,
			}
		}
	}

	data := map[string]interface{}{
		"Project": projectDetail,
		"Id":      id,
	}

	return tmpl.Execute(c.Response(), data)
}

func addmyProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	description := c.FormValue("description")
	StartDate := c.FormValue("startdate")
	EndDate := c.FormValue("enddate")
	nodeJs := c.FormValue("nodeJs")
	reactJs := c.FormValue("reactJs")
	golang := c.FormValue("golang")
	javascript := c.FormValue("javascript")
	image := c.FormValue("image")

	newProject := Project{
		Author:      "Adiwidiawan",
		ProjectName: projectName,
		Durations:   Duration(StartDate, EndDate),
		StartDate:   StartDate,
		EndDate:     EndDate,
		Description: description,
		NodeJs:      (nodeJs == "nodeJs"),
		ReactJs:     (reactJs == "reactJs"),
		Golang:      (golang == "golang"),
		Javascript:  (javascript == "javascript"),
		Image:       image,
	}

	dataProject = append(dataProject, newProject)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deletemyProject(c echo.Context) error {
	id := c.Param("id")

	// append

	// slice -> 3 struct (+ 1 struct)

	// slice = append(slice, structlagi)

	// fmt.Println("persiapan delete index : ", id)

	Id, _ := strconv.Atoi(id)

	dataProject = append(dataProject[:Id], dataProject[Id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(x echo.Context) error {
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	ProjectEdit := Project{}

	for index, data := range dataProject {
		if index == Id {
			ProjectEdit = Project{
				Author:      data.Author,
				ProjectName: data.ProjectName,
				Durations:   data.Durations,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				NodeJs:      data.NodeJs,
				ReactJs:     data.ReactJs,
				Golang:      data.Golang,
				Javascript:  data.Javascript,
				Image:       data.Image,
			}
		}
	}
	data := map[string]interface{}{
		"Project": ProjectEdit,
		"Id":      id,
	}

	tmpl, err := template.ParseFiles("htmls/updateProject.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	return tmpl.Execute(x.Response(), data)
}

func updatedProject(x echo.Context) error {

	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	projectName := x.FormValue("projectName")
	description := x.FormValue("description")
	StartDate := x.FormValue("startdate")
	EndDate := x.FormValue("enddate")
	nodeJs := x.FormValue("nodeJs")
	reactJs := x.FormValue("reactJs")
	golang := x.FormValue("golang")
	javascript := x.FormValue("javascript")
	image := x.FormValue("image")

	UpdatedProject := Project{
		Author:      "Adiwidiawan",
		ProjectName: projectName,
		Durations:   Duration(StartDate, EndDate),
		StartDate:   StartDate,
		EndDate:     EndDate,
		Description: description,
		NodeJs:      (nodeJs == "nodeJs"),
		ReactJs:     (reactJs == "reactJs"),
		Golang:      (golang == "golang"),
		Javascript:  (javascript == "javascript"),
		Image:       image,
	}

	dataProject[Id] = UpdatedProject

	return x.Redirect(http.StatusMovedPermanently, "/")
}

func Duration(StartDate string, EndDate string) string {
	StartDatee, _ := time.Parse("2006-01-02", StartDate)
	EndDatee, _ := time.Parse("2006-01-02", EndDate)

	diff := EndDatee.Sub(StartDatee)
	day := int(diff.Hours() / 24)
	week := day / 7
	month := day / 30
	year := month / 12
	if day < 7 {
		return strconv.Itoa(day) + "day"
	}
	if week < 4 {
		return strconv.Itoa(week) + "week"
	}
	if month < 4 {
		return strconv.Itoa(month) + "month"
	}
	return strconv.Itoa(year) + "year"

}
