package worker

import (
	"../models"
	"../utils"
	"github.com/jinzhu/gorm"
	"github.com/jrallison/go-workers"
	"os"
	"strings"
)

var db *gorm.DB

func GenerateWaveform(msg *workers.Msg) {
	media := models.Media{}
	id, _ := msg.Get("id").Int()
	if err := db.First(&media, id).Error; err != nil {
		return
	}
	data, duration := utils.GenerateSamplesAsString(media.Url, 4)
	db.Model(&media).Update(&models.Media{
		Waveform: strings.Join(data, ","),
		Duration: duration,
	})
}

func Worky(dbm *gorm.DB) {
	workers.Configure(map[string]string{
		// location of redis instance
		"server": os.Getenv("REDIS"),
		// instance of the database
		"database": "3",
		// number of connections to keep open with redis
		"pool": "30",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})
	db = dbm
	workers.Process("GenerateWaveform", GenerateWaveform, 1)

	// Blocks until process is told to exit via unix signal
	go workers.Run()
}
