package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	var cfg Config
	err := cfg.Load()
	if err != nil {
		logger.Error("Error loading or parsing flags.", "err", err)
		os.Exit(1)
	}

	fmt.Println("Blitz ⚡️")

	// Todos:
	// - [ ] blitz deploy
	// - [ ] blitz rollback
	// - [ ] blitz backup
}
