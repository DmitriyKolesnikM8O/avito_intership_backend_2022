package main

import "github.com/DmitriyKolesnikM8O/avito_intership_backend/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
