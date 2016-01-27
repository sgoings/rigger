package environments

import (
  "fmt"
  "io/ioutil"
  "log"
  "gopkg.in/yaml.v2"
  "github.com/spf13/viper"
)

type EnvironmentsFile struct {
	Default string `json:"default" yaml:"default"`
	Rigs map[string]RigInstance `json:"rigs" yaml:"rigs"`
}

type RigInstance struct {
	Rig string `json:"rig" yaml:"rig"`
	Url string `json:"url" yaml:"url"`
  // Vars RigVars `json:"vars" yaml:"vars"`
}

var Environments EnvironmentsFile

func Load() {
  viper.SetConfigName("rigs")
  viper.AddConfigPath("$HOME/.rigger/")
  viper.AddConfigPath(".rigger/")
  viper.SetConfigType("yaml")

  err := viper.ReadInConfig()
  if err != nil {
    if _,ok := err.(viper.UnsupportedConfigError); !ok {
      log.Printf("[ERROR] %s", err)
    }
  }
  viper.Unmarshal(&Environments)
}

func List() {
  fmt.Printf("Default: %s\n", Environments.Default)

  fmt.Println("Available:")
  for val := range Environments.Rigs {
    fmt.Printf(" - %s (%s)\n", val, Environments.Rigs[val].Rig)
  }
}

func Add(name string, rig RigInstance) {
  if Environments.Rigs == nil {
    Environments.Rigs = make(map[string]RigInstance)
  }
  Environments.Rigs[name] = rig
  Environments.Default = name
}

func Save() {
  bytes, err := yaml.Marshal(Environments)
  if err != nil {
    log.Printf("[ERROR] could not marshal environments list to YAML")
  }
  err = ioutil.WriteFile(".rigger/rigs.yaml", bytes, 0644)
  if err != nil {
    log.Printf("[ERROR] failed to write rigs.yaml file")
  }
}
