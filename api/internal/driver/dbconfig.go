// Módulo de intuito apenas "organizacional" de configurações

package driver

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

// Função de atalho de formatação
func (d *DatabaseConfig) GetDNS() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		d.Host, d.Port, d.DBname, d.User, d.Password)
}

// Função apenas para "embelezamento" do código
func LoadDatabaseConfig() DatabaseConfig {
	// default values

	// gotenv.Apply(strings.NewReader("APP_ID=1234567"))

	return DatabaseConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBname:   os.Getenv("DATABASE_NAME"),
	}
}
