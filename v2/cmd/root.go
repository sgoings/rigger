// Copyright © 2015 Seth Goings
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
  "log"
  "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Rigfile struct {
  Workflows []Workflow
}

type Workflow struct {
  Interface WorkflowInterface
  Implementations []WorkflowImplementation
}

type WorkflowInterface struct {
  Url string
}

type WorkflowImplementation struct {
  Name string
  Url string
}

// type WorkflowConfig {
//
// }

// type Rigfile struct {
//   Tools []Tool
//   Commands []Command
// }
//
// type Tool struct {
//   Name string
//   Type string
//   Install InstallMethod
// }
//
// type InstallMethod struct {
//   Method string
//   Url string
//   Env []string
//   Path []string
// }
//
// type Providers struct {
//   Name string
// }

type WorkflowDefinition struct {
  Metadata WorkflowMetadata
  Commands []WorkflowCommands
}

type WorkflowMetadata struct {
  Name string
  Description string
}

type WorkflowCommands struct {
  Name string
  Description string
  Inputs WorkflowInputs
  Outputs WorkflowOutputs
}

type WorkflowOutputs struct {
  Vars []string
}

type WorkflowInputs struct {
  Vars []string
}


var rig Rigfile
var workflow WorkflowDefinition
var commands map[string]*cobra.Command
var cfgFile string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "rigger [command]",
	Short: "rigger helps you assemble and construct things easily",
	Long: `...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
  loadCommands()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.makeup/makeup.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

func loadCommands() {
  commands = make(map[string]*cobra.Command)

  viper.SetConfigName("Rigfile")
  viper.AddConfigPath("$HOME/.rigger/")
  viper.AddConfigPath(".")
  viper.SetConfigType("yaml")

  err := viper.ReadInConfig()
  if err != nil {
    if _,ok := err.(viper.UnsupportedConfigError); ok {
      log.Printf("[ERROR] No Rigfile exists.")
      os.Exit(1)
    } else {
      log.Printf("[ERROR] %s", err)
    }
  }

  viper.Unmarshal(&rig)

  workflowYaml := rig.Workflows[0].Interface.Url + "/workflow.yaml"

  // log.Printf("Location of workflow interface: %v", workflowYaml)

  file, err := os.Open(workflowYaml)
  if err != nil {
    log.Printf("[ERROR] %s", err)
  } else {
    viper.ReadConfig(file)
    viper.Unmarshal(&workflow)
  }

  RootCmd.Long = workflow.Metadata.Description

  for _,element := range workflow.Commands {
    // log.Printf("[DEBUG] Adding command %s", element.Name)
    commands[element.Name] = &cobra.Command{
      Use:   element.Name,
      Short: strings.TrimSpace(element.Description),
      Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
      },
    }
    RootCmd.AddCommand(commands[element.Name])
  }
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  log.Print("[DEBUG] Loading runtime config...")
}
