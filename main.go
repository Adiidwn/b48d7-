package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ProjectName string
	Description string
	Author      string
	PostDate    string
}

var dataProject = []Project{
	{
		ProjectName: "Project 1",
		Description: "ini Project 1",
		Author:      "Adiwidiawan",
		PostDate:    "20July2023",
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
	e.GET("/updateProject/:id", updateProject)
	e.POST("/updatedProject", updatedProject)

	e.Logger.Fatal(e.Start("localhost:666"))
}

// handler
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
				ProjectName: data.ProjectName,
				Description: data.Description,
				Author:      data.Author,
				PostDate:    data.PostDate,
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
	// StartDate := c.FormValue("startdate")
	// EndDate := c.FormValue("enddate")
	// checklis1 := c.FormValue("checklis1")
	// checklis2 := c.FormValue("checklis2")
	// checklis3 := c.FormValue("checklis3")
	// checklis4 := c.FormValue("checklis4")

	newProject := Project{
		ProjectName: projectName,
		Description: description,
		Author:      "Adiwidiawan",
		PostDate:    "27 July 2023",
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

func updateProject(x echo.Context) error {
	tmplate, _ := template.ParseFiles("htmls/updateProject.html")
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)
	projectName := x.FormValue("projectName")
	description := x.FormValue("description")

	UpdatedProject := Project{
		ProjectName: projectName,
		Description: description,
		Author:      "Adiwidiawan",
		PostDate:    "27 July 2023",
	}

	for index, data := range dataProject {
		if index == Id {
			UpdatedProject = Project{
				ProjectName: data.ProjectName,
				Description: data.Description,
				Author:      data.Author,
				PostDate:    data.PostDate,
			}
		}
	}

	dataProject = append(dataProject[:Id], UpdatedProject)

	return tmplate.Execute(x.Response(), nil)
}
func updatedProject(x echo.Context) error {

	return x.Redirect(http.StatusMovedPermanently, "/")
}

// StartDate := c.FormValue("startdate")
// EndDate := c.FormValue("enddate")
// checklis1 := c.FormValue("checklis1")
// checklis2 := c.FormValue("checklis2")
// checklis3 := c.FormValue("checklis3")
// checklis4 := c.FormValue("checklis4")

// return c.Redirect(http.StatusMovedPermanently, "/")
