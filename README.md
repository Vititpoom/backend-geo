# Mini Spatial Data Platform - Backend

This is the Go backend for the Mini Spatial Data Platform, using the Echo framework and MongoDB. It provides a RESTful API to manage location data in GeoJSON format.

## Prerequisites

- **Go** (1.20+ recommended)
- **MongoDB** (Local instance or MongoDB Atlas)

## Project Initialization and Setup

1. **Clone the repository** and navigate to the backend directory:
   ```sh
   cd backend-geo
   ```

2. **Download Go modules**:
   ```sh
   go mod download
   ```

3. **Set up Environment Variables**:
   Create a `.env` file in the root directory (where `main.go` is located) with the following content:
   
   ```env
   PORT=8080
   MONGO_URI=mongodb+srv://<db_username>:<db_password>@cluster.mongodb.net/?appName=Cluster
   DB_NAME=spatial_db
   ```
   *Note: If no `.env` file is present, the application will fallback to default configurations.*

## Running the Application

To run the application locally in development mode:

```sh
go run main.go
```

The server will start on the port specified in your environment variables (default is `8080`).

## Testing the API

For testing the API, you can import the **Postman Collection** included in:
`backend_geo_postman_collection.json`

### Endpoints Overview

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/health` | Health check to verify server is running |
| `GET` | `/api/locations` | Fetch all recorded locations |
| `POST` | `/api/locations` | Create a new location feature (GeoJSON Point) |
| `PUT` | `/api/locations/:id` | Update an existing location feature |
| `DELETE` | `/api/locations/:id` | Delete a location feature by its ID |

### Example Payload (`POST /api/locations`)

```json
{
  "type": "Feature",
  "geometry": {
    "type": "Point",
    "coordinates": [100.5018, 13.7563]
  },
  "properties": {
    "name": "Bangkok",
    "description": "Capital of Thailand"
  }
}
```
