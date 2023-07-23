package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	conect "vodlab/com/connection"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id           int
	Author       string
	ProjectName  string
	Duration     int
	StartDate    any
	EndDate      any
	Description  string
	Technologies []string
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

	dataProject, dataerror := conect.Conn.Query(context.Background(), "SELECT * FROM tb_projects")

	if dataerror != nil {
		return x.JSON(500, err.Error())
	}

	var resultProject []Project
	for dataProject.Next() {
		each := Project{}
		dataProject.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies, &each.Image)
		if dataerror != nil {
			return x.JSON(500, err.Error())
		}
		each.Author = "Adiwidiawan"
		// t1 := each.StartDate
		// t2 := each.EndDate
		// diff:=t1.Sub(t2)
		fmt.Println(each.StartDate)
		resultProject = append(resultProject, each)
	}

	data := map[string]interface{}{
		"Project": resultProject,
	}
	println(resultProject)
	return tmplate.Execute(x.Response(), data)
}

func printf(t1, t2 any) {
	panic("unimplemented")
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
				Id:           idInt,
				Author:       data.Author,
				ProjectName:  data.ProjectName,
				StartDate:    data.StartDate,
				EndDate:      data.EndDate,
				Description:  data.Description,
				Technologies: data.Technologies,
				Image:        data.Image,
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
	checklis1 := c.FormValue("checklis1")
	checklis2 := c.FormValue("checklis2")
	checklis3 := c.FormValue("checklis3")
	checklis4 := c.FormValue("checklis4")
	StartDatee, _ := strconv.Atoi(StartDate)
	EndDatee, _ := strconv.Atoi(EndDate)

	newProject := Project{
		Id:           0,
		Author:       "Adiwidiawan",
		ProjectName:  projectName,
		StartDate:    StartDatee,
		EndDate:      EndDatee,
		Description:  description,
		Technologies: []string{checklis1, checklis2, checklis3, checklis4},
		Image:        "adi.jpg",
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

func newFunction(resultProject []Project) {
	resultProject = dataProject
}

func editProject(x echo.Context) error {
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	ProjectEdit := Project{}

	for index, data := range dataProject {
		if index == Id {
			ProjectEdit = Project{
				Id:           Id,
				Author:       data.Author,
				ProjectName:  data.ProjectName,
				StartDate:    data.StartDate,
				EndDate:      data.EndDate,
				Description:  data.Description,
				Technologies: data.Technologies,
				Image:        data.Image,
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
	// StartDate := x.FormValue("startdate")
	// EndDate := x.FormValue("enddate")
	checklis1 := x.FormValue("checklis1")
	checklis2 := x.FormValue("checklis2")
	checklis3 := x.FormValue("checklis3")
	checklis4 := x.FormValue("checklis4")

	UpdatedProject := Project{
		Id:          Id,
		Author:      "Adiwidiawan",
		ProjectName: projectName,
		// StartDate:    StartDate,
		// EndDate:      EndDate,
		Description:  description,
		Technologies: []string{checklis1, checklis2, checklis3, checklis4},
		Image:        "adi.jpg",
	}

	dataProject[Id] = UpdatedProject

	return x.Redirect(http.StatusMovedPermanently, "/")
}
