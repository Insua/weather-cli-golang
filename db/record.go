package db

import (
	"os/user"
	"path/filepath"
	"time"

	"gorm.io/gorm/logger"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Record(weather string, temp uint) {
	u, err := user.Current()
	if err != nil {
		return
	}

	path := filepath.Join(u.HomeDir, ".local/share/weather-cli", "weather.db")
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	db.Logger = logger.Default.LogMode(logger.Silent)
	if err != nil {
		return
	}

	type Weather struct {
		ID        uint   `gorm:"primaryKey"`
		Weather   string `gorm:"index"`
		Temp      uint   `gorm:"index"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	db.Create(&Weather{
		Weather: weather,
		Temp:    temp,
	})
}
