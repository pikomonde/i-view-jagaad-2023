package initialize

import (
	"i-view-jagaad-2023/model"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewConfig Read and process config file
func NewConfig() (conf model.Config) {

	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Can't find config file, err : %s", err.Error())
		return
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log.Errorf("Unable to decode into config struct : %s", err.Error())
		return
	}

	return
}
