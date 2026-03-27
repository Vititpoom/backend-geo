package handlers

import (
	"net/http"

	"backend-geo/config"
	"backend-geo/database"
	"backend-geo/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getCollection() *mongo.Collection {
	cfg := config.Load()
	return database.GetCollection(database.Client, cfg.DBName, "locations")
}

// GetLocations returns all locations as a JSON array
func GetLocations(c echo.Context) error {
	collection := getCollection()

	cursor, err := collection.Find(c.Request().Context(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch locations",
		})
	}
	defer cursor.Close(c.Request().Context())

	var locations []models.Location
	if err := cursor.All(c.Request().Context(), &locations); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to decode locations",
		})
	}

	// Return empty array instead of null when no results
	if locations == nil {
		locations = []models.Location{}
	}

	return c.JSON(http.StatusOK, locations)
}

// CreateLocation creates a new location from a GeoJSON Feature payload
func CreateLocation(c echo.Context) error {
	var location models.Location
	if err := c.Bind(&location); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Enforce GeoJSON Feature spec
	location.Type = "Feature"
	if location.Geometry.Type == "" {
		location.Geometry.Type = "Point"
	}

	// Validate required fields
	if len(location.Geometry.Coordinates) != 2 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Geometry coordinates must be an array of [longitude, latitude]",
		})
	}
	if location.Properties.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Properties.name is required",
		})
	}

	collection := getCollection()

	// Clear ID so MongoDB generates one
	location.ID = bson.ObjectID{}

	result, err := collection.InsertOne(c.Request().Context(), location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to insert location",
		})
	}

	location.ID = result.InsertedID.(bson.ObjectID)

	return c.JSON(http.StatusCreated, location)
}

// DeleteLocation removes a location by its ID
func DeleteLocation(c echo.Context) error {
	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid location ID format",
		})
	}

	collection := getCollection()

	result, err := collection.DeleteOne(c.Request().Context(), bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete location",
		})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Location not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Location deleted successfully",
	})
}

// UpdateLocation updates an existing location by its ID
func UpdateLocation(c echo.Context) error {
	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid location ID format",
		})
	}

	var location models.Location
	if err := c.Bind(&location); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Enforce GeoJSON Feature spec
	location.Type = "Feature"
	if location.Geometry.Type == "" {
		location.Geometry.Type = "Point"
	}

	// Validate required fields
	if len(location.Geometry.Coordinates) != 2 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Geometry coordinates must be an array of [longitude, latitude]",
		})
	}
	if location.Properties.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Properties.name is required",
		})
	}

	collection := getCollection()

	update := bson.M{
		"$set": bson.M{
			"type":       "Feature",
			"geometry":   location.Geometry,
			"properties": location.Properties,
		},
	}

	result, err := collection.UpdateOne(c.Request().Context(), bson.M{"_id": objectID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update location",
		})
	}
	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Location not found",
		})
	}

	location.ID = objectID
	return c.JSON(http.StatusOK, location)
}