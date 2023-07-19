package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/addProject", addProject)
	e.GET("/Testimonials", Testimonials)
	e.GET("/ProjectDetail/:id", blogDetail)
	e.POST("/addmyProject", addmyProject)
	e.Logger.Fatal(e.Start("localhost:666"))
}

// handler
func home(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/index.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
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
func Testimonials(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/testimonials.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	return tmplate.Execute(x.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id := c.Param("id") // misal : 1

	tmpl, err := template.ParseFiles("htmls/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogDetail := map[string]interface{}{ // interface -> tipe data apapun
		"Id":      id,
		"Title":   "Dumbways ID memang keren",
		"Content": "Dumbways ID adalah bootcamp terbaik sedunia seakhirat!",
	}

	return tmpl.Execute(c.Response(), blogDetail)
}
func addmyProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startdate")
	endDate := c.FormValue("enddate")
	description := c.FormValue("description")
	checklis1 := c.FormValue("checklis1")
	checklis2 := c.FormValue("checklis2")
	checklis3 := c.FormValue("checklis3")
	checklis4 := c.FormValue("checklis4")

	fmt.Println("Project Name: ", projectName)
	fmt.Println("Start Date: ", startDate)
	fmt.Println("End Date: ", endDate)
	fmt.Println("Description: ", description)
	fmt.Println("Tecnologies: ", "1 :", checklis1, "2 :", checklis2, "3 :", checklis3, "4 :", checklis4)

	return c.Redirect(http.StatusMovedPermanently, "/addProject")
}
