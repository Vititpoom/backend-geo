package config

import "os"

type Config struct {
	Port       string
	MongoURI   string
	DBName     string
}

func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		MongoURI: getEnv("MONGO_URI", "mongodb+srv://<db_username>:<db_password>@cluster.kytft.mongodb.net/?appName=Cluster"),
		DBName:   getEnv("DB_NAME", "spatial_db"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
