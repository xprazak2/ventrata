package utils

import (
	"os"
)

const (
	PortEnvKey = "PORT"
)

func PortEnvVar() string {
	port := os.Getenv(PortEnvKey)
	if port == "" {
		return "3000"
	}
	return port
}
