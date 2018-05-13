package config

import (
	"github.com/nawazish-github/airfare-poller/models"
)

type ConfigUnmarshallerInterface interface {
	Unmarshal(file string) (models.Config, error)
}
