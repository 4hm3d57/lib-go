package main

import (
	adminHandlers "lib/handlers/admin"
	// studentHandlers "lib/handlers/student"
	"github.com/gin-gonic/gin"
	"html/template"
	"lib/handlers"
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

	r.POST("/login", handlers.LoginHandler)

	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/signup.html")
	})

	r.POST("/signup", handlers.SignupHandler)

	r.Static("/admin", "./templates/admin")
	r.Static("/student", "./templates/student")

	r.GET("/admin-book", func(c *gin.Context) {
		c.File("templates/admin/add-book.html")
	})

	r.POST("/admini-book", adminHandlers.AddBookHandler)

	// route for rendering the user table
	r.GET("/admin-users", adminHandlers.RenderUserHandler)
	r.GET("/admin-all-books", adminHandlers.BookHandler)
	r.GET("/admin-index", adminHandlers.IndexHandler)
	r.GET("/admin-recomm", adminHandlers.RecommendationHandler)
	r.GET("/admin-request", adminHandlers.RequestHandler)
	r.GET("/admin-message", adminHandlers.MessageHandler)
	r.GET("/admin-curr", adminHandlers.CurrHandler)

	log.Fatal(r.Run(":3000"))
}
