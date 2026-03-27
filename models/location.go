package models

import "go.mongodb.org/mongo-driver/v2/bson"

// GeoJSON Geometry represents a GeoJSON geometry object (Point type)
type Geometry struct {
	Type        string    `json:"type" bson:"type"`               // Must be "Point"
	Coordinates []float64 `json:"coordinates" bson:"coordinates"` // [longitude, latitude]
}

// Properties holds the metadata for a Location feature
type Properties struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

// Location follows the GeoJSON Feature specification
type Location struct {
	ID         bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       string        `json:"type" bson:"type"` // Must be "Feature"
	Geometry   Geometry      `json:"geometry" bson:"geometry"`
	Properties Properties    `json:"properties" bson:"properties"`
}

