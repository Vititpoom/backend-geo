package routes

import (
	"backend-geo/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/health", handlers.HealthCheck)

	// Location endpoints
	api.GET("/locations", handlers.GetLocations)
	api.POST("/locations", handlers.CreateLocation)
	api.DELETE("/locations/:id", handlers.DeleteLocation)
	api.PUT("/locations/:id", handlers.UpdateLocation)
}

