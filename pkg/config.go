package pkg

import (
	"fmt"
	"huangjihui511/script-manager/model"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ConfigPathEnv = "SM_CONF_PATH"
)

func InitConfigPath() {
	configPath, ok := os.LookupEnv(ConfigPathEnv)
	if ok {
		model.ConfigPath = configPath
		return
	}
	CreateErrorUI(fmt.Errorf("Not find the definition yaml file path in environment variable, please see information in help"))
}

func ReadConfig(path string) (model.Config, error) {
	config := model.Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func TransConfig2View(config model.Config) (model.View, error) {
	viewConfigGroups := []model.CommandGroup{
		{
			GroupName: "internal",
			Commands:  model.InternalCommand,
		},
	}
	viewConfigGroups = append(viewConfigGroups, config.CommandGroups...)
	return model.View{
		ViewName:      config.ConfigName,
		CommandGroups: viewConfigGroups,
	}, nil
}
