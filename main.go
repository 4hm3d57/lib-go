package main

import (
	adminHandlers "lib/handlers/admin"
	studentHandlers "lib/handlers/student"
	registerHandlers "lib/handlers/register"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"path/filepath"
)

func main() {

	r := gin.Default()

	// serve login.html as root URI
	r.Static("/frontend", "./templates")

	adminGlob := filepath.Join("templates", "admin", "*.html")
	studentGlob := filepath.Join("templates", "student", "*.html")

	// load the html template
	tmpl := template.Must(template.ParseGlob(adminGlob))
	tmpl = template.Must(tmpl.ParseGlob(studentGlob))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) {
		c.File("templates/login.html")
	})

	r.POST("/login", registerHandlers.LoginHandler)

	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/signup.html")
	})

	r.POST("/signup", registerHandlers.SignupHandler)

	r.Static("/admin", "./templates/admin")
	r.Static("/student", "./templates/student")

	r.GET("/admin-book", func(c *gin.Context) {
		c.File("templates/admin/add-book.html")
	})

	r.POST("/admin-book", adminHandlers.AddBookHandler)

	r.GET("/student-message", func(c *gin.Context){
		c.File("templates/student/message.html")
	})

	r.POST("/student-message", studentHandlers.InsertMessageHandler)

	r.GET("/student-recomm", func(c *gin.Context){
		c.File("templates/student/recommendation.html")
	})

	r.POST("/student-recomm", studentHandlers.InsertRecommHandler)

	// admin routes
	r.GET("/admin-users", adminHandlers.RenderUserHandler)
	r.GET("/admin-all-books", adminHandlers.BookHandler)
	r.GET("/admin-index", adminHandlers.IndexHandler)
	r.GET("/admin-recomm", adminHandlers.RecommendationHandler)
	r.GET("/admin-request", adminHandlers.RequestHandler)
	r.GET("/admin-message", adminHandlers.MessageHandler)
	r.GET("/admin-curr", adminHandlers.CurrHandler)

	// student routes
	r.GET("/student-index", studentHandlers.StudentIndexHandler)
	r.GET("/student-all-books", studentHandlers.StudentBookHandler)
	r.GET("/student-curr-books", studentHandlers.StudentCurrHandler)

	log.Fatal(r.Run(":8080"))
}
