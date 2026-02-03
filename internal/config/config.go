package config

import (
	"os"

	"github.com/joho/godotenv"
)



type Config struct {
    Port 						 string
    Env  						 string
    DatabaseURL 				 string
    StorageProvider 			 string
    S3Bucket        			 string
    S3Region          			 string
	S3CfDistro					 string
	CloudFrontDistributionDomain string
}


func Load() *Config {
    _ = godotenv.Load()

    return &Config{
        Port:            			  getEnv("PORT", ""),
        Env:                          getEnv("ENV", ""),
        DatabaseURL:     			  getEnv("DB_PATH", ""),
        StorageProvider: 			  getEnv("STORAGE_PROVIDER", ""),
        S3Bucket:        			  getEnv("S3_BUCKET", ""),
        S3Region:    				  getEnv("S3_REGION", ""),
		S3CfDistro:		 			  getEnv("S3_CF_DISTRO", ""),
		CloudFrontDistributionDomain: getEnv("CloudFrontDistributionDomain", ""),
    }
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}

// func (c *Config) IsDevelopment() bool {
//     return c.Env == "development"
// }

// func (c *Config) IsProduction() bool {
//     return c.Env == "production"
// }