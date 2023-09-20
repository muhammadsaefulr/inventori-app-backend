package config

import "github.com/gin-contrib/cors"

func ConfigCors() cors.Config {
	configs := cors.DefaultConfig()
	configs.AllowOrigins = []string{"*"}
	configs.AllowMethods = []string{"GET", "PUT", "POST", "DELETE"}
	return configs
}
