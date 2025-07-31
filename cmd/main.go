package main

import (
	"fmt"
	"log/slog"
	"os"
)

type CLI struct {
	config *Config
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	if len(os.Args) < 2 {
		fmt.Println("Blitz ⚡️")
		fmt.Println("Usage: blitz <command>")
		fmt.Println("Commands:")
		fmt.Println("  deploy    Deploy binary using systemd service")
		os.Exit(1)
	}

	command := os.Args[1]
	os.Args = append(os.Args[:1], os.Args[2:]...)

	cfg := &Config{}
	err := cfg.Load()
	if err != nil {
		logger.Error("Error loading or parsing flags.", "err", err)
		os.Exit(1)
	}

	cli := CLI{
		config: cfg,
		logger: logger,
	}

	switch command {
	case "deploy":
		err := cli.Deploy()
		if err != nil {
			logger.Error("Deploy failed", "err", err)
			os.Exit(1)
		}
	default:
		logger.Error("Unknown command", "command", command)
		os.Exit(1)
	}
}
