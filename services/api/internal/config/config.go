package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func LoadDotEnv(paths ...string) error {
	for _, path := range paths {
		loadError := loadDotEnvFile(path)
		if loadError == nil {
			return nil
		}

		if !errors.Is(loadError, os.ErrNotExist) {
			return loadError
		}
	}

	return nil
}

func loadDotEnvFile(path string) error {
	file, openError := os.Open(path)
	if openError != nil {
		return openError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		if setError := os.Setenv(key, value); setError != nil {
			return setError
		}
	}

	return scanner.Err()
}
