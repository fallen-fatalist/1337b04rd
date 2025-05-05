package cmd

import (
	"log"
	"log/slog"
	"os"

	"1337bo4rd/internal/adapter/config"
	flag "1337bo4rd/internal/adapter/config"
	"1337bo4rd/internal/adapter/logger"
)

func Run() {
	// Parse flags
	err := flag.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Load environment variables
	config := config.New()

	// Set logger
	logger.Set(config.App)
	slog.Info("Staring application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	// ctx := context.Background()
	// db, err :=
}
