package db

import (
	"os"
	"os/user"
	"path/filepath"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/gorm"

	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/os/gfile"
)

func Init() {
	createFile()
	migrate()
}

func createFile() {
	u, err := user.Current()
	if err != nil {
		return
	}
	path := filepath.Join(u.HomeDir, ".local/share/weather-cli", "weather.db")
	if gfile.Exists(path) {
		return
	}
	_, _ = os.Create(filepath.Join(u.HomeDir, ".local/share/weather-cli", "weather.db"))
}

func migrate() {
	u, err := user.Current()
	if err != nil {
		return
	}

	path := filepath.Join(u.HomeDir, ".local/share/weather-cli", "weather.db")
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	}()
	db.Logger = logger.Default.LogMode(logger.Silent)
	if err != nil {
		return
	}

	type Weather struct {
		ID        uint   `gorm:"primaryKey"`
		Weather   string `gorm:"index;size:255"`
		Temp      uint8  `gorm:"index"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	if !db.Migrator().HasTable("weathers") {
		_ = db.Migrator().CreateTable(&Weather{})
	}
}
