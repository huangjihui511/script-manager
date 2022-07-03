package model

import "fmt"

var (
	ConfigPath      = "/Users/huangjihui/Library/Mobile Documents/com~apple~CloudDocs/Documents/GitHub/script-manager/config.yaml"
	InternalCommand = []Command{
		{
			Script: fmt.Sprintf("vim \"%s\"", ConfigPath),
			Info:   "edit config with vim",
		},
		{
			Script: fmt.Sprintf("open \"%s\"", ConfigPath),
			Info:   "edit config with your editor",
		},
	}
)

type Command struct {
	Script string `yaml:"script,omitempty"`
	Info   string `yaml:"info,omitempty"`
}

type CommandGroup struct {
	GroupName string    `yaml:"group-name,omitempty"`
	Commands  []Command `yaml:"commands,omitempty"`
}

type Config struct {
	ConfigName    string         `yaml:"config-name,omitempty"`
	CommandGroups []CommandGroup `yaml:"command-groups,omitempty" json:"command_groups,omitempty"`
}
