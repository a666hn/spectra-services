package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	common "github.com/skinnyguy/spectra-services/core/model"
)

type config struct {
	SpectraConfig *common.SpectraConfig
	FileModified  time.Time
}

var (
	appConfig *config
	once      sync.Once
)

func getSpectraConfigInstance() *config {
	once.Do(func() {
		appConfig = new(config)
	})

	return appConfig
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// GetSpectraConfig ...
func GetSpectraConfig() *common.SpectraConfig {
	filename := basepath + "/.app.config.json"

	file, err := os.Stat(filename)
	if err != nil {
		log.Log(errors.BadRequest("common.Config.GetSpectraConfig", err.Error()))
		return nil
	}

	instance := getSpectraConfigInstance()
	if file.ModTime().After(instance.FileModified) {
		log.Info("Load application config from file ...")

		jsonFile, err := os.Open(filename)
		if err != nil {
			log.Log(errors.BadRequest("common.config.GetSpectraConfig", err.Error()))
			return nil
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		var appConfig *common.SpectraConfig

		json.Unmarshal([]byte(byteValue), &appConfig)

		instance.SpectraConfig = appConfig
		instance.FileModified = file.ModTime()
	}

	return instance.SpectraConfig
}
