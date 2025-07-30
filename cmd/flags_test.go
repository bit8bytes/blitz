package main

import (
	"testing"

	"github.com/bit8bytes/blitz/pkg/assert"
)

func TestValidateEnv(t *testing.T) {
	tests := []struct {
		name      string
		env       string
		wantError error
	}{
		{
			name:      "Valid development env",
			env:       "development",
			wantError: nil,
		},
		{
			name:      "Valid staging env",
			env:       "staging",
			wantError: nil,
		},
		{
			name:      "Valid production env",
			env:       "production",
			wantError: nil,
		},
		{
			name:      "Invalid empty env",
			env:       "",
			wantError: ErrEnvCannotBeEmpty,
		},
		{
			name:      "Invalid env",
			env:       "foo",
			wantError: ErrEnvInvalid,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateEnv(test.env)
			assert.Equal(t, err, test.wantError)
		})
	}
}

func TestValidateServiceName(t *testing.T) {
	tests := []struct {
		name        string
		serviceName string
		wantError   error
	}{
		{
			name:        "Valid service name",
			serviceName: "blitz",
			wantError:   nil,
		},
		{
			name:        "Valid single letter",
			serviceName: "a",
			wantError:   nil,
		},
		{
			name:        "Valid long service name",
			serviceName: "verylongservicenamebutwithinlimits",
			wantError:   nil,
		},
		{
			name:        "Empty service name",
			serviceName: "",
			wantError:   ErrServiceNameEmpty,
		},
		{
			name:        "Service name with uppercase",
			serviceName: "Blitz",
			wantError:   nil,
		},
		{
			name:        "Service name with numbers",
			serviceName: "blitz123",
			wantError:   nil,
		},
		{
			name:        "Service name with hyphen",
			serviceName: "my-service",
			wantError:   nil,
		},
		{
			name:        "Service name with underscore",
			serviceName: "my_service",
			wantError:   nil,
		},
		{
			name:        "Service name too long",
			serviceName: "thisfoobarservicenameiswaytoolongandexceedsthemaximumlengthof64chars",
			wantError:   ErrServiceNameTooLong,
		},
		{
			name:        "Service name has invalid chat",
			serviceName: "service$",
			wantError:   ErrServiceNameInvalid,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateServiceName(test.serviceName)
			assert.Equal(t, err, test.wantError)
		})
	}
}

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name      string
		user      string
		wantError error
	}{
		{
			name:      "Valid lowercase user",
			user:      "github",
			wantError: nil,
		},
		{
			name:      "Valid uppercase user",
			user:      "GITHUB",
			wantError: nil,
		},
		{
			name:      "Valid mixed case user",
			user:      "GitHub",
			wantError: nil,
		},
		{
			name:      "Valid user with numbers",
			user:      "user123",
			wantError: nil,
		},
		{
			name:      "Valid user with hyphen",
			user:      "my-user",
			wantError: nil,
		},
		{
			name:      "Valid user with underscore",
			user:      "my_user",
			wantError: nil,
		},
		{
			name:      "Valid complex user",
			user:      "User-123_test",
			wantError: nil,
		},
		{
			name:      "Empty user",
			user:      "",
			wantError: ErrUserEmpty,
		},
		{
			name:      "User with space",
			user:      "my user",
			wantError: ErrUserInvalid,
		},
		{
			name:      "User with special chars",
			user:      "user@example",
			wantError: ErrUserInvalid,
		},
		{
			name:      "User with dot",
			user:      "user.name",
			wantError: ErrUserInvalid,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateUser(test.user)
			assert.Equal(t, err, test.wantError)
		})
	}
}

func TestValidateHost(t *testing.T) {
	tests := []struct {
		name      string
		host      string
		wantError error
	}{
		{
			name:      "Valid IPv4 address",
			host:      "192.168.1.10",
			wantError: nil,
		},
		{
			name:      "Valid IPv6 address",
			host:      "2001:db8::1",
			wantError: nil,
		},
		{
			name:      "Valid hostname",
			host:      "example.com",
			wantError: nil,
		},
		{
			name:      "Valid subdomain",
			host:      "api.example.com",
			wantError: nil,
		},
		{
			name:      "Valid localhost",
			host:      "localhost",
			wantError: nil,
		},
		{
			name:      "Valid hostname with hyphen",
			host:      "my-server.example.com",
			wantError: nil,
		},
		{
			name:      "Empty host",
			host:      "",
			wantError: ErrHostEmpty,
		},
		{
			name:      "Invalid hostname with underscore",
			host:      "my_server.com",
			wantError: ErrHostInvalid,
		},
		{
			name:      "Invalid hostname starting with hyphen",
			host:      "-server.com",
			wantError: ErrHostInvalid,
		},
		{
			name:      "Invalid hostname ending with hyphen",
			host:      "server-.com",
			wantError: ErrHostInvalid,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateHost(test.host)
			assert.Equal(t, err, test.wantError)
		})
	}
}
