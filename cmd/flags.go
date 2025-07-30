package main

import (
	"errors"
	"flag"
	"net"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/bit8bytes/blitz/internal/env"
)

var (
	ErrEnvCannotBeEmpty   = errors.New("env cannot be empty")
	ErrEnvInvalid         = errors.New("invalid env")
	ErrServiceNameEmpty   = errors.New("service name cannot be empty")
	ErrServiceNameTooLong = errors.New("service name too long")
	ErrServiceNameInvalid = errors.New("service name not valid")
	ErrBinaryDirEmpty     = errors.New("binary dir cannot be empty")
	ErrBinaryDirNotExist  = errors.New("binary dir does not exist")
	ErrUnitDirEmpty       = errors.New("unit dir cannot be empty")
	ErrUnitDirNotExist    = errors.New("unit dir does not exist")
	ErrUserEmpty          = errors.New("user cannot be empty")
	ErrUserInvalid        = errors.New("invalid user")
	ErrHostEmpty          = errors.New("host cannot be empty")
	ErrHostInvalid        = errors.New("invalid host")
)

var envs = []string{"development", "staging", "production"}

type Config struct {
	Env         string
	ServiceName string
	BinaryDir   string
	UnitDir     string
	User        string
	Host        string
}

func (cfg *Config) Load() error {
	env.Load()
	flag.StringVar(&cfg.Env, "env", env.GetString("ENV", "production"),
		"Environment (development|staging|production)")
	flag.StringVar(&cfg.ServiceName, "service-name", env.GetString("SERVICE_NAME", ""),
		"Name of service e.g. blitz")
	flag.StringVar(&cfg.BinaryDir, "binary-dir", env.GetString("BINARY_DIR", ""),
		"Directory containing binaries")
	flag.StringVar(&cfg.UnitDir, "unit-dir", env.GetString("UNIT_DIR", ""),
		"Directory containing systemd unit files")
	flag.StringVar(&cfg.User, "user", env.GetString("USER", ""),
		"User")
	flag.StringVar(&cfg.Host, "host", env.GetString("HOST", ""),
		"Hostname or IP")
	flag.Parse()

	if err := validateEnv(cfg.Env); err != nil {
		return err
	}
	if err := validateServiceName(cfg.ServiceName); err != nil {
		return err
	}
	if err := validateBinaryDir(cfg.BinaryDir); err != nil {
		return err
	}
	if err := validateUnitDir(cfg.UnitDir); err != nil {
		return err
	}
	if err := validateUser(cfg.User); err != nil {
		return err
	}
	if err := validateHost(cfg.Host); err != nil {
		return err
	}
	return nil
}

func validateEnv(env string) error {
	if env == "" {
		return ErrEnvCannotBeEmpty
	}
	if !slices.Contains(envs, env) {
		return ErrEnvInvalid
	}
	return nil
}

func validateServiceName(serviceName string) error {
	if serviceName == "" {
		return ErrServiceNameEmpty
	}
	if len(serviceName) > 64 {
		return ErrServiceNameTooLong
	}

	for _, char := range serviceName {
		isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
		isDigit := char >= '0' && char <= '9'
		isAllowedSymbol := char == '-' || char == '_' || char == '.'

		if !isLetter && !isDigit && !isAllowedSymbol {
			return ErrServiceNameInvalid
		}
	}
	return nil
}

func validateBinaryDir(binaryDir string) error {
	if binaryDir == "" {
		return ErrBinaryDirEmpty
	}
	if _, err := os.Stat(binaryDir); os.IsNotExist(err) {
		return ErrBinaryDirNotExist
	}
	return nil
}

func validateUnitDir(unitDir string) error {
	if unitDir == "" {
		return ErrUnitDirEmpty
	}
	if _, err := os.Stat(unitDir); os.IsNotExist(err) {
		return ErrUnitDirNotExist
	}
	return nil
}

func validateUser(user string) error {
	if user == "" {
		return ErrUserEmpty
	}
	for _, char := range user {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_') {
			return ErrUserInvalid
		}
	}
	return nil
}

func validateHost(host string) error {
	if host == "" {
		return ErrHostEmpty
	}

	if net.ParseIP(host) != nil {
		return nil
	}

	if isValidHostname(host) {
		return nil
	}

	return ErrHostInvalid
}

func isValidHostname(hostname string) bool {
	if len(hostname) > 253 {
		return false
	}

	parts := strings.SplitSeq(hostname, ".")
	for part := range parts {
		if len(part) == 0 || len(part) > 63 {
			return false
		}
		for i, char := range part {
			if !((char >= 'a' && char <= 'z') ||
				(char >= 'A' && char <= 'Z') ||
				(char >= '0' && char <= '9') ||
				(char == '-' && i != 0 && i != len(part)-1)) {
				return false
			}
		}
	}
	return true
}

func (cfg *Config) BinaryPath() string {
	return filepath.Join(cfg.BinaryDir, cfg.ServiceName)
}

func (cfg *Config) UnitPath() string {
	return filepath.Join(cfg.UnitDir, cfg.ServiceName+".service")
}
