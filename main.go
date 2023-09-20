package main

import (
	"github.com/gin-contrib/cors"
	"github.com/muhammadsaefulr/inventori-barang/config"
	"github.com/muhammadsaefulr/inventori-barang/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Use(cors.New(config.ConfigCors()))
	r.Run(":8000")
}
