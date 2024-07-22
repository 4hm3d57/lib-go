package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"path/filepath"

	adminHandlers "lib/handlers/admin"
	registerHandlers "lib/handlers/register"
	studentHandlers "lib/handlers/student"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/frontend", "./templates")

	// Session store
	store := cookie.NewStore([]byte("secret_key"))
	r.Use(sessions.Sessions("session_id", store))

	// Load HTML templates
	adminGlob := filepath.Join("templates", "admin", "*.html")
	studentGlob := filepath.Join("templates", "student", "*.html")

	tmpl := template.Must(template.ParseGlob(adminGlob))
	tmpl = template.Must(tmpl.ParseGlob(studentGlob))
	r.SetHTMLTemplate(tmpl)

	// Public routes
	r.GET("/", func(c *gin.Context) {
		c.File("templates/login.html")
	})
	r.POST("/login", registerHandlers.LoginHandler)
	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/signup.html")
	})
	r.POST("/signup", registerHandlers.SignupHandler)

	// static files student and admin folders
	r.Static("/admin", "./templates/admin")
	r.Static("/student", "./templates/student")

	// Admin routes with authentication middleware

	r.GET("/admin-book", adminHandlers.AuthMiddleware(), func(c *gin.Context) {
		c.File("templates/admin/add-book.html")
	})
	r.POST("/admin-book", adminHandlers.AuthMiddleware(), adminHandlers.AddBookHandler)

	r.GET("/admin-users", adminHandlers.AuthMiddleware(), adminHandlers.RenderUserHandler)
	r.GET("/admin-all-books", adminHandlers.AuthMiddleware(), adminHandlers.BookHandler)
	r.GET("/admin-index", adminHandlers.AuthMiddleware(), adminHandlers.IndexPage)
	r.GET("/admin-recomm", adminHandlers.AuthMiddleware(), adminHandlers.RecommendationHandler)
	r.GET("/admin-request", adminHandlers.AuthMiddleware(), adminHandlers.RequestHandler)
	r.GET("/admin-message", adminHandlers.AuthMiddleware(), adminHandlers.MessageHandler)
	r.GET("/admin-curr", adminHandlers.AuthMiddleware(), adminHandlers.CurrHandler)

	// Student routes with authentication middleware
	r.GET("/student-index", studentHandlers.AuthMiddleware(), studentHandlers.IndexPage)
	r.GET("/student-all-books", studentHandlers.AuthMiddleware(), studentHandlers.StudentBookHandler)
	r.GET("/student-curr-books", studentHandlers.AuthMiddleware(), studentHandlers.StudentCurrHandler)
	r.GET("/student-message", studentHandlers.AuthMiddleware(), func(c *gin.Context) {
		c.File("templates/student/message.html")
	})
	r.POST("/student-message", studentHandlers.AuthMiddleware(), studentHandlers.InsertMessageHandler)
	r.GET("/student-recomm", studentHandlers.AuthMiddleware(), func(c *gin.Context) {
		c.File("templates/student/recommendation.html")
	})
	r.POST("/student-recomm", studentHandlers.AuthMiddleware(), studentHandlers.InsertRecommHandler)

	log.Fatal(r.Run(":9000"))
}
