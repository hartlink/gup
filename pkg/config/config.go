package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config representa la configuración de CliGO
type Config struct {
	Language string `json:"language"` // "es" o "en"
}

// GetConfigPath retorna la ruta del archivo de configuración
func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".gup", "config.json")
}

// Load carga la configuración desde el archivo
func Load() (*Config, error) {
	configPath := GetConfigPath()

	// Si no existe el archivo, retornar configuración por defecto
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{Language: ""}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{Language: ""}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &Config{Language: ""}, err
	}

	return &config, nil
}

// Save guarda la configuración en el archivo
func Save(config *Config) error {
	configPath := GetConfigPath()

	// Crear directorio si no existe
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
