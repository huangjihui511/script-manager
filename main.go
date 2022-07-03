package main

import (
	"huangjihui511/script-manager/model"
	"huangjihui511/script-manager/pkg"
)

func init() {

}

func main() {
	pkg.InitConfigPath()
	pkg.CreateUI(func() (model.View, error) {
		cfg, err := pkg.ReadConfig(model.ConfigPath)
		if err != nil {
			return model.View{}, err
		}
		viewData, err := pkg.TransConfig2View(cfg)
		return viewData, err
	})
}
