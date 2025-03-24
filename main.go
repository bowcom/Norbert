package main

import (
	"errors"
	"html/template"
	"log/slog"
	"net/http"

	//	"Norbert/views"

	//	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		slog.Info("Warning: .env file not found")
	}

	// Echo instance
	e := echo.New()

	// Static files
	e.Static("/assets", "assets")

	//	e.GET("/", homeHandler)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", homeHandler)
	e.GET("/home", homeHandler)
	e.GET("/about", aboutHandler)
	e.GET("/login", loginHandler)
	e.GET("/test", testHandler)
	e.GET("/dashboard", dashboardHandler)
	// Start server
	if err := e.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

func homeHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/home.html", "templates/boiler.html"))
	return tmpl.ExecuteTemplate(c.Response().Writer, "home.html", nil)
}

func loginHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/login.html", "templates/boiler.html"))
	return tmpl.ExecuteTemplate(c.Response().Writer, "login.html", nil)

}
func dashboardHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html", "templates/boiler.html"))
	return tmpl.ExecuteTemplate(c.Response().Writer, "dashboard.html", nil)
}

func aboutHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/about.html", "templates/boiler.html"))
	return tmpl.ExecuteTemplate(c.Response().Writer, "about.html", nil)
}
func testHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/test.html", "templates/boiler.html"))
	return tmpl.ExecuteTemplate(c.Response().Writer, "test.html", nil)
}
