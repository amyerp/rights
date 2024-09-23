package entrypoint

import (
  "fmt"
	. "rights/version"
	. "rights/global"

	"github.com/spf13/viper"
)

func Init() {
  Update()
	updateversion()
}

func updateversion() {
setingskey := fmt.Sprintf("%s.entrypointversion", MicroServiceName)
viper.Set(setingskey, VERSIONPLUGIN)
}
