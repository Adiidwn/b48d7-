package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"unicode"
	conect "vodlab/com/connection"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	AuthorId     int
	Idrelate     bool
}
type Users struct {
	Id       int
	Username string
	Email    string
	Password string
	Role     string
}
type UserLoginSessi struct {
	Islogin bool
	Name    string
	Roles   bool
	Id      bool
}

var userLoginSessi = UserLoginSessi{}

var users = []Users{
	// Username:,
	// Email:,
	// Password:,
	// Role:,
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

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("adi"))))
	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/testimonials", testimonials)
	e.GET("/about", about)
	e.GET("/projectDetail/:id", projectDetail)

	// Delete Project
	e.POST("/deletemyProject/:id", deletemyProject)

	// Add Project
	e.GET("/addProject", addProject)
	e.POST("/addmyProject", addmyProject)

	// Update
	e.GET("/updateProject/:id", editProject)
	e.POST("/updatedProject/:id", updatedProject)

	// Login/Register
	e.GET("/form-login", formLogin)
	e.POST("/login", login)

	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	e.POST("/logout", logout)

	e.Logger.Fatal(e.Start("localhost:666"))
}

// handler
// HOME
func home(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/index.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}
	// var datanotstring sql.NullString
	// var startDate = time.DateOnly
	// var endDate = time.DateOnly
	sessi, _ := session.Get("session", x)
	data1, dataerror := conect.Conn.Query(context.Background(), "SELECT tb_projects.id, tb_users.username,tb_projects.author_id, tb_projects.p_name, tb_projects.description,tb_projects.technologies, tb_projects.image,tb_projects.start_date ,tb_projects.end_date FROM tb_projects LEFT JOIN tb_users ON tb_projects.author_id = tb_users.id ")

	if dataerror != nil {
		return x.JSON(500, err.Error())
	}

	dataProject = []Project{}
	for data1.Next() {
		var each = Project{}

		data1.Scan(&each.Id, &each.Author, &each.AuthorId, &each.ProjectName, &each.Description, &each.Technologies, &each.Image, &each.StartDate, &each.EndDate)
		// if dataerror != nil {
		// 	return x.JSON(500, err.Error())
		// }
		// fmt.Println("author/username", each.Author, "authorID", each.AuthorId)
		// each.Author = datanotstring.String
		// each.StartDate = time.DateOnly(startDate)
		// each.EndDate = endDate.String
		each.Durations = Duration(each.StartDate, each.EndDate)
		// fmt.Println("startDate:", each.StartDate, "endDate:", each.EndDate)
		// fmt.Println("id:", each.Id, "namaP;", each.Author, "namaproject", each.ProjectName, "duration", each.Durations, "description", each.Description, "tech", each.Technologies, "image", each.Image)

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

		if sessi.Values["id"] == each.AuthorId {
			each.Idrelate = true
		}
		// fmt.Println("values id ", sessi.Values["id"], "author id ", each.AuthorId)
		// t1 := each.StartDate
		// t2 := each.EndDate
		// diff:=t1.Sub(t2)

		dataProject = append(dataProject, each)
	}

	// sessi, sessierr := session.Get("session", x)
	// if sessierr != nil {
	// 	return x.JSON(500, sessierr.Error())
	// }
	// fmt.Println("message:", sessi.Values["message"])
	// fmt.Println("status:", sessi.Values["status"])

	// flash := map[string]interface{}{
	// 	"FlashM": sessi.Values["message"],
	// 	"FlashS": sessi.Values["status"],
	// }

	// if err != nil {
	// 	return x.JSON(500, err.Error())
	// }
	// return tmplate.Execute(x.Response(), flash)

	if sessi.Values["Islogin"] != true {
		userLoginSessi.Islogin = false
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}

	if sessi.Values["role"] != "admin" {
		userLoginSessi.Roles = false
	} else {
		userLoginSessi.Roles = true
	}
	// dataProject := Project{}
	// if sessi.Values["id"] == dataProject.AuthorId {
	// 	userLoginSessi.Id = true
	// } else {
	// 	userLoginSessi.Id = false
	// }
	// fmt.Println("sessi.value ID", sessi.Values["id"], "authorID", dataProject.AuthorId)
	// println("values id:", sessi.Values["id"], "project ID :", projectId.AuthorId)
	// fmt.Println(userLoginSessi.Name)
	flash := map[string]interface{}{
		"Project":        dataProject,
		"UserLoginSessi": userLoginSessi,

		// "FlashM":         sessi.Values["message"],
		// "FlashS":         sessi.Values["status"],
	}
	delete(sessi.Values, "message")
	delete(sessi.Values, "status")
	sessi.Save(x.Request(), x.Response())

	return tmplate.Execute(x.Response(), flash)
}

// DURATION CALCULATION
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

// CONTACT
func contact(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/contact.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	sessi, _ := session.Get("session", x)

	if sessi.Values["Islogin"] != true {
		userLoginSessi.Islogin = false
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}
	// fmt.Println(userLoginSessi.Name)
	flash := map[string]interface{}{
		"Project":        dataProject,
		"UserLoginSessi": userLoginSessi,
	}

	sessi.Save(x.Request(), x.Response())
	return tmplate.Execute(x.Response(), flash)
}

// ADDPROJECT FORM
func addProject(x echo.Context) error {
	sessi, _ := session.Get("session", x)

	if sessi.Values["Islogin"] != true {
		fmt.Println("addproject form userloginsesi.islogin =  ", userLoginSessi.Islogin)
		fmt.Println("addproject form valueislogin =  ", sessi.Values["Islogin"])
		return x.Redirect(http.StatusMovedPermanently, "/")
	}

	tmplate, err := template.ParseFiles("htmls/project1.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	// fmt.Println(userLoginSessi.Name)
	flash := map[string]interface{}{
		"Project":        dataProject,
		"UserLoginSessi": userLoginSessi,
	}
	delete(sessi.Values, "message")
	delete(sessi.Values, "status")
	sessi.Save(x.Request(), x.Response())

	return tmplate.Execute(x.Response(), flash)
}

// ABOUT FORM
func about(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/about.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	sessi, _ := session.Get("session", x)

	if sessi.Values["Islogin"] != true {
		userLoginSessi.Islogin = false
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}
	// fmt.Println(userLoginSessi.Name)
	flash := map[string]interface{}{
		"Project":        dataProject,
		"UserLoginSessi": userLoginSessi,
	}

	sessi.Save(x.Request(), x.Response())
	return tmplate.Execute(x.Response(), flash)
}

// TESTIMONIALS FORM
func testimonials(x echo.Context) error {
	tmplate, err := template.ParseFiles("htmls/testimonials.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	sessi, _ := session.Get("session", x)

	if sessi.Values["Islogin"] != true {
		userLoginSessi.Islogin = false
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}
	// fmt.Println(userLoginSessi.Name)
	flash := map[string]interface{}{
		"Project":        dataProject,
		"UserLoginSessi": userLoginSessi,
	}

	sessi.Save(x.Request(), x.Response())
	return tmplate.Execute(x.Response(), flash)
}

// PROOJECT DETAIL FORM
func projectDetail(c echo.Context) error {
	sessi, _ := session.Get("session", c)

	if userLoginSessi.Islogin != true {
		userLoginSessi.Islogin = false
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}
	if sessi.Values["role"] != "admin" {
		userLoginSessi.Roles = false
	} else {
		userLoginSessi.Roles = true
	}

	id := c.Param("id") // misal : 1

	tmpl, err := template.ParseFiles("htmls/ProjectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Id, _ := strconv.Atoi(id)

	dataProject := Project{}
	// var datanotstring sql.NullString
	err1 := conect.Conn.QueryRow(context.Background(), "SELECT tb_projects.id, tb_users.username,tb_projects.author_id, tb_projects.p_name, tb_projects.description,tb_projects.technologies, tb_projects.image,tb_projects.start_date ,tb_projects.end_date FROM tb_projects LEFT JOIN tb_users ON tb_projects.author_id = tb_users.id  WHERE tb_projects.id=$1", Id).Scan(&dataProject.Id, &dataProject.Author, &dataProject.AuthorId, &dataProject.ProjectName, &dataProject.Description, &dataProject.Technologies, &dataProject.Image, &dataProject.StartDate, &dataProject.EndDate)

	// dataProject.Author = datanotstring.String
	// fmt.Println("datanotstring", datanotstring)
	fmt.Println("authorid", dataProject.AuthorId)
	fmt.Println("value id", sessi.Values["id"])

	dataProject.Durations = Duration(dataProject.StartDate, dataProject.EndDate)

	if checkValue(dataProject.Technologies, "ReactJs") {
		dataProject.ReactJs = true
	}
	if checkValue(dataProject.Technologies, "Golang") {
		dataProject.Golang = true
	}
	if checkValue(dataProject.Technologies, "NodeJs") {
		dataProject.NodeJs = true
	}
	if checkValue(dataProject.Technologies, "Javascript") {
		dataProject.Javascipt = true
	}
	if sessi.Values["id"] == dataProject.AuthorId {
		dataProject.Idrelate = true
	}

	if err1 != nil {
		return c.JSON(500, err.Error())
	}

	// SESSION

	// if sessi.Values["Islogin"] != true {
	// 	userLoginSessi.Islogin = false
	// } else {
	// 	userLoginSessi.Islogin = true
	// 	userLoginSessi.Name = sessi.Values["username"].(string)
	// }
	// fmt.Println(userLoginSessi.Name)

	sessi.Save(c.Request(), c.Response())

	data := map[string]interface{}{
		"Project":        dataProject,
		"Id":             id,
		"startDateS":     dataProject.StartDate.Format("2006-01-02"),
		"endDateS":       dataProject.EndDate.Format("2006-01-02"),
		"UserLoginSessi": userLoginSessi,
	}

	return tmpl.Execute(c.Response(), data)
}

// ADD PROJECT POST
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
	sessi, _ := session.Get("session", c)

	_, err := conect.Conn.Exec(context.Background(), "INSERT INTO tb_projects(p_name, start_date, end_date, description, technologies, image ,author_id) VALUES ($1, $2, $3, $4, $5, $6 ,$7)", projectName, StartDate, EndDate, description, technologies, image, sessi.Values["id"].(int))
	fmt.Println("id:", sessi.Values["id"])
	if err != nil {
		fmt.Println("error guys")
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// DELET PROJECT
func deletemyProject(c echo.Context) error {
	id := c.Param("id")

	Id, _ := strconv.Atoi(id)

	_, err1 := conect.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", Id)
	if err1 != nil {
		return c.JSON(500, err1.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// EDIT PROJECT FORM
func editProject(x echo.Context) error {
	sessi, _ := session.Get("session", x)
	if sessi.Values["Islogin"] != true {
		return x.Redirect(http.StatusMovedPermanently, "/")
	} else {
		userLoginSessi.Islogin = true
		userLoginSessi.Name = sessi.Values["username"].(string)
	}
	id := x.Param("id")
	Id, _ := strconv.Atoi(id)

	tmpl, err := template.ParseFiles("htmls/updateProject.html")

	if err != nil {
		return x.JSON(500, err.Error())
	}

	var datanotstring sql.NullString
	dataProject := Project{}
	err1 := conect.Conn.QueryRow(context.Background(), "SELECT tb_projects.id, tb_users.username, tb_projects.p_name, tb_projects.description,tb_projects.technologies, tb_projects.image,tb_projects.start_date ,tb_projects.end_date FROM tb_projects LEFT JOIN tb_users ON tb_projects.id = tb_users.id WHERE tb_projects.id=$1", Id).Scan(&dataProject.Id, &datanotstring, &dataProject.ProjectName, &dataProject.Description, &dataProject.Technologies, &dataProject.Image, &dataProject.StartDate, &dataProject.EndDate)
	dataProject.Author = datanotstring.String
	dataProject.Durations = Duration(dataProject.StartDate, dataProject.EndDate)

	if err1 != nil {
		return x.JSON(500, err.Error())
	}

	data := map[string]interface{}{
		"Project":        dataProject,
		"Id":             id,
		"UserLoginSessi": userLoginSessi,

		// "startDateS": dataProject.StartDate.Format("2006-01-02"),
		// "endDateS":   dataProject.EndDate.Format("2006-01-02"),
	}

	delete(sessi.Values, "message")
	delete(sessi.Values, "status")
	sessi.Save(x.Request(), x.Response())

	return tmpl.Execute(x.Response(), data)
}

// EDIT PROJECT POST
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

// CHECKVALUE CALCULATION
func checkValue(x []string, checked string) bool {
	for _, data := range x {
		if data == checked {
			return true
		}
	}
	return false
}

// LOGIN/REGISTER

func formLogin(x echo.Context) error {

	if userLoginSessi.Islogin != false {
		return x.Redirect(http.StatusMovedPermanently, "/")
	}

	// if sessi1.Values["Islogin"] != false {
	// 	userLoginSessi.Islogin = true
	// 	return redirectWMessage(x, "Tidak boleh Login kembali", true, "/")
	// } else {
	// 	userLoginSessi.Islogin = false
	// }

	tmplate, err := template.ParseFiles("htmls/form-login.html")
	if err != nil {
		return x.JSON(500, err.Error())
	}
	sessi1, _ := session.Get("session", x)

	flash := map[string]interface{}{
		"FlashM": sessi1.Values["message"],
		"FlashS": sessi1.Values["status"],
	}

	delete(sessi1.Values, "message")
	delete(sessi1.Values, "status")
	sessi1.Save(x.Request(), x.Response())

	return tmplate.Execute(x.Response(), flash)
}

func login(x echo.Context) error {
	email := x.FormValue("email")
	password := x.FormValue("password")

	dataUser := Users{}
	err := conect.Conn.QueryRow(context.Background(), "SELECT id,role,email,username,password FROM tb_users WHERE email=$1", email).Scan(&dataUser.Id, &dataUser.Role, &dataUser.Email, &dataUser.Username, &dataUser.Password)

	if err != nil {
		fmt.Println("isi login ID", err)
		return redirectWMessage(x, "Login Failed !", false, "/form-login")
	}

	bcrypterr := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))

	if bcrypterr != nil {
		fmt.Println("isi pasword compare has password", bcrypterr)
		return redirectWMessage(x, "Login Failed !", false, "/form-login")
	}

	sessi, _ := session.Get("session", x)
	sessi.Options.MaxAge = 1080 //10800 = 3jam
	sessi.Values["message"] = "Login Succes !"
	sessi.Values["status"] = true
	sessi.Values["Islogin"] = true
	sessi.Values["id"] = dataUser.Id
	sessi.Values["username"] = dataUser.Username
	sessi.Values["role"] = dataUser.Role
	sessi.Values["email"] = dataUser.Email
	fmt.Println("beres login value id apaan nih", sessi.Values["id"])
	sessi.Save(x.Request(), x.Response())

	return x.Redirect(http.StatusMovedPermanently, "/")
}

func formRegister(x echo.Context) error {
	// sessi1, _ := session.Get("session", x)
	if userLoginSessi.Islogin != false {
		return x.Redirect(http.StatusMovedPermanently, "/")
	}
	// if sessi1.Values["Islogin"] != false {
	// 	userLoginSessi.Islogin = true
	// 	return redirectWMessage(x, "Tidak boleh Register kembali", true, "/")
	// } else {
	// 	userLoginSessi.Islogin = false
	// }

	tmplate, err := template.ParseFiles("htmls/form-register.html")
	if err != nil {
		return x.JSON(500, err.Error())
	}
	// fmt.Println("message:", sessi.Values["message"])
	// fmt.Println("status:", sessi.Values["status"])
	sessi, _ := session.Get("session", x)

	flash := map[string]interface{}{
		"FlashM": sessi.Values["message"],
		"FlashS": sessi.Values["status"],
	}

	delete(sessi.Values, "message")
	delete(sessi.Values, "Status")
	sessi.Save(x.Request(), x.Response())

	return tmplate.Execute(x.Response(), flash)
}

func register(x echo.Context) error {

	usernameV := x.FormValue("username")
	emailV := x.FormValue("email")
	passwordV := x.FormValue("password")

	// var ErrHashTooShort = errors.New("crypto/bcrypt: hashedSecret too short to be a bcrypted password")

	passwordHashed, pwerror := bcrypt.GenerateFromPassword([]byte(passwordV), 7)

	if pwerror != nil {
		fmt.Println("password gagal ter enkripsi")
		return x.JSON(http.StatusInternalServerError, pwerror.Error())
	}

	_, err := conect.Conn.Exec(context.Background(), "INSERT INTO tb_users (username, email, password ,role) VALUES ($1, $2, $3 ,$4)", usernameV, emailV, passwordHashed, "user")

	if err != nil {
		fmt.Println("error INSERT DATA guys")
		return redirectWMessage(x, "Regist Failed !", false, "/form-register")
	}

	return redirectWMessage(x, "Regist Succes !", true, "/form-login")
}

func logout(x echo.Context) error {
	sessi, _ := session.Get("session", x)

	sessi.Options.MaxAge = -1
	sessi.Values["Islogin"] = false
	userLoginSessi.Islogin = false
	sessi.Save(x.Request(), x.Response())
	// if sessierr != nil {
	// 	fmt.Println("error logout")
	// 	return redirectWMessage(x, "Logout Failed !", false, "/form-login")
	// }
	fmt.Println("BERES LOGOUT APA NIH", sessi.Values["Islogin"])
	return redirectWMessage(x, "Logout Succesfull", false, "/")
}

func redirectWMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			//return false, false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}
