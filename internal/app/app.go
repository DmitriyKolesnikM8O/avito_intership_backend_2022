package app

import (
	"github.com/DmitriyKolesnikM8O/avito_intership_backend/config"
	"github.com/DmitriyKolesnikM8O/avito_intership_backend/pkg/logger"
	"github.com/DmitriyKolesnikM8O/avito_intership_backend/pkg/postgres"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	logger.SetLogger(cfg.Log.Level)

	log.Info("Initializing postgres")
	pg, err := postgres.New(cfg.Pg.URL, postgres.MaxPoolSize(cfg.Pg.MaxPoolSize))
	if err != nil {
		log.Fatalf("Connection Postgres Error: %s", err)
	}

	defer pg.Close()

	log.Info("Initializing repositories")
	repositories := 

}
